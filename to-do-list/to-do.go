package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args

	if len(args) < 2 {

		fmt.Println("Error encountered. Accepted values are add, list, complete")
		return

	}

	switch os.Args[1] {
	case "add":
		addTask(os.Args[2:])

	case "list":
		listTasks()

	case "complete":
		completeTask(os.Args[2])

	}
}

//function to handle addTask

func addTask(task []string) {
	fmt.Printf("We are trying to add: %v \n", task)

	taskString := strings.Join(task, " ") // Join task strings into one

	// Try to open the file to read existing tasks
	file, err := os.Open("tasks.csv")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("tasks.csv does not exist. Creating a new file.")
			// Create the file if it doesn't exist
			file, err = os.Create("tasks.csv")
			if err != nil {
				fmt.Println("Error creating the file:", err)
				return
			}
			defer file.Close() // Close after creating
		}
	}

	// Read the existing tasks and close the file afterward
	reader := csv.NewReader(file)
	tasks, err := reader.ReadAll() // Read all tasks
	if err != nil {
		fmt.Println("Not able to read file:", err)
		file.Close() // Close the file before returning
		return
	}
	file.Close() // Close the file after reading

	nextIndex := len(tasks) + 1 // Determine next index

	// Now open the file for writing (in append mode)
	file, err = os.OpenFile("tasks.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening the file for writing:", err)
		return
	}
	defer file.Close() // Close the connection after writing

	writer := csv.NewWriter(file)

	// Write the new task to the CSV file
	if err := writer.Write([]string{fmt.Sprintf("%d", nextIndex), taskString, "false"}); err != nil {
		fmt.Println("Error writing to the CSV file:", err)
		return
	}

	// Ensure all buffered data is written to the file
	writer.Flush()

	// Check for any errors during flush
	if err := writer.Error(); err != nil {
		fmt.Println("Error flushing the writer:", err)
		return
	}

	fmt.Println("Task Added:", taskString)

}

func listTasks() {
	fmt.Println("Operation requested:", os.Args[1])
	file, err := os.Open("tasks.csv") //try opening the file first and capture any error

	if err != nil {

		fmt.Println("unable to open file", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file) // create a new reader and read the contents of
	tasks, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error encountered, unable to read data", err)
		return
	}

	if len(tasks) == 0 {

		fmt.Printf("no tasks found in the file")
		return
	}

	for _, task := range tasks {

		fmt.Printf("%s", task)
	}
}

//function to mark a task complete

func completeTask(taskID string) {

	file, err := os.OpenFile("tasks.csv", os.O_RDWR, 0644)

	if err != nil {

		fmt.Println("Unable to open the file", err)
		return

	}

	defer file.Close()

	reader := csv.NewReader(file)

	tasks, err := reader.ReadAll()

	if err != nil {
		fmt.Println("unable to read data", err)
		return
	}

	taskIndex := os.Args[2] // Assuming the task ID is passed as the second argument

	found := false

	for i, task := range tasks {

		if task[0] == taskIndex {

			tasks[i][2] = "true"
			found = true
			break
		}
	}

	if !found {

		fmt.Println("task not found")
		return
	}

	//rewind the file and clear its content to avoid duplicate or corrupt file
	file.Seek(0, 0)
	file.Truncate(0)

	//create a new writer

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, task := range tasks {

		if err := writer.Write(task); err != nil {

			fmt.Println("errr, could not write to the file")
			return
		}
	}
	fmt.Println("task has been marked as completed successfully")

}
