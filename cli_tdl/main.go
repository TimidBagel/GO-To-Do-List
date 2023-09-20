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

type User struct {
	Name  string `json:"name"`
	Tasks []Task `json:"tasks"`
}

func main() {
	loop := true
	var user User

	for loop {
		fmt.Print("Enter your username >> ")
		username := read.ReadLine()

		// look for username in users folder
		path := "users/" + username + ".json"
		_, err := os.Stat(path)

		if err == nil {
			fmt.Println("\nHello,", username)
			user, _ = Deserialize(path)
		} else if os.IsNotExist(err) {

			fmt.Printf("Path '%s' does not exist.\n", path)
			fmt.Printf("Would you like to create a new user? (y/n) >> ")
			response := strings.ToLower(read.ReadLine())
			if response == "y" {

			} else if response == "n" {

			} else {

			}
		} else {

			fmt.Printf("Error checking path '%s': %v\n", path, err)
		}
	}

	// Saves work to json file specific to username
	Serialize(user)
}

func Serialize(user User) {
	// converts user struct to JSON bytes
	jsonData, err := json.MarshalIndent(user, "", "   ")

	if err != nil {
		fmt.Println("Error serializing user:", err)
		return
	}

	// creates a file
	path := "users/" + user.Name + ".json"
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

func Deserialize(path string) (User, error) {
	var user User

	// reads file as bytes
	jsonData, err := os.ReadFile(path)

	// handles error
	if err != nil {
		fmt.Printf("Error reading file '%s': %v\n", path, err)
		return user, err
	}

	valid := json.Valid(jsonData)

	if !valid {
		fmt.Println("yo json aint valid cheif!!1!")
		return user, errors.New("invlaid JSON")
	}

	// deserializes bytes, stored in user
	err = json.Unmarshal(jsonData, &user)

	// handles error
	if err != nil {
		fmt.Println("Error deserializing file:", err)
		return user, err
	}

	// returns address of deserialized user
	return user, err
}
