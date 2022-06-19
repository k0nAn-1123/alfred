package main

import (
	"alfred/service/command"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
	}
	app.Commands = command.Command()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
