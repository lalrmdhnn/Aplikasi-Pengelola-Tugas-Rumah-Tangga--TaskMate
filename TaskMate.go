package main

import "fmt"

// Biodata (group member) :

// Hilal Ramadhan 10301254004
// project description, func main()
// searchMenu(), searchByName(), searchByRoom(), checkIfSorted()
// selectionSortByDifficulty(), insertionSortByDuration()
// showStatistics()

// Bryan Junata Teezar Prasetyo 103012540015
// addTask(), editTask(), deleteTask(), markTaskDone(), displayAll()
// sortByRoom(),isAlreadySortedByDifficulty(), isAlreadySortedByDuration()
// sortMenu()

// Project Description : TaskMate is an application for managing and scheduling various household tasks on a daily basis. The main data used consists of task type, difficulty level, and estimated working time. The users of the application are family members or household residents. Specifications:

// a. Users can add, modify, and delete household task records.
// b. The system can record job descriptions, difficulty levels, and task durations in minutes.
// c. Users can search for task data based on the task name or room category using Sequential Search and Binary Search.
// d. Users can sort task data based on difficulty level or estimated completion time using Selection Sort and Insertion Sort.
// e. The system can display statistics on the number of completed tasks and the average amount of time spent working. Display the results in an informative format.

// listing the sub-program needed to be made
// addTask()
// editTask()
// deleteTask()
// markTaskDone()
// displayAll()
// searchMenu()
// sortMenu()
// showStatistics()
// creating the program ( each member has their part )

const maxTasks = 100

// list elements needed for task management (the data going to be stored)
type Task struct {
	name       string
	room       string
	difficulty int
	duration   int
	timetaken  int
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
			addTask()
		} else if choice == 2 {
			editTask()
		} else if choice == 3 {
			deleteTask()
		} else if choice == 4 {
			markTaskDone()
		} else if choice == 5 {
			displayAll()
		} else if choice == 6 {
			searchMenu()
		} else if choice == 7 {
			sortMenu()
		} else if choice == 8 {
			showStatistics()
		} else if choice == 0 {
			fmt.Println("Goodbye!")
		} else {
			fmt.Println("Invalid choice.")
		}
	}
}

// creating the sub-programs to add data
func addTask() {

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

// creating the sub-programs to edit data
func editTask() {

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

// creating the sub-programs to delete data
func deleteTask() {

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
	for i < taskCount-1 { // shift the tasks after the deleted task to the left
		tableTask[i] = tableTask[i+1]
		i++
	}
	taskCount-- //decrement the task count after deleting a task
	fmt.Println("Task deleted!")

}

// creating the sub-programs to mark task as done
// when marking a task as done, the user will be asked to input how long it took to complete the task,
// and the data will be stored in the timetaken variable, which can be used later for statistics
func markTaskDone() {
	var target string
	fmt.Print("Enter task name to mark as done: ")
	fmt.Scan(&target)

	i := 0
	for i < taskCount {
		if tableTask[i].name == target {
			if tableTask[i].done {
				fmt.Println("Task is already marked as done.")
				return
			}
			tableTask[i].done = true
			fmt.Print("How long did it take? (minutes): ")
			fmt.Scan(&tableTask[i].timetaken)
			fmt.Println("Task marked as done!")
			return
		}
		i++
	}
	fmt.Println("Task not found.")
}

// creating the sub-programs to display all data
// simple loop to display all the tasks in a nice format, showing the name, room, difficulty, duration, and whether the task is done or not
func displayAll() {

	if taskCount == 0 {
		fmt.Println("No tasks recorded.")
		return
	}
	fmt.Println("=======================================================")
	fmt.Printf("%-20s %-12s %-10s %-8s %-5s\n", "Name", "Room", "Difficulty", "Duration", "Done")
	fmt.Println("=======================================================")

	i := 0
	for i < taskCount {
		status := "No"
		if tableTask[i].done {
			status = "Yes"
		}
		fmt.Printf("%-20s %-12s %-10d %-8d %-5s\n",
			tableTask[i].name, tableTask[i].room, tableTask[i].difficulty, tableTask[i].duration, status)
		i++
	}
	fmt.Println("=======================================================")
}

// creating the sub-programs to search data
func searchMenu() {
	var searchchoice int
	fmt.Println("\n-- Search Menu --")
	fmt.Println("1. Linear search by Name")
	fmt.Println("2. Binary search by Room")
	fmt.Println("0. Back")
	fmt.Print("Choice: ")
	fmt.Scan(&searchchoice)

	if searchchoice == 1 { //using linear search
		searchByName()

	} else if searchchoice == 2 { //using binary search
		if checkIfSorted() { // check if the tasks are sorted by room before performing binary search
			searchByRoom()
		} else {
			sortMenu() // if not sorted, sort the tasks by room first
			searchByRoom()
		}
	} else {
		return
	}

}

// seperate search functions for duration and difficulty
func searchByName() { //linear search
	var target string
	fmt.Print("Enter task name to search (back to return): ")
	fmt.Scan(&target)

	if target == "back" {
		return
	}

	i := 0
	found := false

	for i < taskCount { //simple linear search by name, since the tasks are not sorted by name, we have to use linear search to find the target task
		if tableTask[i].name == target {
			fmt.Printf("Task found: \n %s in room %s, difficulty %d, duration %d minutes, done: %t\n",
				tableTask[i].name, tableTask[i].room, tableTask[i].difficulty, tableTask[i].duration, tableTask[i].done)
			found = true
		}
		i++
	}
	if !found {
		fmt.Println("Task not found.")
	}
}

// binary search by room, since the tasks are sorted by room, we can use binary search to find the target room faster than linear search
func searchByRoom() {
	var target string
	fmt.Print("Enter room number to search (back to return): ")
	fmt.Scan(&target)

	if target == "back" {
		return
	}

	left := 0
	right := taskCount - 1
	found := false
	// binary search by room, since the tasks are sorted by room, we can use binary search to find the target room faster than linear search
	for left <= right {
		mid := (left + right) / 2

		if tableTask[mid].room == target && !found {
			fmt.Printf("Task found: \n %s in room %s, difficulty %d, duration %d minutes, done: %t\n",
				tableTask[mid].name, tableTask[mid].room, tableTask[mid].difficulty, tableTask[mid].duration, tableTask[mid].done)
			found = true
			break
		} else if tableTask[mid].room < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}

	if !found {
		fmt.Println("No tasks found in that room.")
	}

}

func checkIfSorted() bool {
	// check if the tasks are sorted by room before performing binary search
	isSorted := true
	i := 1
	for i < taskCount {
		if tableTask[i].room < tableTask[i-1].room {
			isSorted = false
			break
		}
		i++
	}

	if !isSorted {
		fmt.Println("Tasks are not sorted by room. Please sort the tasks first.")
		return false
	}
	return true
}

// create sortMenu
func sortByRoom() {
	// insertion sort by room
	i := 1
	for i < taskCount {
		key := tableTask[i]
		j := i - 1
		for j >= 0 && tableTask[j].room > key.room {
			tableTask[j+1] = tableTask[j]
			j--
		}
		tableTask[j+1] = key
		i++
	}
	fmt.Println("Tasks sorted by room.")
}

// check if the tasks are already sorted by difficulty
// did alot of sub-programs to check if the data is already sorted before performing the sorting algorithm, to optimize the program and avoid unnecessary sorting
func isAlreadySortedByDifficulty() bool {
	i := 0
	for i < taskCount-1 {
		if tableTask[i].difficulty > tableTask[i+1].difficulty {
			return false
		}
		i++
	}
	return true
}

// another function to check if the tasks are already sorted by duration
// just a simple incremental check to see if the duration of the current task is greater than the next task, if it is, then the data is not sorted by duration
func isAlreadySortedByDuration() bool {
	i := 0
	for i < taskCount-1 {
		if tableTask[i].duration > tableTask[i+1].duration {
			return false
		}
		i++
	}
	return true
}

// creating the sub-programs to sort data
// sortbydifficulty, sortbyduration, sortbyroom (for binary search)

func sortMenu() {
	choice := -1
	for choice != 0 {
		fmt.Println("\n-- Sort Menu --")
		fmt.Println("1. Selection Sort by Difficulty")
		fmt.Println("2. Insertion Sort by Duration")
		fmt.Println("3. Sort by Room (for binary search)")
		fmt.Println("0. Back")
		fmt.Print("Choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			if isAlreadySortedByDifficulty() {
				fmt.Println("Data is already sorted by difficulty, no need to sort.")
			} else {
				selectionSortByDifficulty()
			}
		} else if choice == 2 {
			if isAlreadySortedByDuration() {
				fmt.Println("Data is already sorted by duration, no need to sort.")
			} else {
				insertionSortByDuration()
			}
		} else if choice == 3 {
			if checkIfSorted() {
				fmt.Println("Data is already sorted by room, no need to sort.")
			} else { //being called from the searchMenu, so we can directly call the sortByRoom function without having to go back to the sortMenu again
				sortByRoom()
			}
		} else if choice == 0 {
			// back
		} else {
			fmt.Println("Invalid choice.")
		}
	}
}

// selection sort by difficulty, since the tasks are not sorted by difficulty
// we have to use selection sort to sort the tasks by difficulty
// which is a simple sorting algorithm that works by repeatedly selecting the minimum element from the unsorted part and putting it at the beginning of the sorted part
func selectionSortByDifficulty() {
	i := 0
	for i < taskCount-1 {
		minIndex := i
		j := i + 1
		for j < taskCount {
			if tableTask[j].difficulty < tableTask[minIndex].difficulty {
				minIndex = j
			}
			j++
		}
		if minIndex != i {
			tableTask[i], tableTask[minIndex] = tableTask[minIndex], tableTask[i]
		}
		i++
	}
	fmt.Println("Tasks sorted by difficulty.")
}

// insertion sort by duration, since the tasks are not sorted by duration
// we have to use insertion sort to sort the tasks by duration
// which is a simple sorting algorithm that works by repeatedly inserting the current element into the sorted part of the array
func insertionSortByDuration() {
	i := 1
	for i < taskCount {
		key := tableTask[i]
		j := i - 1
		for j >= 0 && tableTask[j].duration > key.duration {
			tableTask[j+1] = tableTask[j]
			j--
		}
		tableTask[j+1] = key
		i++
	}
	fmt.Println("Tasks sorted by duration.")
}

// just like the displayAll function, we will display the statistics in a nice format
//
//	showing the total number of tasks, the number of completed tasks, and the average time spent on completed tasks
func showStatistics() {
	if taskCount == 0 {
		fmt.Println("No tasks recorded.")
		return
	}

	completedTasks := 0
	totalTime := 0

	i := 0
	// just a simple loop to count the number of completed tasks and the total time spent on completed tasks
	// we will use the timetaken variable to calculate the average time spent on completed tasks, which is the total time divided by the number of completed tasks
	for i < taskCount {
		if tableTask[i].done {
			completedTasks++
			totalTime += tableTask[i].timetaken
		}
		i++
	}

	averageTime := 0
	// since when adding a task we ask for estimated duration
	// we added a feature to ask for the actual time taken when marking a task as done, so we can use that data to calculate the average time spent on completed tasks
	if completedTasks > 0 {
		averageTime = totalTime / completedTasks
	}

	fmt.Printf("Total tasks: %d\n", taskCount)
	fmt.Printf("Completed tasks: %d\n", completedTasks)
	fmt.Printf("Average time spent on completed tasks: %d minutes\n", averageTime)
}
