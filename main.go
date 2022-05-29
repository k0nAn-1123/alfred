package main

import (
	"alfred/service/action"
	"alfred/service/flags"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:   "alfred",
		Usage:  "my personal butler",
		Action: action.Action,
		Flags: []cli.Flag{
			flags.Lang,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
