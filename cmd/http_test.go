package cmd

import (
	"io"
	"io/ioutil"
	"testing"
)

type httpProcessTest struct {
	line     string
	expected string
}

var httpProcessTestCases = []httpProcessTest{
/*{"one,two,three", ""},
  {"two,three,four", ""},
  {"club,breakfast", ""},*/
}

func TestHttpPrint(t *testing.T) {}

func TestHttpMake(t *testing.T) {}

func TestHTTPProcess(t *testing.T) {
	for _, test := range httpProcessTestCases {
		r, w := io.Pipe()
		go func() {
			cmd := http{writer: w}
			cmd.Print()
			w.Close()
		}()
		result, _ := ioutil.ReadAll(r)
		if test.expected+"\n" != string(result) {
			t.Error("Expect '"+test.line+"\n'", "'"+string(result)+"'")
		}
	}
}
