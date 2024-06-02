package iomanager

import (
	"bufio"
)

type IOManager interface {
	Readfile() (*bufio.Scanner, error)
	WriteJSON(data interface{}) error
}
