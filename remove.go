package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func rmFunc(command *cobra.Command, args []string) error {
	fmt.Println(newContext)
	topic, lines, err := getTopic(newContext, true)
	fmt.Println(topic, lines, err)
	return nil
}

var rmCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove a topic",
	Long:    "Remove topic. If topicName doesn't exist, it is silent.",
	Aliases: []string{"rm"},
	Example: "goque remove --topic topicName\nOR\ngoque rm -t topic",
	Args:    cobra.MinimumNArgs(0),
	RunE:    addFunc,
}
