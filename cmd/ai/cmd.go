package ai

import (
	"flag"
	"fmt"
	"github.com/urfave/cli/v2"
)

var cmdPrompt string = "Please generate command for description %s. \r\n" +
	"Response should be as in console format. \r\n" +
	"Response should contains example. \r\n"

func Cmd(c *cli.Context) {
	s := c.String("c")
	flag.Parse()
	if len(s) == 0 {
		fmt.Println("You must provide a command for help with ai")
		return
	}
	Tts(fmt.Sprintf(cmdPrompt, s))
}
