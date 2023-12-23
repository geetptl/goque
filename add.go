package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func addFunc(command *cobra.Command, args []string) error {
	fmt.Println(newContext)
	return nil
}

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a message to a topic",
	Long:    "Add a message to a topic. If topicName doesn't exist, it is created.",
	Aliases: []string{"a"},
	Example: "goque add --topic topicName --msg messageContent\nOR\ngoque a -t topic -m message",
	Args:    cobra.MinimumNArgs(0),
	RunE:    addFunc,
}
