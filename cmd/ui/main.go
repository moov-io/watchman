package main

import (
	"cmp"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/moov-io/watchman"
	"github.com/moov-io/watchman/internal/ui"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/moov-io/base/log"
	"github.com/urfave/cli/v2"
)

var (
	flagBaseAddress = &cli.StringFlag{
		Name:  "address",
		Value: cmp.Or(os.Getenv("WATCHMAN_ADDRESS"), "http://localhost:8084"),
		Usage: "Address to connect with Watchman",
	}
	flagVerbose = &cli.BoolFlag{
		Name:  "verbose",
		Value: false,
		Usage: "Output verbose logging",
	}
)

func main() {
	logger := log.NewDefaultLogger().With(log.Fields{
		"app":     log.String("watchman"),
		"version": log.String(watchman.Version),
	})
	logger.Log("Starting watchman UI")

	ctx := context.Background()

	app := &cli.App{
		Name: "watchman-ui",
		// UsageText:   "watchman-ui [global options] command [command options]",
		Description: "Watchman GUI",
		Authors: []*cli.Author{
			{Name: "Moov OSS", Email: "oss@moov.io"},
		},
		Flags: []cli.Flag{
			// Common Flags
			flagBaseAddress, flagVerbose,
		},
		Commands: []*cli.Command{
			// commandFind,
		},
		Action: func(cliContext *cli.Context) error {
			env := ui.Environment{
				Logger: logger,
				Client: createWatchmanClient(cliContext.String(flagBaseAddress.Name)),
			}

			// cli.ShowAppHelp(ctx)

			return showUI(ctx, env)
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("ERROR running command: %v\n", err)
		os.Exit(127)
	}
}

func showUI(ctx context.Context, env ui.Environment) error {
	app := ui.New(ctx, env)
	app.Run()

	return nil
}

func createWatchmanClient(baseAddress string) search.Client {
	var httpClient *http.Client = nil

	return search.NewClient(httpClient, baseAddress)
}
