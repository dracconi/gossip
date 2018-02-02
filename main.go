package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/dracconi/gossip/client"
	"github.com/dracconi/gossip/server"
	"github.com/urfave/cli"
)

// Checks whether user passed IP is real
func sanitizeIP(ip string) (bool, error) {
	return regexp.MatchString("((\\d{1,3}\\.){3}\\d{1,3}\\:\\d{1,5})", ip)
}

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
					sanip, err := sanitizeIP(c.Args()[0])
					if err != nil {
						panic(err)
					}
					if sanip {
						client.Client(c.Args()[0])
					} else {
						fmt.Print("Please specify 1-255.1-255.1-255.1-255:1-65535")
					}
				} else {
					fmt.Print("Please specify IP:PORT")
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
