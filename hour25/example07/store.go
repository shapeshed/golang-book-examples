package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"github.com/boltdb/bolt"
)

var (
	// errKeyNotFound is returned when a key is not found
	errKeyNotFound = errors.New("Key Not Found")
)

type Store struct {
	// db is the underlying handle to the db.
	db *bolt.DB
}

// initialize is used to set up all of the buckets.
func NewStore() (*Store, error) {
	handle, err := bolt.Open("tasks.db", 0600, nil)
	store := &Store{
		db: handle,
	}

	if err != nil {
		return nil, err
	}

	return store, nil
}

// Initialize sets up BoltDB
func (s *Store) Initialize() error {
	return s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil {
			return err
		}
		return nil
	})
}

// CreateTask persists a task
func (s *Store) CreateTask(t *Task) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))

		id, _ := b.NextSequence()
		t.Id = int(id)

		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		return b.Put(itob(t.Id), buf)
	})
}

// UpdateTask updates a given task
func (s *Store) UpdateTask(t *Task) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))

		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		return b.Put(itob(t.Id), buf)
	})
}

// GetTasks fetches all tasks from the store
func (s *Store) GetTasks() ([]*Task, error) {
	tasks := []*Task{}
	err := store.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))

		b.ForEach(func(k, v []byte) error {
			var t Task
			err := json.Unmarshal(v, &t)
			if err != nil {
				return err
			}
			tasks = append(tasks, &t)
			return nil
		})

		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetTask fetches a single task from the store
func (s *Store) GetTask(id int) (*Task, error) {
	var t Task
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		v := b.Get(itob(id))
		if v == nil {
			return errKeyNotFound
		} else {
			err := json.Unmarshal(v, &t)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// DeleteTask removes a task from the store
func (s *Store) DeleteTask(id int) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		err := b.Delete(itob(id))
		if err != nil {
			return err
		}
		return nil
	})
}

// Close is used to gracefully close the DB connection.
func (b *Store) Close() error {
	return b.db.Close()
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
