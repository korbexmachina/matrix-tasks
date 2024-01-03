package utils

import (
	"testing"
)

// Test initializing a database connection
func TestInitDB(t *testing.T) {
	err := InitDB("test.db")
	if err != nil {
		t.Errorf("Error initializing database: %s", err)
	}
}

// Test creating a bucket
func TestCreateBucket(t *testing.T) {
	err := CreateBucket("test")
	if err != nil {
		t.Errorf("Error creating bucket: %s", err)
	}
}

// Test inserting a key/value pair
func TestInsert(t *testing.T) {
	err := Insert("test", "key", "value")
	if err != nil {
		t.Errorf("Error inserting key/value pair: %s", err)
	}
}

// Test getting a key/value pair
func TestGet(t *testing.T) {
	val, err := Get("test", "key")
	if err != nil {
		t.Errorf("Error getting value: %s", err)
	}

	if val != "value" {
		t.Errorf("Expected value to be 'value', got '%s'", val)
	}
}

// Test deleting a key/value pair
func TestDelete(t *testing.T) {
	err := Delete("test", "key")
	if err != nil {
		t.Errorf("Error deleting key/value pair: %s", err)
	}

	// Print the matrix
	PrintMatrix("test.db")

	val, err := Get("test", "key")
	if err != nil {
		t.Errorf("Error getting value: %s", err)
	}

	if val != "" {
		t.Errorf("Expected value to be '', got '%s'", val)
	}
}

// Test deleting a bucket
func TestDeleteBucket(t *testing.T) {
	err := DeleteBucket("test")
	if err != nil {
		t.Errorf("Error deleting bucket: %s", err)
	}
}

// Test closing a database connection
func TestCloseDB(t *testing.T) {
	err := CloseDB()
	if err != nil {
		t.Errorf("Error closing database: %s", err)
	}
}
