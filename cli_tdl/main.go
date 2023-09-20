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

func HelpHelp() {
	fmt.Println("- 'help task' - displays 'task' command specific information")
	fmt.Println("- 'help quit' - displays 'quit' command specific information")
}

func TaskHelp(args ...string) error {
	if len(args) != 2 {
		return errors.New("unexpected number of arguments")
	}

	fmt.Printf("--> %50s - %-30s\n", "'task new [name] [description]'", "initializes and saves a new task")
	fmt.Printf("--> %50s - %-30s\n", "'task edit [name] [new name] [new description]'", "edits the name and description of a task")
	fmt.Printf("--> %50s - %-30s\n", "'task del [name]'", "deletes given task")
	fmt.Printf("--> %50s - %-30s\n", "'task done [name]'", "marks a given task as completed")

	return nil
}

func QuitHelp(args ...string) error {
	if len(args) != 2 {
		return errors.New("unexpected number of arguments")
	}

	fmt.Println("- 'quit 0' - ends program with 0 exit code")
	return nil
}

func NewTask(args ...string) error {
	if len(args) != 4 {
		return errors.New("unexpected number of arguments")
	}

	name := args[2]
	desc := args[3]

	path := "tasks/" + name + ".json"

	_, err := os.Stat(path)

	if os.IsExist(err) {
		return errors.New("task name already exists")
	} else if !os.IsNotExist(err) {
		return err
	}

	file, err := os.Create(path)

	if err != nil {
		return err
	}

	task := Task{name, desc, false}

	err = Serialize(task, file)

	if err != nil {
		return err
	}

	return nil
}

func EditTask(args ...string) error {
	if len(args) != 5 {
		return errors.New("unexpected number of arguments")
	}

	return nil
}

func DeleteTask(args ...string) error {
	if len(args) != 3 {
		return errors.New("unexpected number of arguments")
	}

	return nil
}

func CompleteTask(args ...string) error {
	if len(args) != 3 {
		return errors.New("unexpected number of arguments")
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
			"0": func(args ...string) error {
				os.Exit(0)
				return nil
			},
		},
	}

	loop := true

	for loop {
		fmt.Print(">")
		input := strings.Split(strings.ToLower(read.ReadLine()), " ")
		fmt.Printf("Arguments: %v - %v\n", len(input), input)

		if len(input) == 1 && input[0] == "help" {
			HelpHelp()
			continue
		}
		if com, valid := commands[input[0]]; valid {
			if len(input) > 1 {
				if com2, present := com[input[1]]; present {
					result := com2(input...)
					if result == nil {
						continue
					} else {
						fmt.Printf("- Error: %v\n- *see 'help %s'* -\n", result, input[0])
						continue
					}
				} else {
					fmt.Printf("- Error: %v\n- *see 'help'* -\n", errors.New("unrecognized command string"))
				}

			} else {
				fmt.Printf("- Error: %v\n- *see 'help'* -\n", errors.New("unrecognized command string"))
			}
		} else {
			fmt.Printf("- Error: %v\n- *see 'help'* -\n", errors.New("unrecognized command string"))
		}
	}
}

func Serialize(task Task, file *os.File) error {
	// converts task to JSON bytes
	jsonData, err := json.MarshalIndent(task, "", "   ")

	if err != nil {
		return err
	}

	// writes converted JSON to file
	_, err = file.Write(jsonData)

	if err != nil {
		return err
	}

	// closes file
	file.Close()

	return nil
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
