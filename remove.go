package main

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

func rmFunc(command *cobra.Command, args []string) error {
	fmt.Println(newContext)
	topic, lines, err := getTopic(newContext, false)
	fmt.Println("rm", topic, lines, err)
	if err != nil {
		return err
	}

	if topic == "" {
		return nil
	}
	return removeFileAndEntry(topic)
}

func removeFileAndEntry(topic string) error {
	fmt.Println("removing " + topic)
	err := removeTopic(topic)
	if err != nil {
		return err
	}

	err = os.Remove(path.Join(DATADIR_PATH, topic))
	if err != nil {
		return err
	}

	return nil
}

var rmCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove a topic",
	Long:    "Remove topic. If topicName doesn't exist, it is silent.",
	Aliases: []string{"rm"},
	Example: "goque remove --topic topicName\nOR\ngoque rm -t topic",
	Args:    cobra.MinimumNArgs(0),
	RunE:    rmFunc,
}
