package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	// JWTDefaultTTL default lieftime of jwt token
	JWTDefaultTTL int = 30

	// APIv1 api version
	APIv1 string = "1.0"

	// APIv2 api version
	APIv2 string = "2.0"

	// APIv3 api version
	APIv3 string = "3.0"

	// BasicAuth type
	BasicAuth string = "BasicAuth"
	// JWTAuth type
	JWTAuth string = "JWTAuth"

	// MDAPICategory md api url prefix
	MDAPICategory string = "md"

	// TRADEAPICategory traders api url prefix
	TRADEAPICategory string = "trade"

	// CurrentAPIVersion default api version
	CurrentAPIVersion string = APIv2

	emptyString           string = ""
	commaSeparator        string = ","
	slashSeparator        string = "/"
	jwtAuthTokenPrefix    string = "Bearer"
	basicAuthPreffix      string = "Basic"
	acceptJSON            string = "application/json"
	acceptStream          string = "application/x-json-stream"
	acceptEncoding        string = "gzip"
	contentEncodingHeader string = "Content-Encoding"
	authHeader            string = "Authorization"
)

// APIVersions list of available api verisons
var APIVersions = [3]string{APIv1, APIv2, APIv3}

// Scopes is JWT Auth scopes
var Scopes = []string{
	"crossrates", "change", "crossrates", "summary",
	"symbols", "feed", "ohlc", "orders", "transactions",
	"accounts",
}

// Represents last remember auth token
var currentJwtAuthToken string

func setReqHeader(r *http.Request, h, v string) { r.Header.Set(h, v) }

// Standard jwt-go claims does not support multiple audience
type claimsWithMultiAudSupport struct {
	Aud []string `json:"aud"`
	jwt.StandardClaims
}

type requestData struct {
	category, action, version, auth, pathParams string
	queryStringParams                           map[string]string
}

// getAuth getter for requestData.auth field
// return JWT is default auth type
func (h HTTPApi) getAuth(u requestData) IAuth {
	switch u.auth {
	case JWTAuth:
		return h.Auth.JWT
	case BasicAuth:
		return h.Auth.Basic
	default:
		return h.Auth.JWT
	}
}

// getAPIVersion getter for requestData.version field
// return APIv2 is default API version
func (r requestData) getAPIVersion() string {
	switch r.version {
	case APIv1:
	case APIv2:
	case APIv3:
	default:
		return APIv2
	}
	return r.version
}

// getCategory getter for requestData.category field
// return md is default api category
func (r requestData) getCategory() string {
	switch r.category {
	case MDAPICategory:
	case TRADEAPICategory:
	case emptyString:
		return MDAPICategory
	default:
		panic(errUndefinedCategoryAPIMessage)
	}
	return r.category
}

// libTransport custom transport layer
type libTransport struct {
	underlyingTransport http.RoundTripper
}

// RoundTrip with custom headers setup
func (t *libTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Accept", acceptJSON)
	req.Header.Add("Accept-Encoding", acceptEncoding)
	if req.Method == http.MethodPost {
		req.Header.Add("Content-type", acceptJSON)
	}
	if strings.Contains(req.URL.Path, "stream") {
		req.Header.Add("Accept", acceptStream)
	}
	return t.underlyingTransport.RoundTrip(req)
}

// JWTAuthMethod ...
type JWTAuthMethod struct {
	ApplicationID, ClientID, SharedKey string
	JwtTTL                             int
}

// BasicAuthMethod ...
type BasicAuthMethod struct{ Username, Password string }

// Auth ...
type Auth struct {
	JWT   JWTAuthMethod
	Basic BasicAuthMethod
}

// HTTPApi struct provide communication with xnt services
type HTTPApi struct {
	Auth                Auth
	httpClient          *http.Client
	baseAPIURL, version string
}

// IAuth ...
type IAuth interface{ authenticate(r *http.Request) }

func (b BasicAuthMethod) authenticate(r *http.Request) {
	r.SetBasicAuth(b.Username, b.Password)
}

func (j JWTAuthMethod) authenticate(r *http.Request) {
	isValidToken := j.validateJwtToken()
	if !isValidToken {
		currentJwtAuthToken = j.getJWTToken()
	}
	authValue := fmt.Sprintf(
		"%s %s", jwtAuthTokenPrefix, currentJwtAuthToken)
	setReqHeader(r, authHeader, authValue)
}

func (j JWTAuthMethod) validateJwtToken() bool {
	if currentJwtAuthToken == emptyString {
		return false
	}
	token, _ := jwt.ParseWithClaims(
		currentJwtAuthToken, &claimsWithMultiAudSupport{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SharedKey), nil
		})
	return token.Valid
}

func (j JWTAuthMethod) getJWTToken() string {
	now := time.Now()
	ttl := j.getTTL()
	jwtExpiresAt := now.Add(ttl).Unix()
	jwtIssueAt := now.Unix()

	claims := claimsWithMultiAudSupport{
		Scopes,
		jwt.StandardClaims{
			Issuer:    j.ClientID,
			Subject:   j.ApplicationID,
			IssuedAt:  jwtIssueAt,
			ExpiresAt: jwtExpiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(j.SharedKey))
	return tokenString
}

func (j JWTAuthMethod) getTTL() time.Duration {
	v := JWTDefaultTTL
	if j.JwtTTL > 0 {
		v = j.JwtTTL
	}
	ttl := time.Second * time.Duration(v)
	return ttl
}

func (h HTTPApi) getVersion() string {
	for _, apiVersion := range APIVersions {
		if apiVersion == h.version {
			return apiVersion
		}
	}
	panic(errUndefinedAPIVersionMessage)
}

func (h HTTPApi) request(
	method, url string, a IAuth, p *bytes.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, p)
	if err != nil {
		return nil, err
	}
	a.authenticate(req)
	resp, err := h.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h HTTPApi) buildURL(u requestData) string {
	varsion := u.getAPIVersion()
	apiURL := fmt.Sprintf(
		"%s/%s/%s/%s", h.baseAPIURL, u.getCategory(), varsion, u.action)

	if len(u.pathParams) > 0 {
		apiURL = fmt.Sprintf("%s/%s/", apiURL, u.pathParams)
	}
	if len(u.queryStringParams) > 0 {
		qParams := url.Values{}
		for k := range u.queryStringParams {
			qParams.Add(k, u.queryStringParams[k])
		}
		apiURL = apiURL + "?" + qParams.Encode()
	}
	return apiURL
}

func (h HTTPApi) get(m interface{}, u requestData) error {
	err := h.fetch(http.MethodGet, m, u, emptyPostPayload)
	return err
}

func (h HTTPApi) post(m interface{}, u requestData) error {
	objByte, err := h.preparePostPayload(m)
	if err != nil {
		return err
	}
	err = h.fetch(http.MethodPost, m, u, objByte)
	return err
}

func (h HTTPApi) preparePostPayload(m interface{}) (*bytes.Reader, error) {
	o, err := typeToJSON(m)
	if err != nil {
		return nil, err
	}
	objByte := bytes.NewReader(o)
	return objByte, err
}

func (h HTTPApi) stream(u requestData, stopChan chan bool, outChan chan []byte) {
	resp, err := h.request(
		http.MethodGet, h.buildURL(u), h.getAuth(u), emptyPostPayload)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	g, err := gzip.NewReader(resp.Body)
	io.TeeReader(g, &buf)
	for {
		select {
		default:
			r, err := ioutil.ReadAll(&buf)
			if err != nil {
				panic(err)
			}
			outChan <- r
		case <-stopChan:
			return
		}
	}
}

func (h HTTPApi) runStream(u requestData) (chan []byte, chan bool) {
	stopChan := make(chan bool)
	outChan := make(chan []byte)
	go h.stream(u, stopChan, outChan)
	return outChan, stopChan
}

func (h HTTPApi) fetch(
	httpMethod string, m interface{}, u requestData, payload *bytes.Reader) error {

	resp, err := h.request(httpMethod, h.buildURL(u), h.getAuth(u), payload)
	if err != nil {
		return err
	}
	preparedData, err := h.processResponse(resp)
	if err != nil {
		return err
	}
	err = h.serialize(preparedData, m)
	return err
}

func (h HTTPApi) processResponse(resp *http.Response) ([]byte, error) {
	var responseBody []byte
	switch resp.Header.Get(contentEncodingHeader) {
	case acceptEncoding:
		g, _ := gzip.NewReader(resp.Body)
		responseBody, _ = ioutil.ReadAll(g)
		defer g.Close()
	default:
		responseBody, _ = ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"Response status code - %d \n %s", resp.StatusCode, responseBody)
	}
	return responseBody, nil
}

func (h HTTPApi) serialize(data []byte, model interface{}) (err error) {
	err = json.Unmarshal(data, &model)
	return
}
