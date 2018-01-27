package main

import (
	"fmt"
	"os"

	"github.com/dracconi/gossip/client"
	"github.com/dracconi/gossip/server"
	"github.com/urfave/cli"
)

func main() {

	var port string

	app := cli.NewApp()

	app.Name = "gossip"
	app.Usage = "peer-to-peer chat"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "port, p",
			Value:       "1122",
			Usage:       "Port you want to listen on",
			Destination: &port,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Run the server",
			Action: func(c *cli.Context) error {
				server.Server(":" + port)
				return nil
			},
		},
		{
			Name:    "client",
			Aliases: []string{"c"},
			Usage:   "Run the client",
			Action: func(c *cli.Context) error {
				if c.NArg() > 0 {
					// fmt.Print(c.Args()[0])
					client.Client(c.Args()[0])
				} else {
					fmt.Print("Please specify IP")
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
