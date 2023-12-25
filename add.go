package main

import (
	"log"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a message to a topic",
	Long:    "Add a message to a topic. If topicName doesn't exist, it is created.",
	Aliases: []string{"a"},
	Example: "goque add --topic topicName --msg messageContent\nOR\ngoque a -t topic -m message",
	Args:    cobra.MinimumNArgs(0),
	RunE:    addFunc,
}

var addCommand addCommand_

func init() {
	addCommand = AddCommand()
	addCmd.PersistentFlags().StringVarP(&addCommand.topic, "topic", "t", "", "Topic name for your message")
	addCmd.PersistentFlags().StringVarP(&addCommand.message, "message", "m", "", "A message to put into topic")
	rootCmd.AddCommand(addCmd)
}

func addFunc(command *cobra.Command, args []string) error {
	log.Println(addCommand)
	topic, err := getTopic(addCommand.topic, true)
	log.Println(topic, err)
	if err != nil {
		return err
	}

	return addMessageInTopic(topic, addCommand.message)
}
