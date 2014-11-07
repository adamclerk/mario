package cmd

import (
	"io"
	"io/ioutil"
	"testing"
)

type csvTemplateTest struct {
	line     string
	template string
	expected string
}

var csvTestCases = []csvTemplateTest{
	{"one,two,three", "{{index . 1}}", "two"},
	{"two,three,four", "{{index . 2}},{{index . 1}}", "four,three"},
	{"club,breakfast", "{{index . 1}} {{index . 0}}", "breakfast club"},
}

func TestCSVTemplatePrint(t *testing.T) {
	for _, test := range csvTestCases {
		r, w := io.Pipe()
		go func() {
			cmd := csvTemplate{line: test.line, writer: w, template: test.template, separator: ","}
			cmd.Print()
			w.Close()
		}()
		result, _ := ioutil.ReadAll(r)
		if test.expected+"\n" != string(result) {
			t.Error("Expect '"+test.line+"\n'", "'"+string(result)+"'")
		}
	}
}
