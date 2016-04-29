package main

import (
	"log"
	"os"
)

/*
 * Logs configuration.
 */
func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}

func main() {
	directory := getHTTPDirectory()
	dir = &directory

	/*
	 * Creates the HTTP directory if it does not exist.
	 */
	setUpHTTPDir(directory)

	/**
	 * Starts the HTTP server
	 */
	serveFiles()
}

func getHTTPDirectory() string {
	if os.Getenv("HTTP_DIR") == "" {
		return "./"
	} else {
		return os.Getenv("HTTP_DIR")
	}
}

func setUpHTTPDir(directory string) {
	_, err := os.Stat(directory)
	if err == nil {
		log.Printf("The %s directory already exists.", directory)
		return
	}

	if !os.IsNotExist(err) {
		log.Fatalln("Error while setting up the HTTP directory", err)
		return
	}

	log.Printf("Creating the %s directory", directory)
	err = os.MkdirAll(directory, 0644)
	if err != nil {
		log.Fatalln("Error while creating the HTTP directory", err)
		return
	}
}
