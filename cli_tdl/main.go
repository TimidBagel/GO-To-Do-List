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
	// checks if correct number of arguments are passed
	if len(args) != 2 {
		return errors.New("unexpected number of arguments")
	}

	fmt.Printf("--> %50s - %-30s\n", "'task new [name] [description]'", "initializes and saves a new task")
	fmt.Printf("--> %50s - %-30s\n", "'task edit [name] [new name] [new description]'", "edits the name and description of a task")
	fmt.Printf("--> %50s - %-30s\n", "'task del [name]'", "deletes given task")
	fmt.Printf("--> %50s - %-30s\n", "'task done [name]'", "marks a given task as completed")

	// return nil error if success
	return nil
}

func QuitHelp(args ...string) error {
	// checks if correct number of arguments are passed
	if len(args) != 2 {
		return errors.New("unexpected number of arguments")
	}

	fmt.Println("- 'quit 0' - ends program with 0 exit code")

	// return nil error if success
	return nil
}

// Takes a list of strings as a parameter, will return an error if there are more or less than 4 strings in that list.
// Iniitializes a new task, creates a new task file, and writes the serialized task JSON to that file. Returns a nil error
// if successful.
func NewTask(args ...string) error {
	// checks if correct number of arguments are passed
	if len(args) != 4 {
		return errors.New("unexpected number of arguments")
	}

	// initializes name and description
	name := args[2]
	desc := args[3]

	// initialize path based on name
	path := "tasks/" + name + ".json"

	// check if path exists
	_, err := os.Stat(path)

	// if does exist, return error. if doesn't exist and error, return error.
	if os.IsExist(err) {
		return errors.New("task name already exists")
	} else if !os.IsNotExist(err) {
		return err
	}

	// create file at path
	file, err := os.Create(path)

	// return error if failed
	if err != nil {
		return err
	}

	// closes file after completion
	defer file.Close()

	// initialize new task
	task := Task{name, desc, false}

	// convert task to JSON, write JSON to task file
	err = Serialize(task, file)

	// return error if failed
	if err != nil {
		return err
	}

	// return nil error if success
	return nil
}

func EditTask(args ...string) error {
	// checks if correct number of arguments are passed
	if len(args) != 5 {
		return errors.New("unexpected number of arguments")
	}

	path := "tasks/" + args[2] + ".json"

	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return errors.New("task not found")
	} else if !os.IsExist(err) {
		return err
	}

	name := args[3]
	desc := args[4]

	task := Task{name, desc, false}

	fmt.Print(task)

	// return nil error if success
	return nil
}

func DeleteTask(args ...string) error {
	// checks if correct number of arguments are passed
	if len(args) != 3 {
		return errors.New("unexpected number of arguments")
	}

	// return nil error if success
	return nil
}

func CompleteTask(args ...string) error {
	// checks if correct number of arguments are passed
	if len(args) != 3 {
		return errors.New("unexpected number of arguments")
	}

	// return nil error if success
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
				}

			} else {
				fmt.Printf("- Error: %v\n- *see 'help'* -\n", errors.New("unrecognized command string"))
			}
		} else {
			fmt.Printf("- Error: %v\n- *see 'help'* -\n", errors.New("unrecognized command string"))
		}
	}
}

// Serializes a Task struct into JSON, then writes that JSON byte array to a file. Takes
// two parameters, the Task struct, and type *os.File.
// Returns an error message or nil.
func Serialize(task Task, file *os.File) error {
	// converts task to JSON bytes
	jsonData, err := json.MarshalIndent(task, "", "   ")

	if err != nil {
		return err
	}

	// closes file
	defer file.Close()

	// writes converted JSON to file
	_, err = file.Write(jsonData)

	if err != nil {
		return err
	}

	return nil
}

// Deserializes a given task file. Takes one parameter: the path of the file. Returns
// a Task struct and an error, possibly nil.
func Deserialize(path string) (Task, error) {

	var task Task

	// reads file as bytes
	jsonData, err := os.ReadFile(path)

	// handles error
	if err != nil {
		return task, err
	}

	// checks if JSON is valid
	valid := json.Valid(jsonData)

	if !valid {
		fmt.Println("yo json aint valid cheif!!1!")
		return task, errors.New("invlaid JSON")
	}

	// deserializes bytes, stored in task list
	err = json.Unmarshal(jsonData, &task)

	// handles error
	if err != nil {
		return task, err
	}

	// returns deserialized task list
	return task, nil
}
