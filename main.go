package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// cli Flags
	filename := flag.String("filename", "", "Filename of the csv file you want to convert")
	newFilename := flag.String("newname", *filename+"Converted", "Name of the converted file")
	help := flag.Bool("help", false, "Show CLI options")
	flag.Parse()

	// show help command if user had '-help' as one of the options
	if *help {
		fmt.Printf("CSV semicolon to comma converter usage: \n 'go run . [option] [argument] ... [option] [argument]' \n \n Options:\n -filename , takes the name of the file as argument with '.csv' suffix. This is a necessary option! \n -newname , takes the given string as the name of the generated file. If new name was not given, then generates the new filename based on original filename. This is optional \n \n Example for file 'examplefile.csv': \n $ go run . -filename examplefile \n Generates a new file called 'examplefileConverted.csv' in the current directory\n \n Suffixes in filenames can be used, but are not required.\n")
		os.Exit(0)
	}

	// add suffixes if they were not given
	if !strings.HasSuffix(*filename, ".csv") {
		*filename = *filename + ".csv"
	}
	if !strings.HasSuffix(*newFilename, ".csv") {
		*newFilename = *newFilename + ".csv"
	}

	// if empty, throw error and prompt user to use flag '-filename'
	if *filename == "" {
		log.Fatal("Filename is empty. Please input filename to be converted after option '-filename'.")
	}

	// read all the data in file to data
	data, err := os.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}

	// convert []byte to string
	str := string(data)

	// create new string with semicolons replaced with commas
	strNew := strings.ReplaceAll(str, ";", ",")

	// write the string in []byte form to new file
	err = os.WriteFile(*newFilename, []byte(strNew), 0777)
}
