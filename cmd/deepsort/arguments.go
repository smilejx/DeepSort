package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/akamensky/argparse"
)

func argumentParsing(args []string, argument *Arguments) {
	// Create new parser object
	parser := argparse.NewParser("deepsort", "AI powered image tagger backed by DeepDetect")
	// Create string flag
	URL := parser.String("u", "url", &argparse.Options{Required: true, Help: "URL of your DeepDetect instance (i.e: http://localhost:8080)"})
	input := parser.String("i", "input", &argparse.Options{Required: true, Help: "Your input folder."})
	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}
	// Handle the input flag
	inputFolder, _ := filepath.Abs(*input)
	// Finally print the collected string
	argument.Input = inputFolder
	argument.URL = *URL
}