package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func readFunc(command *cobra.Command, args []string) error {
	fmt.Println(newContext)
	topic, lines, err := getTopic(newContext, true)
	fmt.Println(topic, lines, err)
	return nil
}

var readCmd = &cobra.Command{
	Use:     "read",
	Short:   "Read messages from a topic",
	Long:    "Read messages from a topic. If topicName doesn't exist, it returns empty.",
	Aliases: []string{"r"},
	Example: "goque read --topic topicName --number numberOfMessages\nOR\ngoque a -t topic -n nMessages",
	Args:    cobra.MinimumNArgs(0),
	RunE:    readFunc,
}
