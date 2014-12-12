package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/wm/go-flowdock/flowdock"
)

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
