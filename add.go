package main

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

func addFunc(command *cobra.Command, args []string) error {
	fmt.Println(newContext)
	topic, lines, err := getTopic(newContext, true)
	fmt.Println("add", topic, lines, err)
	if err != nil {
		return err
	}

	return writeToTopicFile(topic, newContext.message)
}

func writeToTopicFile(topic string, message string) error {
	topicFile, _ := os.OpenFile(path.Join(DATADIR_PATH, topic), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer topicFile.Close()
	_, err := topicFile.WriteString(message + "\n")
	return err
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
