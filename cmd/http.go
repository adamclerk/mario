package cmd

import (
	"fmt"
	"io"
	nethttp "net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/adamclerk/mario/lib"
	"github.com/codegangsta/cli"
)

type httpFactory struct{}

type http struct {
	response     *nethttp.Response
	request      *nethttp.Request
	err          error
	responseTime time.Duration
	writer       io.Writer
}

func (f *httpFactory) Make(line string, context *cli.Context) lib.Task {
	values := strings.Split(line, context.String("separator"))

	request := nethttp.Request{}
	url, err := url.Parse(values[0])

	if err != nil {
		return &http{
			err: err,
		}
	} else {
		request.URL = url
	}

	request.Method = "GET"
	if len(values) > 1 {
		request.Method = values[1]
	}

	request.Header = map[string][]string{}
	if len(values) > 2 {
		for _, header := range strings.Split(values[2], ";") {
			keyval := strings.Split(header, ":")
			request.Header.Add(keyval[0], keyval[1])
		}
	}

	return &http{
		request: &request,
		writer:  os.Stdout,
		err:     nil,
	}
}

func (h *http) Print() {
	fmt.Fprintf(h.writer, "%d,%s,%d\n", h.response.StatusCode, h.responseTime, h.responseTime)
}

func (h *http) Process() {

	if h.err != nil {
		return
	}
	client := nethttp.Client{}
	time_start := time.Now()
	response, err := client.Do(h.request)
	if err != nil {
		h.err = err
		return
	}
	h.responseTime = time.Since(time_start)
	h.response = response
}

func AddHTTP() cli.Command {
	return cli.Command{
		Name:      "http",
		ShortName: "w",
		Usage:     "request a http resource",
		Action: func(c *cli.Context) {
			d := httpFactory{}
			lib.Run(&d, c)
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "separator",
				Value: ",",
				Usage: "separator character",
			},
		},
	}
}
