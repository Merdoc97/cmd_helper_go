package ai

import (
	"flag"
	"fmt"
	"github.com/urfave/cli/v2"
)

var prompt string = "current %s can help with any command which you need to help," +
	"it's analogue of man command but with examples"

func Help(c *cli.Context) {
	s := c.String("c")
	flag.Parse()
	if len(s) == 0 {
		fmt.Println("You must provide a command for help with ai")
		return
	}
	Tts(fmt.Sprintf(prompt, s))
}
