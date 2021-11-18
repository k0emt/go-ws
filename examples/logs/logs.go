package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("default log output goes to StdOut")
	setupLogging()
	log.Println("This should go to the log file")
}

func setupLogging() {
	logFile, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("something went wrong setting up the application log")
		log.Print(err)
		return
	}
	log.SetOutput(logFile)

	// set up an error
	config, err := readConfig("/path/to/config.file")
	if err != nil {
		log.Printf("error: %+v", err)
		os.Exit(1)
	}
	fmt.Println(config)
}

type Config struct {
	appName string
} // application settings

func readConfig(path string) (*Config, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return nil, err // must use *Config to return nil
	}
	defer configFile.Close()

	config := &Config{}
	// parse file here
	return config, nil
}
