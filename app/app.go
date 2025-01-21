package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// The Generate functions returns the command line application ready to run
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search IPs and server names on the internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Seach IP addresses on the internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: buscarIPs,
		},
	}

	return app
}

func buscarIPs(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ips := range ips {
		fmt.Println(ips)
	}
}
