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

var newContext context

var GOQUE_PATH string
var INDEX_PATH string
var DATADIR_PATH string

func init() {
	GOQUE_PATH = path.Join(os.Getenv("HOME"), ".goque")
	INDEX_PATH = path.Join(GOQUE_PATH, "index")
	DATADIR_PATH = path.Join(GOQUE_PATH, "data/")

	newContext = NewContext()

	addCmd.PersistentFlags().StringVarP(&newContext.topic, "topic", "t", "", "Topic name for your message")
	addCmd.PersistentFlags().StringVarP(&newContext.message, "message", "m", "", "A message to put into topic")

	readCmd.PersistentFlags().StringVarP(&newContext.topic, "topic", "t", "", "Topic name to read from")
	readCmd.PersistentFlags().IntVarP(&newContext.numberOfMessages, "number", "n", 0, "Number of messages to read from a topic")

	rmCmd.PersistentFlags().StringVarP(&newContext.topic, "topic", "t", "", "Topic name to remove")

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
