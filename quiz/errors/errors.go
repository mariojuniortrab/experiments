package errors

import (
	"fmt"
	"os"
)

type IError interface {
	Exit()
}

type FileNotFoundError struct {
	FileName string
}

func (e *FileNotFoundError) Exit() {
	fmt.Printf("Failed to Open file: %s\n", e.FileName)
	os.Exit(1)
}

func ShowErrorAndExit(error IError) {
	error.Exit()
}
