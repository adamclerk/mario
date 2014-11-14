package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	nethttp "net/http"
	"net/url"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/adamclerk/mario/lib"
	"github.com/codegangsta/cli"
	"github.com/twinj/uuid"
)

type httpFactory struct{}

type http struct {
	response     *nethttp.Response
	request      *nethttp.Request
	err          error
	responseTime time.Duration
	tmpdir       string
	line         string
	file         string
	writer       io.Writer
}

func (f *httpFactory) Make(line string, context *cli.Context) lib.Task {
	return MakeHTTP(line, context.String("separator"), context.String("tmpdir"))
}

func MakeHTTP(line string, separator string, tmpdir string) *http {
	values := strings.Split(line, separator)

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
		tmpdir:  tmpdir,
		line:    line,
		err:     nil,
	}
}

func (h *http) Print() {

	fmt.Fprintf(h.writer, "%s,%d,%d,%s\n", h.line, h.response.StatusCode, h.responseTime/1000000, h.file)
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

	usr, _ := user.Current()
	dirErr := os.MkdirAll(usr.HomeDir+"/"+h.tmpdir+"/http", 0755)
	h.file = usr.HomeDir + "/" + h.tmpdir + "/http/" + uuid.Formatter(uuid.NewV4(), uuid.Clean) + ".body"
	defer h.response.Body.Close()
	body, err := ioutil.ReadAll(h.response.Body)
	if dirErr != nil {
		body = []byte("error")
	}
	writeErr := ioutil.WriteFile(h.file, body, 0664)
	if writeErr != nil {
		panic(writeErr)
		h.file = "write error"
	}
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
			cli.StringFlag{
				Name:  "tmpdir",
				Value: ".mario",
				Usage: "temp directory where we store the body",
			},
		},
	}
}
