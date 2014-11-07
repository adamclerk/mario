package lib

import (
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/codegangsta/cli"
)

type csvTemplateFactory struct{}

type csvTemplate struct {
	separator string
	template  string
	line      string
	writer    io.Writer
}

func (f *csvTemplateFactory) make(line string, context *cli.Context) task {
	return &csvTemplate{
		line:      line,
		writer:    os.Stdout,
		template:  context.String("out"),
		separator: context.String("separator"),
	}
}

func (d *csvTemplate) print() {
	values := strings.Split(d.line, d.separator)
	tmpl, _ := template.New("csv").Parse(d.template + "\n")
	tmpl.Execute(d.writer, values)
}

func (d *csvTemplate) process() {
	return
}

func AddCSVTemplate() cli.Command {
	return cli.Command{
		Name:      "csvTemplate",
		ShortName: "c",
		Usage:     "change csv input",
		Action: func(c *cli.Context) {
			d := csvTemplateFactory{}
			Run(&d, c)
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "out",
				Value: "{{. index 1}},{{. index 2}}",
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
