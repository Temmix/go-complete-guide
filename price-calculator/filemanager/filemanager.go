package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"

	"fmt"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (f FileManager) Readfile() (*bufio.Scanner, error) {
	file, err := os.Open(f.InputFilePath)

	if err != nil {
		fmt.Println("Could not open the file")
		fmt.Println(err)
		return nil, errors.New("could not open file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	return scanner, nil
}

func (f FileManager) WriteJSON(data interface{}) error {
	file, err := os.Create(f.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	defer file.Close()

	time.Sleep(3 * time.Second) // added this for concurrency exercise
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed to convert data to JSON")
	}
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
