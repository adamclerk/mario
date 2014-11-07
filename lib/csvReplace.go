package lib

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io"
	"os"
	//"strings"
)

type csvReplaceFactory struct{}

type csvReplace struct {
	separator string
	template  string
	line      string
	writer    io.Writer
}

func (f *csvReplaceFactory) make(line string, context *cli.Context) task {
	return &csvReplace{
		line:      line,
		writer:    os.Stdout,
		template:  context.String("out"),
		separator: context.String("separator"),
	}
}

func (d *csvReplace) print() {
	//values := strings.Split(d.line, ",")

	fmt.Fprintln(d.writer, d.template)
}

func (d *csvReplace) process() {
	return
}

func AddCsvReplace() cli.Command {
	return cli.Command{
		Name:      "csvReplace",
		ShortName: "c",
		Usage:     "change csv input",
		Action: func(c *cli.Context) {
			d := csvReplaceFactory{}
			Run(&d, c)
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "out",
				Value: "[...3],[4],[5...6],[7][8...]",
				Usage: "new line constructed from the args of the input",
			},
			cli.StringFlag{
				Name:  "separator",
				Value: ",",
				Usage: "separator character",
			},
		},
	}
}
