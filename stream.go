package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/wm/glowdock/auth"
	"github.com/wm/go-flowdock/flowdock"
)

type flowMessage struct {
	Flow    string
	Message flowdock.Message
}

type flowStream struct {
	Flow    string
	MsgChan chan flowdock.Message
}

func Stream(c *cli.Context) {
	events := strings.Split(c.String("events"), " ")

	org := c.String("org")
	if len(org) < 1 {
		cli.ShowCommandHelp(c, c.Command.Name)
		os.Exit(1)
	}

	flow := c.String("flow")
	if len(flow) < 1 {
		cli.ShowCommandHelp(c, c.Command.Name)
		os.Exit(1)
	}
	flows := strings.Split(flow, " ")

	streamFlowMessages(org, flows, events)
}

func streamFlowMessages(org string, flows []string, events []string) {
	var flowMsg flowMessage
	var streams []flowStream
	stream := flowStream{}
	flowMsgChan := make(chan flowMessage)

	for _, flow := range flows {
		msgChan, es, _ := client.Messages.Stream(auth.Token.AccessToken, org, flow)
		defer es.Close()
		stream.Flow = flow
		stream.MsgChan = msgChan
		streams = append(streams, stream)
	}

	fanIn(streams, flowMsgChan)

	for {
		flowMsg = <-flowMsgChan
		displayMessageData(flowMsg.Message, flowMsg.Flow, events)
	}
}

func fanIn(msgChannels []flowStream, out chan flowMessage) {
	var flowMsg flowMessage
	for _, ch := range msgChannels {
		go func(in flowStream) {
			for message := range in.MsgChan {
				flowMsg.Message = message
				flowMsg.Flow = in.Flow
				out <- flowMsg
			}
		}(ch)
	}
}

func displayMessageData(msg flowdock.Message, room string, events []string) {
	if stringInSlice(*msg.Event, events) {
		id, _ := strconv.ParseInt(*msg.UserID, 10, 32)
		user := users[int(id)]
		fmt.Println("\nMSG:", room, *msg.ID, *user.Nick, *msg.Event, msg.Content())
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
