package main

import (
	"encoding/json"
	"fmt"
	"os"

	"iainmods.com/read" // local module
)

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type User struct {
	Name  string   `json:"name"`
	Tasks [10]Task `json:"tasks"`
}

func main() {
	fmt.Print("Enter your username >> ")
	username := read.ReadLine()
	fmt.Println("\nHello,", username)

	task := Task{"task", "desc", false}
	user := User{username, [10]Task{task}} // copies task 10 times, not intended

	// ERROR: task name and description not in json serialization

	Serialize(&user)
}

func Serialize(user *User) {
	jsonData, err := json.MarshalIndent(user, "", "   ")

	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	path := "users/" + user.Name + ".json"

	file, err := os.Create(path)

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	_, err = file.Write(jsonData)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	file.Close()
}
