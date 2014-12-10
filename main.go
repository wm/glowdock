package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/wm/go-flowdock/flowdock"
	"github.com/wm/godock/auth"
)

var (
	clientSecret string
	clientId     string
	client       *flowdock.Client
)

func init() {
	client = flowdock.NewClient(auth.AuthenticationRequest(clientSecret, clientId))
}

func main() {
	app := cli.NewApp()
	app.Name = "godock"
	app.Usage = "Interact with your flowdock flows"

	app.Commands = []cli.Command{
		{
			Name:   "message",
			Usage:  "Send messages to a Flow",
			Action: Message,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "flow, f",
					Usage: "the flow to write to",
				},
				cli.StringFlag{
					Name:  "prepend, p",
					Value: "",
					Usage: "Prepend text to message",
				},
			},
		},
		{
			Name:   "flows",
			Usage:  "List the flows",
			Action: ListFlows,
		},
	}

	app.Run(os.Args)
}

func ListFlows(c *cli.Context) {
	opt := flowdock.FlowsListOptions{User: true}
	flows, _, err := client.Flows.List(true, &opt)

	if err != nil {
		log.Fatal("Get:", err)
	}

	fmt.Printf("%35s\t| %20s\t| %s\t| %s\n\n", "ID", "Name", "Organization", "URL")
	for _, flow := range flows {
		displayFlowData(flow)
	}
}

func displayFlowData(flow flowdock.Flow) {
	org := flow.Organization
	fmt.Printf("%35s\t| %20s\t| %s\t| %s\n", *flow.Id, *flow.Name, *org.Name, *flow.Url)
}

func Message(c *cli.Context) {
	var content bytes.Buffer

	flow := c.String("flow")
	if len(flow) < 1 {
		log.Fatal("A flow is required")
		return
	}

	prefix := c.String("prepend")
	if len(prefix) > 0 {
		content.WriteString(prefix)
		content.WriteString(" ")
	}

	if len(c.Args()) > 0 {
		content.WriteString(strings.Join(c.Args(), " "))
	} else {
		log.Fatal("A message is required")
		return
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
