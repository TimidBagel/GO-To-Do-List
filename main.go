/*
File: main.go
Author: Iain Broomell
A console application, functions like a CLI.
Start by entering 'help' to see available commands.
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"read"
	"strconv"
	"strings"
)

// Task struct, takes a name, description, and a true/false for whether its been completed.
type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// Displays available base commands
func HelpHelp() {
	fmt.Println("- 'help-task' - displays 'task' command specific information")
	fmt.Println("- 'help-quit' - displays 'quit' command specific information")
}

// Displays task specific information
func TaskHelp(args ...string) error {
	// checks if correct number of arguments are passed
	if len(args) != 2 {
		return errors.New("unexpected number of arguments")
	}

	fmt.Printf("--> %-48s - %-30s\n", "'task-new-[name]-[description]'", "initializes and saves a new task")
	fmt.Printf("--> %-48s - %-30s\n", "'task-edit-[name]-[new name]-[new description]'", "edits the name and description of a task")
	fmt.Printf("--> %-48s - %-30s\n", "'task-view'", "prints all tasks in task directory")
	fmt.Printf("--> %-48s - %-30s\n", "'task-del-[name]'", "deletes given task")
	fmt.Printf("--> %-48s - %-30s\n", "'task-done-[name]'", "marks a given task as completed")

	// return nil error if success
	return nil
}

// displays quit specific information
func QuitHelp(args ...string) error {
	// checks if correct number of arguments are passed
	if len(args) != 2 {
		return errors.New("unexpected number of arguments")
	}

	fmt.Println("- 'quit-0' - ends program with 0 exit code")

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

	// check if path exists, captures possible error
	_, err := os.Stat(path)

	// if does exist, return error. if doesn't exist and error, return error
	if err == nil {
		return errors.New("task already exists")
	}

	// initialize new task
	task := Task{name, desc, false}

	// convert task to JSON, write JSON to task file, captures possible error
	err = Serialize(task, path)

	// if error, return error
	if err != nil {
		return err
	}

	// return nil error if success
	return nil
}

// Takes three arguments: the file name, the new file name, and the new description. Reads old file, deletes old file,
// then creates a new file and writes the new task object to that file
func EditTask(args ...string) error {
	// checks if correct number of arguments are passed
	if len(args) != 5 {
		return errors.New("unexpected number of arguments")
	}

	// specifies local task location
	path := "tasks/" + args[2] + ".json"

	// attempts to find task file, captures possible error
	_, err := os.Stat(path)

	// checks if task exists, if not or error returned, return error
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("task not found")
		} else {
			return err
		}
	}

	// initializes new name and new description strings
	name := args[3]
	desc := args[4]

	// initializes new task with new information
	task := Task{name, desc, false}

	// removes task file from directory, captures possible error
	err = os.Remove(path)

	// if error, return error
	if err != nil {
		return err
	}

	// specifies local task location
	path = "tasks/" + name + ".json"

	// serializes and writes task to file, captures possible error
	err = Serialize(task, path)

	// if error, return error
	if err != nil {
		return err
	}

	// return nil error if success
	return nil
}

// Reads all task files in 'task/' directory, then prints them to the console
func ViewTasks(args ...string) error {
	// checks if correct number of arguments passed
	if len(args) != 2 {
		return errors.New("unexpected number of arguments")
	}

	// specifies local task directory
	path := "tasks/"

	// finds all files in task directory, captures possible error
	files, err := os.ReadDir(path)

	// if error, return error
	if err != nil {
		return err
	}

	// initialize output string
	var output string

	// loop through files in task directory
	for _, file := range files {
		// deserialize task object and capture possible error
		task, err := Deserialize(path + file.Name())

		// if error, return error
		if err != nil {
			return err
		}

		// append task information to output string
		output += "\n" + task.Name + " - Done: " + strconv.FormatBool(task.Done) + "\n\t" + task.Description
	}

	// print entire output string
	fmt.Println(output)

	// return nil error if success
	return nil
}

// Takes one argument: the name of the task. Finds the task, then removes it.
func DeleteTask(args ...string) error {
	// checks if correct number of arguments are passed
	if len(args) != 3 {
		return errors.New("unexpected number of arguments")
	}

	// specifies local task location
	path := "tasks/" + args[2] + ".json"

	// if not found, return 'task not found', if error, return error
	_, err := os.Stat(path)

	// if error, return error
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("task not found")
		}
		return err
	}

	// removes task file from directory, captures possible error
	err = os.Remove(path)

	// if error, return error
	if err != nil {
		return err
	}

	// return nil error if success
	return nil
}

// Takes one argument: the name of the task. Finds the task, deserializes it, marks it as done,
// then serializes it and writes the modified task to the same file
func CompleteTask(args ...string) error {
	// checks if correct number of arguments are passed
	if len(args) != 3 {
		return errors.New("unexpected number of arguments")
	}

	// specifies local task location
	path := "tasks/" + args[2] + ".json"

	// deserializes task, captures returned task and possible error
	task, err := Deserialize(path)

	// if error, return error
	if err != nil {
		return err
	}

	// set task as complete
	task.Done = true

	// serialize task and write to file, captures possible error
	err = Serialize(task, path)

	// if error, return error
	if err != nil {
		return err
	}

	// return nil error if success
	return nil
}

func main() {
	// nested dictionary of functions
	commands := map[string]map[string]func(args ...string) error{
		"help": {
			"task": TaskHelp,
			"quit": QuitHelp,
		},
		"task": {
			"new":  NewTask,
			"edit": EditTask,
			"view": ViewTasks,
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

	// runtime loop
	for loop {
		fmt.Print("\n>")

		// capture user input, set it to lowercase, and split it by the '-' character
		input := strings.Split(strings.ToLower(read.ReadLine()), "-")

		// if only input is 'help', display help information
		if len(input) == 1 && input[0] == "help" {
			HelpHelp()
			continue
		}

		// chekcs if first argument is in first level of commands map
		if com, valid := commands[input[0]]; valid {
			// checks if there is more than one argument
			if len(input) > 1 {
				// checks if second argument is in second level of commands map
				if com2, present := com[input[1]]; present {
					// calls function on third level of commands map, captures possible error
					result := com2(input...)
					// if no error, do nothing. if error, print error, suggest help
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

// Serializes a Task struct into JSON, then writes that JSON byte array to a file. Takes
// two parameters, the Task struct, and type *os.File.
// Returns an error message or nil.
func Serialize(task Task, path string) error {
	// converts task to JSON bytes, captures byte array and possible error
	jsonData, err := json.MarshalIndent(task, "", "   ")

	// if error, return error
	if err != nil {
		return err
	}

	// creates new task file, captures file and possible error
	file, err := os.Create(path)

	// if error, return error
	if err != nil {
		return err
	}

	// closes file
	defer file.Close()

	// writes converted JSON to file, captures possible error
	_, err = file.Write(jsonData)

	// if error, return error
	if err != nil {
		return err
	}

	// returns nil error if success
	return nil
}

// Deserializes a given task file. Takes one parameter: the path of the file. Returns
// a Task struct and an error, possibly nil.
func Deserialize(path string) (Task, error) {
	// initializes new empty task struct
	var task Task

	// attempts to find task file, captures possible error
	_, err := os.Stat(path)

	// if not found, return empty task and 'task not found', if error, return error
	if err != nil {
		if os.IsNotExist(err) {
			return task, errors.New("task not found")
		}
		return task, err
	}

	// reads file as bytes
	jsonData, err := os.ReadFile(path)

	// if error, return empty task and error
	if err != nil {
		return task, err
	}

	// checks if JSON is valid
	valid := json.Valid(jsonData)

	// if invalid, return empty task and error
	if !valid {
		fmt.Println("yo json aint valid cheif!!1!")
		return task, errors.New("invlaid JSON")
	}

	// deserializes bytes, stored in task list
	err = json.Unmarshal(jsonData, &task)

	// if error, return error
	if err != nil {
		return task, err
	}

	// returns deserialized task list
	return task, nil
}
