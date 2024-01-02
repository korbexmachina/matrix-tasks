package utils

import (
	bolt "go.etcd.io/bbolt"
)

var (
	// DB is the global database connection
	DB *bolt.DB
)

// InitDB initializes the database
func InitDB(dbPath string) error {
	var err error
	DB, err = bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return err
	}

	return nil
}

// CloseDB closes the database
func CloseDB() error {
	return DB.Close()
}

// CreateBucket creates a bucket
func CreateBucket(bucketName string) error {
	return DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})
}

// DeleteBucket deletes a bucket
func DeleteBucket(bucketName string) error {
	return DB.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte(bucketName))
	})
}

// Insert inserts a key/value pair into a bucket
func Insert(bucketName string, key string, value string) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		return b.Put([]byte(key), []byte(value))
	})
}

// Delete deletes a key/value pair from a bucket
func Delete(bucketName string, key string) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		return b.Delete([]byte(key))
	})
}

// Get gets a value from a bucket
func Get(bucketName string, key string) (string, error) {
	var value string
	err := DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		value = string(b.Get([]byte(key)))
		return nil
	})
	return value, err
}
