package main

import (
	"os"

	"github.com/adamclerk/mario/lib"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "mario"
	app.Usage = "mama-mia I love pipes. Pipe input to me and I'll do amazing things."
	app.Commands = []cli.Command{
		lib.AddDebug(),
		lib.AddCsvReplace(),
	}
	app.Run(os.Args)
}
