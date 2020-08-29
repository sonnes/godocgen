package main

import (
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/sonnes/godocgen"
	"github.com/urfave/cli/v2"
)

func generateMarkdown(cli *cli.Context) error {

	src := cli.String("src")
	mdName := cli.String("markdown")

	directories, err := godocgen.GetSourcePackages(src)

	if err != nil {
		return err
	}

	for _, d := range directories {
		log.Debug().Msgf("Package: %s", d)

		idx, err := godocgen.GetPackageIndex(d)

		if err != nil {
			return err
		}

		outPath := filepath.Join(d, mdName)
		log.Debug().Msgf("Writing to %s", outPath)
		err = godocgen.WriteMarkdown(idx, outPath)

		if err != nil {
			return err
		}
	}

	return nil
}
