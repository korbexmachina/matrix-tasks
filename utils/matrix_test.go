package utils

import (
	"os"
	"testing"
)

func TestCreateMatrix(t *testing.T) {
	err := CreateMatrix("test.db")
	if err != nil {
		t.Errorf("CreateMatrix failed: %s", err)
	}

	// Check if the database file exists
	if _, err := os.Stat("test.db"); os.IsNotExist(err) {
		t.Errorf("CreateMatrix failed: %s", err)
	}

	// Check if the buckets exist
	for _, bucket := range Buckets {
		if !BucketExists(bucket) {
			t.Errorf("CreateMatrix failed: bucket %s does not exist", bucket)
		}
	}
}

func TestInsertTask(t *testing.T) {
	task := Task{ID: 1, Name: "Test", Bucket: "UrgentImportant"}
	err := InsertTask(task)
	if err != nil {
		t.Errorf("InsertTask failed: %s", err)
	}

	// Check if the task exists
	_, err = GetTask("UrgentImportant", "1")
	if err != nil {
		t.Errorf("InsertTask failed: %s", err)
	}
}

func TestDeleteTask(t *testing.T) {
	task := Task{ID: 1, Name: "Test", Bucket: "UrgentImportant"}
	err := DeleteTask(task)
	if err != nil {
		t.Errorf("DeleteTask failed: %s", err)
	}

	// Check if the task exists
	_, err = GetTask("UrgentImportant", "1")
	if err == nil {
		t.Errorf("DeleteTask failed: task still exists")
	}
}

func TestGetTask(t *testing.T) {
	task := Task{ID: 1, Name: "Test", Bucket: "UrgentImportant"}
	err := InsertTask(task)
	if err != nil {
		t.Errorf("GetTask failed: %s", err)
	}

	// Check if the task exists
	_, err = GetTask("UrgentImportant", "1")
	if err != nil {
		t.Errorf("GetTask failed: %s", err)
	}
}

func TestGetBucketTasks(t *testing.T) {
	task := Task{ID: 1, Name: "Test", Bucket: "UrgentImportant"}
	err := InsertTask(task)
	if err != nil {
		t.Errorf("GetBucketTasks failed: %s", err)
	}

	// Check if the task exists
	_, err = GetBucketTasks("UrgentImportant")
	if err != nil {
		t.Errorf("GetBucketTasks failed: %s", err)
	}
}

func TestDeleteMatrix(t *testing.T) {
	err := DeleteMatrix("test.db")
	if err != nil {
		t.Errorf("DeleteMatrix failed: %s", err)
	}

	// Check if the database file exists
	if _, err := os.Stat("test.db"); !os.IsNotExist(err) {
		t.Errorf("DeleteMatrix failed: %s", err)
	}
}
