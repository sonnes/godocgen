package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sonnes/godocgen/version"
	"github.com/urfave/cli/v2"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	console := zerolog.ConsoleWriter{
		Out: os.Stderr,
	}
	log.Logger = log.Output(console)

	commands := []*cli.Command{
		{
			Name:  "markdown",
			Usage: "generates markdown documentation for all packages in `src` folder",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "source",
					Aliases: []string{"src"},
					Usage:   "path to packages that have to be documented",
				},
				&cli.StringFlag{
					Name:    "markdown",
					Aliases: []string{"md"},
					Usage:   "name of markdown file to output to",
					Value:   "GODOC.md",
				},
			},
			Action: generateMarkdown,
		},
	}

	app := &cli.App{
		Name:     "godocgen",
		Usage:    "Document & maintain your Go documentation in markdown files",
		Commands: commands,
		Version:  version.Version,
	}

	app.Run(os.Args)
}
