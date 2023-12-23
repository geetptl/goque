package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

// Returns: path relative to DATADIR, linesRead, error
// Not returning file descriptors,
// because we need them with different permissions every now and then
func getTopic(context context, create bool) (string, int, error) {
	// create indexFile if not found, keep in READONLY mode
	indexFile, err := os.OpenFile(INDEX_PATH, os.O_RDONLY, 0644)
	if errors.Is(err, os.ErrNotExist) {
		indexFile, err = os.Create(INDEX_PATH)
		if err != nil {
			return "", -1, err
		}
		indexFile, _ = os.OpenFile(INDEX_PATH, os.O_RDONLY, 0644)
	} else if err != nil {
		return "", -1, err
	}

	// create values for topicExists and topicReadCount based on indexFile
	scanner := bufio.NewScanner(indexFile)
	topicExists := false
	topicLinesRead := -1
	for scanner.Scan() {
		line := scanner.Text()
		slices := strings.Split(line, " ")
		fmt.Println(slices)
		if slices[0] == context.topic {
			topicExists = true
			topicLinesRead, err = strconv.Atoi(slices[1])
			if err != nil {
				return "", -1, err
			}
		}
	}

	// work with indexFile is done here.
	// Only thing is to append if creating a datafile,
	// for which we'll need a O_APPEND|O_WRONLY file descriptor anyways,
	// so no harm in closing this file descriptor here.
	indexFile.Close()

	// if: topic exists, then check for the data file. It must exist, call it consistancy.
	// else:
	//     if: create flag is true, then create the data file, otherwise just go home
	if topicExists {
		dataFile, err := os.OpenFile(path.Join(DATADIR_PATH, context.topic), os.O_RDONLY, 0644)
		if err != nil {
			return "", -1, err
		}
		dataFile.Close()
		return context.topic, topicLinesRead, nil
	} else if create {
		dataFile, err := os.Create(path.Join(DATADIR_PATH, context.topic))
		if err != nil {
			return "", -1, err
		}
		dataFile.Close()

		indexFile_W, _ := os.OpenFile(INDEX_PATH, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		_, err = indexFile_W.WriteString(context.topic + " 0\n")
		if err != nil {
			return "", -1, err
		}

		return context.topic, 0, nil
	} else {
		// didn't ask to create, and couldn't find
		return "", -1, nil
	}
}