package main

import (
	"github.com/mesogii/svld/core"
	"github.com/mesogii/svld/util"

	"fmt"
	"os"

	"github.com/urfave/cli"
)

const version = "0.1.6"

const usage = `Itty Bitty encryption utility.

	Made to quickly secure secretive files within repos.

	Built for teams sharing many sensitive files across projects.
`


func main() {

	app := cli.NewApp()
	app.Name = "Svalbard"
	app.Usage = usage
	app.Version = version
	app.Commands = commands()

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("svld help - for usage options")
		return nil
	}


	app.Run(os.Args)

}

func commands() []cli.Command {
	return []cli.Command{

		{
			Name:    "encrypt",
			Aliases: []string{"e"},
			Usage:   "encrypt a file",
			Action:  func(c *cli.Context) {

				filepath := ""
				if c.NArg() > 0 {
					filepath = c.Args().Get(0)
				}

				err := core.EncryptIntf(filepath)

				if err != nil {
					util.Error(err)
				}
			},
		},
		{
			Name:    "decrypt",
			Aliases: []string{"d"},
			Usage:   "decrypt a file",
			Action:  func(c *cli.Context){
				filepath := ""
				if c.NArg() > 0 {
					filepath = c.Args().Get(0)
				}

				err := core.DecryptIntf(filepath)

				if err != nil {
					util.Error(err)
				}
			},
		},

		{
			Name:        "key",
			Aliases:     []string{"k"},
			Usage:       "options for svalbard keys",
			Subcommands: []cli.Command{
				{
					Name:  "gen",
					Usage: "generate a new svalbard key",
					Action: func(c *cli.Context) {

						err := core.Key_Gen_Intf(c.Args().First())

						if err != nil {
							util.Error(err)
						}
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
		},

	}
}
