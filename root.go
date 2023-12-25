package main

import (
	"os"
	"path"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "goque",
	Args: cobra.MinimumNArgs(1),
}

var GOQUE_PATH string
var INDEX_PATH string
var DATADIR_PATH string

func init() {
	GOQUE_PATH = path.Join(os.Getenv("HOME"), ".goque")
	INDEX_PATH = path.Join(GOQUE_PATH, "index")
	DATADIR_PATH = path.Join(GOQUE_PATH, "data/")

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(readCmd)
	rootCmd.AddCommand(rmCmd)
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		rootCmd.PrintErrf("Error: %v\nRun `goque help` for help.\n", err)
		return err
	}
	return nil
}
