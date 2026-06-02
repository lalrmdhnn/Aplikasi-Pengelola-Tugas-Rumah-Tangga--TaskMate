package main

import "fmt"

const maxTasks = 100

// list elements needed for task management (the data going to be stored)
type Task struct {
	name       string
	room       string
	difficulty int
	duration   int
	done       bool
}

var tableTask [maxTasks]Task
var taskCount int

func main() {
	choice := -1

	//listing the required sub-programs for the task management system
	for choice != 0 {
		fmt.Println("\n====== TaskMate ======")
		fmt.Println("1. Add Task")
		fmt.Println("2. Edit Task")
		fmt.Println("3. Delete Task")
		fmt.Println("4. Mark Task as Done")
		fmt.Println("5. Display All Tasks")
		fmt.Println("6. Search Task")
		fmt.Println("7. Sort Tasks")
		fmt.Println("8. Statistics")
		fmt.Println("0. Exit")
		fmt.Print("Choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			//addTask()
		} else if choice == 2 {
			//editTask()
		} else if choice == 3 {
			//deleteTask()
		} else if choice == 4 {
			//markTaskDone()
		} else if choice == 5 {
			//displayAll()
		} else if choice == 6 {
			//searchMenu()
		} else if choice == 7 {
			// sortMenu()
		} else if choice == 8 {
			// showStatistics()
		} else if choice == 0 {
			fmt.Println("Goodbye!")
		} else {
			fmt.Println("Invalid choice.")
		}
	}
}
