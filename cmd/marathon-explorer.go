package main

import (
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
	"log"
	"marathon-explorer/internal/command"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "Marathon Explorer",
		Usage: "command line tool for getting information from a Mesos+Marathon cluster",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Aliases:  []string{"u"},
				Usage:    "marathon url",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "user",
				Aliases: []string{"U"},
				Usage:   "marathon basic auth user",
			},
			&cli.StringFlag{
				Name:    "password",
				Aliases: []string{"p"},
				Usage:   "marathon basic auth password",
			},
			&cli.StringSliceFlag{
				Name:    "project",
				Aliases: []string{"P"},
				Usage:   "filter information by project",
			},
			&cli.IntFlag{
				Name:    "instances",
				Aliases: []string{"i"},
				Usage:   "filter information by instances count",
			},
			&cli.StringFlag{
				Name:    "image",
				Aliases: []string{"I"},
				Usage:   "filter information by image substring",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "applications",
				Usage: "get full info about applications",
				Action: func(c *cli.Context) error {
					result := command.GetApplications(c)
					printResult(result)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func printResult(result [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(result[0])

	for _, r := range result[1:] {
		table.Append(r)
	}
	table.Render()
}
