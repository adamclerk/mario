package lib

import (
	"io"
	"io/ioutil"
	"testing"
)

type test struct {
	line string
}

var testCases = []test{
	{"Lorem ipsum dolor sit amet, consectetur adipisicing elit."},
	{"Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."},
	{"Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."},
}

func TestDebugPrint(t *testing.T) {

	for _, test := range testCases {
		r, w := io.Pipe()
		go func() {
			cmd := debug{line: test.line, writer: w}
			cmd.print()
			w.Close()
		}()
		result, _ := ioutil.ReadAll(r)
		if test.line+"\n" != string(result) {
			t.Error("Expect '"+test.line+"\n'", "'"+string(result)+"'")
		}
	}
}
