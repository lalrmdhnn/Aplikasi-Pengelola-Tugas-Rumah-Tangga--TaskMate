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

func addTask() {
	//good
	if taskCount >= maxTasks {
		fmt.Println("Task table is full")
		return
	}

	fmt.Print("Task name    : ")
	fmt.Scan(&tableTask[taskCount].name)
	fmt.Print("Room/category: ")
	fmt.Scan(&tableTask[taskCount].room)
	fmt.Print("Difficulty (1-5): ")
	fmt.Scan(&tableTask[taskCount].difficulty)
	fmt.Print("Duration (minutes): ")
	fmt.Scan(&tableTask[taskCount].duration)
	tableTask[taskCount].done = false
	taskCount++
	fmt.Println("Task added!")

}

func editTask() {
	//good
	var tasktarget string
	fmt.Print("Enter task name to edit: ")
	fmt.Scan(&tasktarget)

	index := -1
	i := 0
	for i < taskCount && index == -1 {
		if tableTask[i].name == tasktarget {
			index = i
		}
		i++
	}

	if index == -1 {
		fmt.Println("Task not found")
		return
	}

	fmt.Print("New name    : ")
	fmt.Scan(&tableTask[index].name)
	fmt.Print("New room    : ")
	fmt.Scan(&tableTask[index].room)
	fmt.Print("New difficulty (1-5): ")
	fmt.Scan(&tableTask[index].difficulty)
	fmt.Print("New duration (minutes): ")
	fmt.Scan(&tableTask[index].duration)
	fmt.Println("Task updated!")

}

func deleteTask() {
	//okay good
	var tasktarget string
	fmt.Print("Enter task name to delete: ")
	fmt.Scan(&tasktarget)

	index := -1
	i := 0
	for i < taskCount && index == -1 {
		if tableTask[i].name == tasktarget {
			index = i
		}
		i++
	}

	if index == -1 {
		fmt.Println("Task not found")
		return
	}

	i = index
	for i < taskCount-1 {
		tableTask[i] = tableTask[i+1]
		i++
	}
	taskCount--
	fmt.Println("Task deleted!")

}
