package httplib

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_libTransport_RoundTrip(t *testing.T) {
	type fields struct {
		underlyingTransport http.RoundTripper
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &libTransport{
				underlyingTransport: tt.fields.underlyingTransport,
			}
			got, err := l.RoundTrip(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("libTransport.RoundTrip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("libTransport.RoundTrip() = %v, want %v", got, tt.want)
			}
		})
	}
}
