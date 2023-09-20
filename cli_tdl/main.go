package main

/*
	WORK ON ERROR HANDLING!!!!!!!
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"read"
	"strings"
)

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func TaskHelp(args ...string) bool {
	if len(args) != 2 {
		return false
	}

	fmt.Println("- 'task new [name] [description] - initializes and saves a new task")
	fmt.Println("- 'task edit [name] [new name] [new description] - edits the name and description of a task")
	fmt.Println("- 'task del [name] - deletes given task")
	fmt.Println("- 'task done [name] - marks a given task as completed")

	return true
}

func QuitHelp(args ...string) bool {
	fmt.Println("- 'quit prog' - ends program with 0 exit code")
	return true
}

func NewTask(args ...string) bool {

	return true
}

func EditTask(args ...string) bool {
	return true
}

func DeleteTask(args ...string) bool {
	return true
}

func CompleteTask(args ...string) bool {
	return true
}

func main() {
	commands := map[string]map[string]func(args ...string) bool{
		"help": {
			"task": TaskHelp,
			"quit": QuitHelp,
		},
		"task": {
			"new":  NewTask,
			"edit": EditTask,
			"del":  DeleteTask,
			"done": CompleteTask,
		},
		"quit": {
			"prog": func(args ...string) bool {
				os.Exit(0)
				return true
			},
		},
	}

	loop := true

	for loop {
		fmt.Print(">")
		input := strings.Split(strings.ToLower(read.ReadLine()), " ")

		if com, valid := commands[input[0]]; valid {
			if input[0] == "help" && len(input) < 2 {
				fmt.Println("help task")
				fmt.Println("help quit")
			} else if len(input) > 1 {
				if com2, valid := com[input[1]]; valid {
					if input[0] == "help" || input[0] == "quit" {
						com2()
					} else if len(input) == 4 {
						com2()
					} else {
						fmt.Printf("'%s' is invalid. See 'help task'.\n", strings.Join(input, " "))
					}
				} else {
					fmt.Printf("'%s' is not a recognized command.\n", input[1])
					continue
				}
			}
		} else {
			fmt.Printf("'%s' is not a recognized command.\n", input[0])
			continue
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
