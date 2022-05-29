package action

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func Action(c *cli.Context) error {
	name := "world"
	if c.NArg() > 0 {
		name = c.Args().Get(0)
	}

	if c.String("lang") == "english" {
		fmt.Println("hello", name)
	} else {
		fmt.Println("你好", name)
	}
	return nil
}
