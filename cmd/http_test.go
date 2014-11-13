package cmd

import (
	"io"
	"io/ioutil"
	"testing"
)

type httpTest struct {
	line string
}

var httpTestCases = []httpTest{
	{"one,two,three"},
	{"two,three,four"},
	{"club,breakfast"},
}

func TestHTTPProcess(t *testing.T) {
	for _, test := range httpTestCases {
		r, w := io.Pipe()
		go func() {
			cmd := http{line: test.line, writer: w}
			cmd.Print()
			w.Close()
		}()
		result, _ := ioutil.ReadAll(r)
		if test.expected+"\n" != string(result) {
			t.Error("Expect '"+test.line+"\n'", "'"+string(result)+"'")
		}
	}
}
