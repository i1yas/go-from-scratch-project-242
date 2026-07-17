package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory",
		ArgsUsage: "<path>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "recursive",
				Usage:       "recursive size of directories",
				Aliases:     []string{"r"},
				Value:       false,
				DefaultText: "false",
			},
			&cli.BoolFlag{
				Name:        "human",
				Usage:       "human-readable sizes (auto-select unit)",
				Aliases:     []string{"H"},
				Value:       false,
				DefaultText: "false",
			},
			&cli.BoolFlag{
				Name:        "all",
				Usage:       "include hidden files and directories",
				Aliases:     []string{"a"},
				Value:       false,
				DefaultText: "false",
			},
		},
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name: "path",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.StringArg("path")
			isHumanReadable := cmd.Bool("human")
			includeHidden := cmd.Bool("all")
			isRecursive := cmd.Bool("recursive")

			size, err := code.GetPathSize(path, isRecursive, isHumanReadable, includeHidden)

			if err != nil {
				return err
			}

			fmt.Printf("%s\t%s\n", size, path)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
