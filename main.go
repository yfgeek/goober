package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

const usage = "goober is a smallest implementation of docker"

func main() {
	app := cli.NewApp()
	app.Name = "goober"
	app.Usage = usage

	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}

	app.Before = func(context *cli.Context) error {
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
