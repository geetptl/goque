package main

import (
	"log"

	"github.com/spf13/cobra"
)

var removeCommand removeCommand_

var rmCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove a topic",
	Long:    "Remove topic. If topicName doesn't exist, it is silent.",
	Aliases: []string{"rm"},
	Example: "goque remove --topic topicName\nOR\ngoque rm -t topic",
	Args:    cobra.MinimumNArgs(0),
	RunE:    rmFunc,
}

func init() {
	removeCommand = RemoveCommand()
	rmCmd.PersistentFlags().StringVarP(&removeCommand.topic, "topic", "t", "", "Topic name to remove")
	rootCmd.AddCommand(rmCmd)
}

func rmFunc(command *cobra.Command, args []string) error {
	log.Println(removeCommand)
	topic, err := getTopic(removeCommand.topic, false)
	log.Println(topic, err)
	if err != nil {
		return err
	}

	if topic == "" {
		return nil
	}

	return removeTopic(topic)
}
