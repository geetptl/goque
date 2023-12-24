package main

import (
	"bufio"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

func readFunc(command *cobra.Command, args []string) error {
	fmt.Println(newContext)
	topic, lines, err := getTopic(newContext, false)
	fmt.Println("read", topic, lines, err)
	if err != nil {
		return err
	}

	data, err := readNLines(topic, lines, newContext.numberOfMessages)
	if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}

func readNLines(topic string, lineStart int, n int) (string, error) {
	dataFile, err := os.OpenFile(path.Join(DATADIR_PATH, topic), os.O_RDONLY, 0644)
	if err != nil {
		return "", err
	}
	data := ""
	scanner := bufio.NewScanner(dataFile)
	linesScanned := 0
	linesRead := 0
	for scanner.Scan() {
		if linesScanned < lineStart {
			linesScanned++
			continue
		} else {
			if linesScanned < n {
				data += scanner.Text()
				linesScanned++
				linesRead++
			} else {
				break
			}
		}
	}
	updateIndex(topic, lineStart+linesRead)

	return data, nil
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
