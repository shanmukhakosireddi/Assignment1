package main

import (
	"bufio"
	"fmt"
	"os"

	"strings"
)


func main() {
todo()
}

func todo() {
	

	var todolist []string
	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nSlect an option:")
		fmt.Println("\n1. view task")
		fmt.Println("2. add tasks")
		fmt.Println("3. Exit")
		options, _ := input.ReadString('\n')

		option := strings.TrimSpace(options)
		if option == "1" {
			if len(todolist) == 0 {
				fmt.Println("\nNo tasks available. Please add a task first.")
			}
			for index, todolist := range todolist {
				fmt.Printf("your %v task is %v \n", index+1, todolist)
			}
		} else if option == "2" {
			fmt.Println("\nenter your task:")

			todoReder,_ := input.ReadString('\n')
			task := strings.TrimSpace(todoReder)
			
			todolist = append(todolist, task)
			
			if len(todolist) != 0 {
				fmt.Println("\nTask added successfully!")
			} else {
				fmt.Println("\nNo task added")
			}

		} else if option == "3" {
			break
		} else {
			fmt.Println("\nInvalid option. Please try again.")
		}

	}
}
