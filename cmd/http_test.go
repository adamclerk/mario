package cmd

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

type httpMakeTest struct {
	line    string
	method  string
	url     string
	headers string
}

var httpMakeTestCases = []httpMakeTest{
	{"http://www.google.com", "GET", "http://www.google.com", ""},
	{"http://www.facebook.com,POST", "POST", "http://www.facebook.com", ""},
	{"http://www.twitter.com,PUT,accept: text/html", "PUT", "http://www.twitter.com", "Accept: text/html"},
}

func TestHttpMake(t *testing.T) {
	for _, test := range httpMakeTestCases {
		http := MakeHTTP(test.line, ",", ".mariotest")
		if test.url != http.request.URL.String() {
			t.Error("Expect '"+test.url, "'"+http.request.URL.String()+"'")
		}

		if test.method != http.request.Method {
			t.Error("Expect '"+test.method, "'"+http.request.Method+"'")
		}

		r, w := io.Pipe()
		go func() {
			http.request.Header.Write(w)
			w.Close()
		}()
		result, _ := ioutil.ReadAll(r)

		if test.headers != strings.TrimSpace(string(result)) {
			t.Error("Expect '"+test.headers, "'"+strings.TrimSpace(string(result))+"'")
		}
	}
}

func TestHttpPrint(t *testing.T) {}

func TestHTTPProcess(t *testing.T) {}
