package command

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli/v2"
)

func Command() []*cli.Command {
	var command = []*cli.Command{
		{
			Name:    "record",
			Aliases: []string{"r"},
			Usage:   "record health data",
			Action: func(c *cli.Context) error {
				var weight, height int

				fmt.Println("Please input your weight")
				fmt.Scanf("%d\n", &weight)
				fmt.Println("Please input your height")
				fmt.Scanf("%d\n", &height)
				fmt.Printf("Your height: %d, your weight: %d", height, weight)

				return nil
			},
		},

		{
			Name:    "signin",
			Aliases: []string{"s"},
			Usage:   "sign in",
			Action: func(c *cli.Context) error {
				// var user model.User
				// user.Name = "Nicolas"
				// user.Age = 25
				// user.Gender = user_constant.Gender_Male
				// user.Permission = user_constant.Permission_Master
				// user.Extra = model.UserExtra{
				// 	PhoneNumber: []string{"18671056539", "17665358971"},
				// 	Address:     "",
				// }

				// if err := dal.CreateUser(user); err != nil {
				// 	return err
				// }
				var filePath = "/d/Code/go/src/alfred/conf/config.toml"
				fmt.Println(filePath)

				f, err := os.Open(filePath)
				if err != nil {
					fmt.Println(err.Error())
				}
				defer f.Close()
				bytes, _ := ioutil.ReadAll(f)
				fmt.Println(string(bytes))

				return nil
			},
		},
	}

	return command
}
