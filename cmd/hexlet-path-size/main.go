package main

import (
	"code/internal"
	"context"
	"errors"
	"fmt"
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

			extraArguments := cmd.Args().Len()
			if extraArguments > 0 {
				return fmt.Errorf("expecting one argument (path), got %d more", extraArguments)
			}

			if path == "" {
				return errors.New("missing path")
			}

			size, err := internal.GetPathSize(path, isRecursive, includeHidden)

			if err != nil {
				return err
			}

			fmt.Println(internal.FormatCLIOutput(path, size, isHumanReadable))
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
