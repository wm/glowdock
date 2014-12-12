package main

import (
	"bytes"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/wm/go-flowdock/flowdock"
)

func Message(c *cli.Context) {
	var content bytes.Buffer

	flow := c.String("flow")
	if len(flow) < 1 {
		cli.ShowCommandHelp(c, c.Command.Name)
		os.Exit(1)
	}

	prefix := c.String("prepend")
	if len(prefix) > 0 {
		content.WriteString(prefix)
		content.WriteString(" ")
	}

	if len(c.Args()) > 0 {
		content.WriteString(strings.Join(c.Args(), " "))
	} else {
		cli.ShowCommandHelp(c, c.Command.Name)
		os.Exit(1)
	}

	opt := &flowdock.MessagesCreateOptions{
		FlowID:  flow,
		Event:   "message",
		Content: content.String(),
	}
	_, _, err := client.Messages.Create(opt)
	if err != nil {
		log.Fatal("Get:", err)
	}

	return
}
