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
			isHumanFormat := cmd.Bool("human")
			includeHidden := cmd.Bool("all")
			isRecursive := cmd.Bool("recursive")

			size, err := code.GetPathSize(path, includeHidden, isRecursive)

			if err != nil {
				return err
			}

			fmt.Printf("%s\t%s\n", formatFileSize(size, isHumanFormat), path)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

type SizeUnitEntry struct {
	size    int64
	postfix string
}

var sizeUnits = []SizeUnitEntry{
	{size: int64(1) << 0, postfix: "B"},
	{size: int64(1) << 10, postfix: "KB"},
	{size: int64(1) << 20, postfix: "MB"},
	{size: int64(1) << 30, postfix: "GB"},
	{size: int64(1) << 40, postfix: "TB"},
	{size: int64(1) << 50, postfix: "PB"},
	{size: int64(1) << 60, postfix: "EB"},
}

func formatFileSize(size int64, isHuman bool) string {
	if !isHuman {
		return fmt.Sprintf("%dB", size)
	}

	unit := sizeUnits[0]

	for _, u := range sizeUnits {
		if size < u.size {
			break
		}
		unit = u
	}

	sizeInUnit := float32(size) / float32(unit.size)

	return fmt.Sprintf("%.1f%s", sizeInUnit, unit.postfix)
}
