package read

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLine() string {
	var input string

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input = scanner.Text()
	}

	if scanner.Err() != nil {
		fmt.Println("Error:", scanner.Err())
		input = "nil"
	}

	return input
}
