package main

import (
	"awesomeProject/cmd/ai"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "ai-help",
				Usage: "current command can help with any command which" +
					"you need to help," +
					"it's analogue of man command but with examples",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "c",
						Usage: "command for help with ai",
					},
				},
				Action: func(c *cli.Context) error {
					ai.Help(c)
					return nil
				},
			},
			{
				Name:  "ai-cmd",
				Usage: "current command created for generating command by openai",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "c",
						Usage: "command for help with ai",
					},
				},
				Action: func(c *cli.Context) error {
					ai.Cmd(c)
					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}
