package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetMultiLineInput(prompt string) string {
	fmt.Printf("%s (type EOF as the last line to finish the input)\n", prompt)
	reader := bufio.NewReader(os.Stdin)
	var input strings.Builder

	for {
		line, err := reader.ReadString('\n') // Read until newline
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		line = strings.TrimSuffix(line, "\n") // Remove trailing newline
		if line == "EOF" {                   // Check for termination condition
			break
		}
		input.WriteString(line + "\n") // Append line to the builder
	}
	
	return input.String()
}