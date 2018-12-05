package main

import (
	"os"

	"github.com/yogihardi/guestbook/cli/run"
	"github.com/yogihardi/guestbook/version"

	"github.com/inconshreveable/log15"
	"github.com/urfave/cli"
)

var logHandler log15.Handler

func main() {
	app := cli.NewApp()
	app.Name = "guestbook"
	app.Usage = "Guset Book API"
	app.Version = version.Version + " (" + version.GitCommit + ")"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug, d",
			Usage:  "Enable verbose logging",
			EnvVar: "GUESTBOOK_DEBUG",
		},
	}
	app.Before = func(c *cli.Context) error {
		f := log15.JsonFormat()
		if c.Bool("debug") {
			log15.Root().SetHandler(log15.CallerStackHandler("%+v", log15.StreamHandler(os.Stdout, f)))
		} else {
			log15.Root().SetHandler(log15.MultiHandler(
				log15.LvlFilterHandler(log15.LvlInfo, log15.CallerFileHandler(log15.StreamHandler(os.Stdout, f))),
			))
		}

		return nil
	}
	app.Commands = []cli.Command{
		run.Command,
	}

	if err := app.Run(os.Args); err != nil {
		log15.Crit(err.Error())
	}
}
