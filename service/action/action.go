package action

import (
	"alfred/constant/command_constant"
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)

func Action(c *cli.Context) error {
	if c.NArg() == 0 {
		now := time.Now().Local()
		var nowTime string
		if now.Hour() >= 6 && now.Hour() < 12 {
			nowTime = "Good morning, "
		} else if now.Hour() >= 12 && now.Hour() < 20 {
			nowTime = "Good afternoon, "
		} else {
			nowTime = "Good evening, "
		}
		fmt.Println(nowTime + "please input command. ")
	} else {
		switch c.Args().Get(0) {
		case command_constant.RECORD:
			// TODO by 1123 write code here
			fmt.Println("some trash talk")
			var weight int
			fmt.Scanf("Pleash input your weight here: %d", weight)
			fmt.Println(weight)
		case command_constant.ADD:
			// TODO by 1123 write code here too
			fmt.Println("some trash talk again")
		}
	}
	return nil
}
