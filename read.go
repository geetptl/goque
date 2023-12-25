package main

import (
	"log"

	"github.com/spf13/cobra"
)

var readCommand readCommand_

var readCmd = &cobra.Command{
	Use:     "read",
	Short:   "Read messages from a topic",
	Long:    "Read messages from a topic. If topicName doesn't exist, it returns empty.",
	Aliases: []string{"r"},
	Example: "goque read --topic topicName --number numberOfMessages\nOR\ngoque a -t topic -n nMessages",
	Args:    cobra.MinimumNArgs(0),
	RunE:    readFunc,
}

func init() {
	readCommand = ReadCommand()
	readCmd.PersistentFlags().StringVarP(&readCommand.topic, "topic", "t", "", "Topic name to read from")
	readCmd.PersistentFlags().IntVarP(&readCommand.number, "number", "n", 0, "Number of messages to read from a topic")
	rootCmd.AddCommand(readCmd)
}

func readFunc(command *cobra.Command, args []string) error {
	log.Println(readCommand)
	topic, err := getTopic(readCommand.topic, false)
	log.Println(topic, err)
	if err != nil {
		return err
	}

	if topic == "" {
		return nil
	}

	messages, err := readFromTopic(topic, readCommand.number)
	if err != nil {
		return err
	}
	log.Println(messages)
	return nil
}
