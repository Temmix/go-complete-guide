package cmdmanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type CMDManager struct{}

func (cmd CMDManager) Readfile() (*bufio.Scanner, error) {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("Could not open the file")
		fmt.Println(err)
		file.Close()
		return nil, errors.New("could not open file")
	}
	scanner := bufio.NewScanner(file)
	return scanner, nil
}

func (cmd CMDManager) WriteJSON(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
