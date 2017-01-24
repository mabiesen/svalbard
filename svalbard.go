package main

import (
	"github.com/mesogii/svalbard/core"
	"github.com/mesogii/svalbard/util"

	"os"

	"github.com/urfave/cli"
)

const version = "0.1.6"

const usage = `Itty Bitty encryption utility.

	Made to quickly secure secretive files within repos.

	Built for teams sharing many sensitive files across projects.
`


func main() {

	// Initialize a new CLI app
	app := cli.NewApp()
	app.Name = "Svalbard"
	app.Usage = usage
	app.Version = version
	app.Commands = commands()

	// Define default action
	app.Action = func(c *cli.Context) error {
		util.Notice("svld help - for usage options")
		return nil
	}


	app.Run(os.Args)

}


// CLI Command Objects
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
					Name:  "list",
					Usage: "list all key files",
					Action: func(c *cli.Context) {

						err := core.KeyListIntf()

						if err != nil {
							util.Error(err)
						}
					},
				},
				{
					Name:  "gen",
					Usage: "generate a new svalbard key",
					Action: func(c *cli.Context) {

						err := core.KeyGenIntf(c.Args().First())

						if err != nil {
							util.Error(err)
						}
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing key file",
					Action: func(c *cli.Context) {
						err := core.KeyRemoveIntf(c.Args().First())

						if err != nil {
							util.Error(err)
						}
					},
				},
			},
		},

	}
}
