package db

import (
	"errors"
	"log"
	"path/filepath"
	"time"

	"github.com/boltdb/bolt"
)

var (
	db      *bolt.DB
	buckets = []string{
		"abilities",
		"heroes",
		"items",
		"units",
		"general",
	}
)

// Close closes the bolt db.
// Should be called with defer after init in the main() func.
func Close() {
	err := db.Close()
	if err != nil {
		log.Println(err)
	}
}

// Setup opens a new bolt db and ensures default buckets exist.
func Setup(dir string) {
	path := filepath.Join(dir, "database.db")
	options := &bolt.Options{
		Timeout: 1 * time.Second,
	}

	var err error
	db, err = bolt.Open(path, 0600, options)
	if err != nil {
		log.Fatal(err)
	}

	// Ensure main Buckets exist
	err = db.Update(func(tx *bolt.Tx) error {
		for _, name := range buckets {
			_, err := tx.CreateBucketIfNotExists([]byte(name))
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalln("Cannot create default buckets: ", err)
	}
}

// PutToBucket adds an item to bucket
func PutToBucket(key string, content []byte, bucket string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put([]byte(key), content)
	})
	return err
}

// GetFromBucket returns an item from a bucket
func GetFromBucket(key string, bucket string) ([]byte, error) {
	var data []byte
	var err error
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		data = b.Get([]byte(key))
		return nil
	})
	if len(data) == 0 {
		err = errors.New("No data found")
	}
	return data, err
}

// DeleteFromBucket deletes an item from a bucket
func DeleteFromBucket(key string, bucket string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Delete([]byte(key))
	})
	return err
}

// GetAllFromBucket returns all items from bucket
func GetAllFromBucket(bucket string) (map[string][]byte, error) {
	items := make(map[string][]byte)
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		b.ForEach(func(k, v []byte) error {
			items[string(k)] = v
			return nil
		})
		return nil
	})
	return items, err
}
