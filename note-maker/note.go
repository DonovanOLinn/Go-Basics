package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// An interface is not about defining a method, but stating that a method exists, the name, the return value, parameters, etc. 
type saver interface { 
	// naming convention is what it is doing plus 'er' at the end. 
	Save() error 
} 

// type displayer interface {
// 	Display()
// }

// If you have more than one method, you want to describe the key characteristics of the type you want to output

type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }


func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}
	
	err = outputData(userNote)

}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content :=getUserInput("Note Content:")

	return title, content
}

func getUserInput(prompt string) (string) {
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



func printSomething(value any) {
	intVal, ok := value.(int)

	if !ok {
		fmt.Println("Integer:", intVal)
		return 
	}

	floatVal, ok := value.(float64)

	if !ok {
		fmt.Println("Integer:", floatVal)
		return 
	}

	floatVal, ok := value.(float64)

	if !ok {
		fmt.Println("Integer:", floatVal)
		return 
	}
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64: 
	// 	fmt.Println("Float:", value)
	// case string: 
	// fmt.Println(value)
	// default: 
	// }
}