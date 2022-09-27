package main

import (
	"experiments/quiz/errors"
	"flag"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "A csv file in the format of 'questions,answer'.")
	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		error := errors.FileNotFoundError{FileName: *csvFileName}
		errors.ShowErrorAndExit(&error)
	}

	_ = file
}
