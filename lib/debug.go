package lib

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io"
	"os"
)

type debugFactory struct{}

type debug struct {
	line   string
	writer io.Writer
}

func (f *debugFactory) make(line string, context *cli.Context) task {
	return &debug{line: line, writer: os.Stdout}
}

func (d *debug) print() {
	fmt.Fprintln(d.writer, "DEBUG")
	fmt.Fprintln(d.writer, d.line)
}

func (d *debug) process() {
	return
}

func AddDebug() cli.Command {
	return cli.Command{
		Name:      "debug",
		ShortName: "d",
		Usage:     "just pass the line through",
		Action: func(c *cli.Context) {
			d := debugFactory{}
			Run(&d, c)
		},
	}
}
