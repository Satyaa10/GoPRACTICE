package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var todos []string
var reader = bufio.NewReader(os.Stdin)

func input(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func show() {
	if len(todos) == 0 {
		fmt.Println("No tasks yet.")
		return
	}
	for i, t := range todos {
		fmt.Printf("%d. %s\n", i+1, t)
	}
}

func main() {
	for {
		fmt.Println("\n1. Add  2. View  3. Delete  4. Exit")
		choice := input("Choose: ")

		switch choice {
		case "1":
			task := input("Enter task: ")
			todos = append(todos, task)
			fmt.Println("Added.")
		case "2":
			show()
		case "3":
			show()
			numStr := input("Delete task #: ")
			var n int
			fmt.Sscan(numStr, &n)
			if n < 1 || n > len(todos) {
				fmt.Println("Invalid number.")
				continue
			}
			todos = append(todos[:n-1], todos[n:]...)
			fmt.Println("Deleted.")
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}
