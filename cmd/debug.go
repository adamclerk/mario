package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/adamclerk/mario/lib"
	"github.com/codegangsta/cli"
)

type debugFactory struct{}

type debug struct {
	line   string
	writer io.Writer
}

func (f *debugFactory) Make(line string, context *cli.Context) lib.Task {
	return &debug{line: line, writer: os.Stdout}
}

func (d *debug) Print() {
	fmt.Fprintln(d.writer, d.line)
}

func (d *debug) Process() {
	return
}

func AddDebug() cli.Command {
	return cli.Command{
		Name:      "debug",
		ShortName: "d",
		Usage:     "just pass the line through",
		Action: func(c *cli.Context) {
			d := debugFactory{}
			lib.Run(&d, c)
		},
	}
}
