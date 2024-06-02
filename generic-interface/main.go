package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"temmix.com/generic-interface/todo"
)

type SaveAndDisplay interface {
	Display()
	Save() error
}

func main() {
	// Testing Genericx
	fmt.Println("************************ USING GENERICS ******************************")
	valInt := addNthing(1, 9)
	valFloat := addNthing(1.8, 2.9)
	valString := addNthing("Hello ", "World")
	fmt.Println(valInt, valFloat, valString)
	fmt.Println("************************ USING GENERICS END ******************************")
	fmt.Println()

	// Using Interface in func
	fmt.Println("************************ USING INTERFACES ******************************")
	printNThing(1)
	printNThing(1.0)
	printNThing("1")
	fmt.Println("************************ USING INTERFACES END ******************************")
	fmt.Println()

	// Using any in func
	fmt.Println("************************ USING ANY ******************************")
	CheckNThing(1)
	CheckNThing(1.0)
	CheckNThing("1")
	fmt.Println("************************ USING ANY END ******************************")
	fmt.Println()

	content := getTodoData()
	userTodo, err := todo.New(content)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("************************ USING INTERFACES FOR ANY STRUCT ******************************")
	// instead of call the fx on the struct, you can use interface with any struct with the contract
	// and you do not need to join the interface with struct whatsoever :)
	err = outputData(userTodo)

	// userTodo.Display()
	// err = userTodo.Save()

	if err != nil {
		fmt.Println("Saving the todo failed.")
	}

	fmt.Println("Saving the todo successed .")
}

func getTodoData() string {
	content := getUserInput("Content: ")
	return content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

// Using the interface type for the func
func outputData(data SaveAndDisplay) error {
	data.Display()
	return data.Save()
}

// Using Generic for something :)
func addNthing[T int | float64 | string](a, b T) T {
	return a + b
}

// Using interface
func printNThing(value interface{}) {
	// Not value.(type) its only for switch statement
	switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float64:", value)
	case string:
		fmt.Println("String:", value)
	}
}

func CheckNThing(value any) {
	intTypedVale, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intTypedVale)
	}

	floatTypedVale, ok := value.(float64)
	if ok {
		fmt.Println("Float 64:", floatTypedVale)
	}

	stringTypedVale, ok := value.(string)
	if ok {
		fmt.Println("String:", stringTypedVale)
	}
}
