package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Mantle"
	app.Usage = "Package manager for CoreOS"
	app.Action = func(c *cli.Context) {
		println("TBA")
	}

	app.Run(os.Args)
}
