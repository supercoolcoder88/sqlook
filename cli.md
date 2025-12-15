1. Setup

Initialize your module and install the v3 library:
Bash

go mod init my-cli-app
go get github.com/urfave/cli/v3

2. The Basic Skeleton

In v3, the concept of cli.App has been merged into cli.Command. Your "Application" is simply the root command.

Create a main.go file:
Go

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	// In v3, the root application is defined as a cli.Command
	cmd := &cli.Command{
		Name:    "mytool",
		Usage:   "A helpful CLI tool built with v3",
		Version: "1.0.0",
		// The default action if no subcommand is provided
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println("👋 Hello! Welcome to mytool.")
			return nil
		},
	}

	// Run the command. Note that v3 requires a context.
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

3. Adding Flags

Flags allow you to pass parameters to your commands (e.g., --name John or --verbose).

Update the cmd definition to include Flags:
Go

cmd := &cli.Command{
    Name:  "greet",
    Usage: "Greet a user",
    // Define flags here
    Flags: []cli.Flag{
        &cli.StringFlag{
            Name:    "name",
            Aliases: []string{"n"},
            Value:   "World", // Default value
            Usage:   "The name of the person to greet",
        },
        &cli.BoolFlag{
            Name:    "verbose",
            Aliases: []string{"v"},
            Usage:   "Enable verbose output",
        },
    },
    // The Action function signature now accepts (ctx, cmd)
    Action: func(ctx context.Context, cmd *cli.Command) error {
        // Retrieve flag values using the cmd object
        name := cmd.String("name")
        verbose := cmd.Bool("verbose")

        if verbose {
            fmt.Println("[DEBUG] Preparing to greet...")
        }

        fmt.Printf("Hello, %s!\n", name)
        return nil
    },
}

4. Adding Subcommands

Complex applications often have subcommands (like git commit, git push). You can nest commands using the Commands field.
Go

cmd := &cli.Command{
    Name:  "math",
    Usage: "Perform math operations",
    Commands: []*cli.Command{
        {
            Name:    "add",
            Aliases: []string{"a"},
            Usage:   "Add two numbers",
            Action: func(ctx context.Context, cmd *cli.Command) error {
                // cmd.Args() contains positional arguments
                if cmd.NArg() < 2 {
                    return fmt.Errorf("please provide two numbers")
                }
                fmt.Println("Result:", cmd.Args().Get(0) + " + " + cmd.Args().Get(1))
                return nil
            },
        },
        {
            Name:  "subtract",
            Usage: "Subtract two numbers",
            Action: func(ctx context.Context, cmd *cli.Command) error {
                fmt.Println("Subtraction feature coming soon!")
                return nil
            },
        },
    },
}