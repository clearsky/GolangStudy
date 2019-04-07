package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func errPanic(write http.ResponseWriter, request *http.Request) error{
	panic(123)
}

func TestErrWrapper(t *testing.T){
	tests := []struct{
		h appHandle
		code int
		message string
	}{
		{errPanic, 500, "Internal Server Error"},
	}

	for _, tt := range tests{
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.immoc.com", nil)
		f(response, request)
		b, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(b), "\n")
		if response.Code != tt.code || body != tt.message{
			t.Errorf("expect(%d, %s), get(%d, %s)", tt.code, tt.message, response.Code, body)
		}
	}
}
