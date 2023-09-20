package main

/*
	WORK ON ERROR HANDLING!!!!!!!
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"read" // local module
)

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var openCommands = []string{"help", "task", "quit"}
var helpCommands = []string{"task"}
var taskCommands = []string{"new", "edit", "del", "done"}

var commads = make(map[string]map[string]string)

func main() {
	//var tasks []Task

	loop := true
	for loop {
		fmt.Print("->")
		input := (strings.Split(strings.ToLower(read.ReadLine()), " "))

		switch input[0] {
		case "help":
			if len(input) > 1 {
				switch input[1] {
				case "task":
					for _, val := range taskCommands {
						fmt.Println(val, ": explanation")
					}

				case "quit":
					fmt.Println("Quit: exits the program.")
				default:
					fmt.Printf("'%s' is not a recognized command.\n", input[1])
				}
			}

		case "task":
			if len(input) > 1 {
				switch input[1] {
				case "new":
					// new
				case "edit":
					// edit
				case "del":
					// delete
				case "done":
					// done
				default:
					fmt.Printf("'%s' is not a recognized command.\n", input[1])
				}
			}

		case "quit":
			return
		default:
			fmt.Printf("'%s' is not a recognized command.\n", input[0])
		}
	}
}

func Serialize(tasks []Task) {
	// converts task list to JSON bytes
	jsonData, err := json.MarshalIndent(tasks, "", "   ")

	if err != nil {
		fmt.Println("Error serializing tasks:", err)
		return
	}

	// creates a file
	path := "tasks/tasks.json"
	file, err := os.Create(path)

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// writes converted JSON to file
	_, err = file.Write(jsonData)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// closes file
	file.Close()
}

func Deserialize(path string) ([]Task, error) {

	var tasks []Task

	// reads file as bytes
	jsonData, err := os.ReadFile(path)

	// handles error
	if err != nil {
		fmt.Printf("Error reading file '%s': %v\n", path, err)
		return tasks, err
	}

	valid := json.Valid(jsonData)

	if !valid {
		fmt.Println("yo json aint valid cheif!!1!")
		return tasks, errors.New("invlaid JSON")
	}

	// deserializes bytes, stored in task list
	err = json.Unmarshal(jsonData, &tasks)

	// handles error
	if err != nil {
		fmt.Println("Error deserializing file:", err)
		return tasks, err
	}

	// returns deserialized task list
	return tasks, nil
}
