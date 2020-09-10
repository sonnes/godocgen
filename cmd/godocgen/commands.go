package main

import (
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/sonnes/godocgen"
	"github.com/urfave/cli/v2"
)

func generateMarkdown(c *cli.Context) error {

	src := c.String("src")
	name := c.String("name")

	if src == "" {
		cli.ShowAppHelp(c)
		cli.ShowVersion(c)
		return nil
	}

	log.Info().Msgf("Version: %s", c.App.Version)
	log.Debug().Msgf("Package: %s", src)

	idx, err := godocgen.GetPackageIndex(src)

	if err != nil {
		return err
	}

	outPath := filepath.Join(src, name)

	log.Debug().Msgf("Writing to %s", outPath)
	err = godocgen.WriteMDTemplate(idx, outPath)

	if err != nil {
		return err
	}

	return nil
}
