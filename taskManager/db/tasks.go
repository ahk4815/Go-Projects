package db

// Compare with this : https://github.com/gophercises/task/blob/p2_bolt/main.go

import (
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = "taskBucket"
var db *bolt.DB

func init() {
	var err error
	db, err = bolt.Open("tasks.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(taskBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

func Add(task string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(taskBucket))
		id, _ := b.NextSequence()
		return b.Put(itob(int(id)), []byte(task))
	})
}

func List() error {
	return db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(taskBucket))
		c := b.Cursor()
		// TODO : return a slice of tasks
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Println(btoi(k), string(v))
		}
		return nil
	})
}

func get(key int, tx *bolt.Tx) []byte {
	b := tx.Bucket([]byte(taskBucket))
	return b.Get(itob(key))
}

func Do(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(taskBucket))
		val := get(key, tx)
		if val == nil {
			log.Fatal("No task was found with the key")
		}

		err := b.Delete(itob(key))
		if err != nil {
			return err
		}

		return nil
	})
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
