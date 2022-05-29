package flags

import "github.com/urfave/cli/v2"

var Lang = &cli.StringFlag{
	Name:  "lang",
	Value: "english",
	Usage: "language for the meeting",
}
