package cmd

import (
	"fmt"
	"io"
	nethttp "net/http"
	"os"
	"strings"

	"github.com/adamclerk/mario/lib"
	"github.com/codegangsta/cli"
)

type httpFactory struct{}

type http struct {
	response     nethttp.Response
	responseTime int
	line         string
	url          string
	separator    string
	method       string
	headers      string
	protocol     string
	form         string
	writer       io.Writer
}

func (f *httpFactory) Make(line string, context *cli.Context) lib.Task {
	values := strings.Split(line, context.String("separator"))

	protocol := "http"
	if strings.Index(values[0], "https") == 0 {
		protocol = "https"
	}

	method := "GET"
	if len(values) > 1 {
		method = values[1]
	}

	headers := ""
	if len(values) > 2 {
		headers = values[2]
	}

	return &http{
		url:       values[0],
		method:    method,
		protocol:  protocol,
		headers:   headers,
		line:      line,
		writer:    os.Stdout,
		separator: context.String("separator"),
	}
}

func (h *http) Print() {
	fmt.Fprintln(h.writer, "Test")
}

func (h *http) Process() {
	switch h.method {

	}
	return
}

func AddHTTP() cli.Command {
	return cli.Command{
		Name:      "http",
		ShortName: "h",
		Usage:     "request a http resource",
		Action: func(c *cli.Context) {
			d := httpFactory{}
			lib.Run(&d, c)
		},
	}
}
