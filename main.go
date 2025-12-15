package main

import (
	"context"
	"log"
	"os"
	"sqlook/commandOrchestrator"

	"github.com/urfave/cli/v3"
)

func main() {
	// CLI command
	cmd := &cli.Command{
		Name:    "sqlook",
		Usage:   "Testing utility to look into SQLite databases locally",
		Version: "0.1.0",
		Action:  commandOrchestrator.CommandOrchestrator,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "query",
				Aliases: []string{"q"},
				Usage:   "SQLite query",
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
