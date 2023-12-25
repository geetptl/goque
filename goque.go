package main

import (
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func main() {
	if err := Execute(); err != nil {
		os.Exit(1)
	}
}
