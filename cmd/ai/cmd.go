package ai

import (
	"flag"
	"fmt"
	"github.com/urfave/cli/v2"
)

var cmdPrompt string = "Please generate console command for description %s." +
	"Response should be as in console format." +
	"Response should contains example."

func Cmd(c *cli.Context) {
	s := c.String("c")
	flag.Parse()
	if len(s) == 0 {
		fmt.Println("You must provide a command for help with ai")
		return
	}
	Tts(fmt.Sprintf(cmdPrompt, s))
}
