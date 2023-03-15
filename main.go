package main

import (
	"fmt"
	"os"

	"github.com/hapoon/ak4/action"
	"github.com/urfave/cli/v2"
)

const (
	name    = "ak4"
	version = "0.2.0"
)

func main() {
	app := &cli.App{
		Version:     version,
		Description: "A cli appliation for AKASHI",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "profile", Usage: "profile name"},
			&cli.BoolFlag{Name: "V", Usage: "vervose"},
		},
		Commands: []*cli.Command{
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
			{
				Name:      "stamp",
				Usage:     "Process employee imprinting",
				UsageText: "Process employee imprinting",
				Action:    action.ActStamp,
			},
			{
				Name:      "token",
				Usage:     "Reissue token",
				UsageText: "Reissue token",
				Action:    action.ActToken,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
