package main

import (
	"os"

	"github.com/adamclerk/mario/cmd"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "mario"
	app.Usage = "mama-mia I love pipes. Pipe input to me and I'll do amazing things."
	app.Commands = []cli.Command{
		cmd.AddDebug(),
		cmd.AddCSVTemplate(),
		cmd.AddHTTP(),
	}
	app.Run(os.Args)
}
