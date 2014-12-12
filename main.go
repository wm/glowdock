package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/wm/go-flowdock/flowdock"
	"github.com/wm/godock/auth"
)

var (
	users        map[int]flowdock.User
	clientSecret string
	clientId     string
	client       *flowdock.Client
)

func init() {
	client = flowdock.NewClient(auth.AuthenticationRequest(clientSecret, clientId))
	users, _ = fetchUsers()
}

func main() {
	app := cli.NewApp()
	app.Name = "godock"
	app.Usage = "Interact with your flowdock flows."

	app.Commands = []cli.Command{
		{
			Name:   "stream",
			Usage:  "Stream messages from a flow.",
			Action: Stream,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "org, o",
					Usage: "The organization name of the flow(s).",
				},
				cli.StringFlag{
					Name:  "flow, f",
					Usage: "The flow(s) to listen to - separated with spaces.",
				},
				cli.StringFlag{
					Name:  "events, e",
					Value: "comment, message",
					Usage: `The events to listen to - separated with spaces.
                    Events: message status comment action tag-change message-edit
                    activity.user file discussion user-edit file mail zendesk
                    twitter`,
				},
			},
		},
		{
			Name:   "message",
			Usage:  "Send a message to a Flow.",
			Action: Message,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "flow, f",
					Usage: "The flow to write to.",
				},
				cli.StringFlag{
					Name:  "prepend, p",
					Usage: "Prepend text to message.",
				},
			},
		},
		{
			Name:   "flows",
			Usage:  "List the flows.",
			Action: ListFlows,
		},
	}

	app.Run(os.Args)
}

func fetchUsers() (map[int]flowdock.User, error) {
	users := map[int]flowdock.User{}
	userList, _, err := client.Users.List()
	if err == nil {
		for _, user := range *userList {
			users[*user.ID] = user
		}
	}

	return users, err
}
