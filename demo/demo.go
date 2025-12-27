package demo

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func UrfaveCli() {
	app := cli.NewApp()
	app.Name = "myapp"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "language, l",
			Value: "english",
			Usage: "language to greet",
		},
		&cli.IntFlag{
			Name:  "config, c",
			Usage: "config file to use",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println("complete a task on the list")
				for i, v := range c.Args().Slice() {
					fmt.Println(i, v)
				}
				return nil
			},
		}, {
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) error {
				fmt.Println("add a task to the list")
				for i, v := range c.Args().Slice() {
					fmt.Println(i, v)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
