package utils

import (
	"fmt"
	"os"
)

var (
	// Buckets is the list of buckets in the Eisenhower Matrix
	Buckets = []string{"UrgentImportant", "NotUrgentImportant", "UrgentNotImportant", "NotUrgentNotImportant"}
)

// Task represents a task in the matrix
type Task struct {
	ID int
	Name string
	Bucket string
}

// CreateMatrix initializes the Eisenhower Matrix and the associated database
func CreateMatrix (path string) error {
	// Create the database
	InitDB(path)

	// Create the buckets
	for _, bucket := range Buckets {
		err := CreateBucket(bucket)
		if err != nil {
			return err
		}
	}

	return nil
}

// Insert task inserts a task into the database
func InsertTask (task Task) error {
	return Insert(task.Bucket, fmt.Sprint(task.ID), task.Name)
}

// DeleteTask deletes a task from the database
func DeleteTask (task Task) error {
	return Delete(task.Bucket, fmt.Sprint(task.ID))
}

// GetTask gets a task from the database
func GetTask (bucket string, id string) (Task, error) {
	name, err := Get(bucket, id)
	if err != nil {
		return Task{}, err
	}

	return Task{ID: int(id[0]), Name: name, Bucket: bucket}, nil
}

// GetBucketTasks gets all tasks from a bucket
func GetBucketTasks (bucket string) ([]Task, error) {
	var tasks []Task
	var task Task
	var err error

	// Get all tasks from the bucket
	keys, err := GetBucketKeys(bucket)
	if err != nil {
		return nil, err
	}

	// Get the tasks
	for _, key := range keys {
		task, err = GetTask(bucket, key)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

// PrintMatrix prints the Eisenhower Matrix for debugging & testing purposes
// Prints errors to stdout
func PrintMatrix (path string) {
	// Create the database connection
	InitDB(path)

	// Print the matrix
	for _, bucket := range Buckets {
		fmt.Println(bucket)
		keys, _ := GetBucketKeys(bucket)
		for _, key := range keys {
			value, err := Get(bucket, key)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(key, value)
		}
	}
}

// DeleteMatrix deletes the Eisenhower Matrix and the associated database
func DeleteMatrix (path string) error {
	// Delete the database
	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}
