package db

import (
	"fmt"
)

// Database is an in-memory KV store with rollback support.
type Database struct {
	state   map[string][]string
	history []string
}

// DatabaseError handles errors specific to db operations.
type DatabaseError string

func (e DatabaseError) Error() string {
	return string(e)
}

// ErrKeyNonexistant is returned when trying to read a key that never existed
const ErrKeyNonexistant = DatabaseError("key does not exist")

// ErrKeyMissing is returned when a
const ErrKeyMissing = DatabaseError("key not present")

// ErrKeyDeleted is returned when trying to read a deleted key
const ErrKeyDeleted = DatabaseError("key deleted")

// New in-memory database.
func New() (d *Database) {
	return &Database{
		state:   make(map[string][]string),
		history: make([]string, 0),
	}
}

// Write the value `v` to the database under key `k`.
func (d *Database) Write(k, v string) {
	d.state[k] = append(d.state[k], v)
	d.history = append(d.history, k)
}

// Read the value stored under key `k`. Returns `ErrMissingKey` if empty.
func (d *Database) Read(k string) (v string, err error) {

	vals, ok := d.state[k]
	if !ok {
		// key never written
		return "", ErrKeyNonexistant
	}

	l := len(vals)
	if l == 0 {
		// key has been rolled back to before it was first written
		return "", ErrKeyMissing
	}

	v = vals[l-1]
	if v == "" {
		// key has been deleted
		err = ErrKeyDeleted
	}
	return v, err
}

// Delete the key `k` from the database, if present.
func (d *Database) Delete(k string) {
	// slightly hacky - only records a delete if the key already exists
	_, ok := d.state[k]
	if ok {
		d.Write(k, "")
	}
}

// Print the current state of the database.
func (d *Database) Print() {
	for k, v := range d.state {
		l := len(v)
		if l == 0 {
			continue
		}

		latestValue := v[l-1]
		if latestValue != "" {
			fmt.Println(k, latestValue)
		}
	}
}

// Rollback the previous write/delete command.
func (d *Database) Rollback() {
	lenHistory := len(d.history)
	if lenHistory == 0 {
		// nothing to rollback
		return
	}

	// check most recent key change
	latestKey := d.history[lenHistory-1]
	lenValues := len(d.state[latestKey])
	if lenValues == 0 {
		// nothing to rollback
		return
	}

	// roll back key's value to previous
	d.state[latestKey] = d.state[latestKey][:lenValues-1]

	// pop most recent key
	d.history = d.history[:lenHistory-1]
}
