package main

import (
	"fmt"
	"os"

	"github.com/hapoon/ak4/action"
	"github.com/urfave/cli/v2"
)

const (
	name    = "ak4"
	version = "0.1.0"
)

func main() {
	app := &cli.App{
		Description: "A cli appliation for AKASHI",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "profile", Usage: "profile name"},
			&cli.BoolFlag{Name: "V", Usage: "vervose"},
		},
		Commands: []*cli.Command{
			{
				Name:      "stamp",
				Usage:     "Process employee imprinting",
				UsageText: "Process employee imprinting",
				Action:    action.ActStamp,
			},
			{
				Name:      "init",
				Usage:     "Initialize setting",
				UsageText: "Initialize setting",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "login_company_code", Usage: "AKASHI企業ID"},
					&cli.StringFlag{Name: "token", Usage: "トークン"},
				},
				Action: action.ActInit,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
