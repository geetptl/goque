package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path"
	"strconv"
)

/*
The new plan:
No index file.
To list all topics, just list all the directories in DATADIR
DATADIR/topic should contain two files: index & log
*/

/*
if topic is found, return (topic, nil)
if not found and create is true, create directory and return (topic, nil)
if not found and create is false, return ("", nil)
if any error, return ("", err)
*/
func getTopic(topic string, create bool) (string, error) {
	// check for topic dir. check for subfiles. if anything is missing, throw error.
	topicStat, err := os.Stat(path.Join(DATADIR_PATH, topic))
	log.Println(topicStat, err)
	if err != nil {
		if os.IsNotExist(err) {
			if create {
				return createTopic(topic)
			} else {
				return "", nil
			}
		}
		return "", err
	}
	if !topicStat.IsDir() {
		return "", errors.New("Malformed topic directory")
	}

	// topic is good, check subfiles.
	indexStat, err := os.Stat(path.Join(DATADIR_PATH, topic, "index"))
	log.Println(indexStat, err)
	if err != nil {
		return "", err
	}
	if indexStat.IsDir() {
		return "", errors.New(topic + "/index is a directory")
	}

	logStat, err := os.Stat(path.Join(DATADIR_PATH, topic, "log"))
	log.Println(logStat, err)
	if err != nil {
		return "", err
	}
	if logStat.IsDir() {
		return "", errors.New(topic + "/log is a directory")
	}

	// all good here, I guess
	return topic, nil
}

func createTopic(topic string) (string, error) {
	err := os.Mkdir(path.Join(DATADIR_PATH, topic), os.ModePerm)
	log.Println(err)
	if err != nil {
		return "", err
	}

	indexFile_W, err := os.OpenFile(path.Join(DATADIR_PATH, topic, "index"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	log.Println(err)
	if err != nil {
		return "", err
	}
	defer indexFile_W.Close()
	_, err = indexFile_W.WriteString("0")
	if err != nil {
		return "", err
	}

	logFile_W, err := os.OpenFile(path.Join(DATADIR_PATH, topic, "log"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	log.Println(err)
	if err != nil {
		return "", err
	}
	defer logFile_W.Close()

	return topic, nil
}

/*
read from topic/log, and update topic/index.
*/
func readFromTopic(topic string, number int) ([]string, error) {
	if number < 0 {
		return nil, errors.New("number < 0 not allowed")
	}

	indexFile, err := os.OpenFile(path.Join(DATADIR_PATH, topic, "index"), os.O_RDONLY, os.ModePerm)
	log.Println(err)
	if err != nil {
		return nil, err
	}
	defer indexFile.Close()
	indexbuf := [64]byte{}
	n, err := indexFile.Read(indexbuf[:])
	log.Println(err)
	if err != nil {
		return nil, err
	}
	lineStart, err := strconv.Atoi(string(indexbuf[0:n]))
	log.Println(err)
	if err != nil {
		return nil, err
	}

	logFile, err := os.OpenFile(path.Join(DATADIR_PATH, topic, "log"), os.O_RDONLY, os.ModePerm)
	log.Println(err)
	if err != nil {
		return nil, err
	}
	defer logFile.Close()

	messages := []string{}
	scanner := bufio.NewScanner(logFile)
	linesScanned := 0
	linesRead := 0
	log.Println(lineStart)
	for scanner.Scan() {
		if linesScanned < lineStart {
			linesScanned++
		} else {
			if linesScanned < n+lineStart {
				messages = append(messages, scanner.Text())
				linesScanned++
				linesRead++
			} else {
				break
			}
		}
	}

	indexFile_W, err := os.OpenFile(path.Join(DATADIR_PATH, topic, "index"), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	log.Println(err)
	if err != nil {
		return nil, err
	}
	_, err = indexFile_W.WriteString(strconv.Itoa(lineStart + linesRead))
	if err != nil {
		return nil, err
	}
	return messages, nil
}

/* append message to topic/log */
func addMessageInTopic(topic string, message string) error {
	logFile_W, err := os.OpenFile(path.Join(DATADIR_PATH, topic, "log"), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	log.Println(err)
	if err != nil {
		return err
	}
	defer logFile_W.Close()

	n, err := logFile_W.WriteString(message + "\n")
	log.Println(err)
	log.Println(strconv.Itoa(n) + " bytes written")
	if err != nil {
		return err
	}
	return nil
}

/* remove topic directory */
func removeTopic(topic string) error {
	err := os.RemoveAll(path.Join(DATADIR_PATH, topic))
	log.Println(err)
	if err != nil {
		return err
	}
	return nil
}
