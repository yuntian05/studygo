package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errpanicHandler(writer http.ResponseWriter, request *http.Request) error {
	panic(122)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown error")
}

func errNo(writer http.ResponseWriter, request *http.Request) error {
	_, _ = fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	handler appHandler
	code    int
	message string
}{
	{errpanicHandler, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{errNo, 200, "no error"},
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.handler)
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.imooc.com",
			nil)
		response := httptest.NewRecorder()
		f(response, request)
		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.handler)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectMsg string, t *testing.T)  {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode ||
		body != expectMsg {
		t.Errorf("expected:(%d, %s), got (%d, %s)", expectedCode, expectMsg, resp.StatusCode, body)
	}
}
