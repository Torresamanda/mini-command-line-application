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

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Seach IP addresses on the internet",
			Flags:  flags,
			Action: searchIPs,
		},
		{
			Name:   "servers",
			Usage:  "Search the name of servers on the internet",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIPs(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ips := range ips {
		fmt.Println(ips)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host) // name server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servers := range servers {
		fmt.Println(servers.Host)
	}

}
