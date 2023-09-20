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

func HelpHelp(){
  fmt.Println("- 'help task' - displays 'task' command specific information")
  fmt.Println("- 'help quit' - displays 'quit' command specific information")
}

func TaskHelp(args ...string) error {
	if len(args) != 2 {
		return errors.New("Unexpected number of arguments")
	}

	fmt.Println("- 'task new [name] [description] - initializes and saves a new task")
	fmt.Println("- 'task edit [name] [new name] [new description] - edits the name and description of a task")
	fmt.Println("- 'task del [name] - deletes given task")
	fmt.Println("- 'task done [name] - marks a given task as completed")

	return nil
}

func QuitHelp(args ...string) error {
	if len(args) != 2 {
		return errors.New("Unexpected number of arguments")
	}

	fmt.Println("- 'quit prog' - ends program with 0 exit code")
	return nil
}

func NewTask(args ...string) error {
	if len(args) != 4 {
		return errors.New("Unexpected number of arguments")
	}

	return nil
}

func EditTask(args ...string) error {
	if len(args) != 5 {
		return errors.New("Unexpected number of arguments")
	}

	return nil
}

func DeleteTask(args ...string) error {
	if len(args) != 3 {
		return errors.New()
	}

	return nil
}

func CompleteTask(args ...string) error {
	if len(args) != 3 {
		return errors.New()
	}

	return nil
}

func main() {
	commands := map[string]map[string]func(args ...string) error{
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
			"prog": func(args ...string) error {
				os.Exit(0)
				return nil
			},
		},
	}

	loop := true

	for loop {
		fmt.Print(">")
		input := strings.Split(strings.ToLower(read.ReadLine()), " ")

    if input[0] == "help"{
      HelpHelp()
      continue
    }
		if com, valid := commands[input[0]]; valid{
      if len(input) > 1 && com[input[1]] != true{
        fmt.Printf("- '%s' is not a valid command string.")
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
