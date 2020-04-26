package main

import (
	"fmt"
	"log"
	"os"

	"github.com/arithran/covercmp/compare"
	"github.com/arithran/covercmp/driver/golang"
	"github.com/urfave/cli/v2"
)

const (
	goUsage     = "covercmp go before.txt after.txt"
	goUsageText = `
Each input file should be from:
	'go test -count=1 -cover > [before,after].txt'

Covercmp compares before and after for each unit test runs.
	'covercmp go before.txt after.txt'
`
)

func main() {
	app := &cli.App{
		Name:  "covercmp",
		Usage: "The covercmp command displays code coverage changes between unit test runs",
		Commands: []*cli.Command{
			{
				Name:      "go",
				Usage:     goUsage,
				UsageText: goUsageText,
				Action: func(c *cli.Context) error {
					if c.NArg() < 2 {
						return fmt.Errorf("not enough arguments usage:\n %s", goUsage)
					}
					return compare.Cmp(golang.New(), c.Args().Get(0), c.Args().Get(1))
				},
			},

			// Add other language support in the future
			// ...
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
