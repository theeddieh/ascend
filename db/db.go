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
		history: make([]string, 1),
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

	v = vals[len(vals)-1]
	if v == "" {
		// key has been deleted
		err = ErrKeyDeleted
	}
	return v, err
}

// Delete the key `k` from the database.
func (d *Database) Delete(k string) {
	d.Write(k, "")
}

// Print the current state of the database.
func (d *Database) Print() {
	for k, v := range d.state {
		latestValue := v[len(v)-1]
		if latestValue != "" {
			fmt.Println(k, latestValue)
		}
	}
}

// Rollback the previous write/delete command.
func (d *Database) Rollback() {
	l := len(d.history) - 1
	lastCommand := d.history[l]
	d.history = d.history[:l]

	l = len(d.state[lastCommand]) - 1
	d.state[lastCommand] = d.state[lastCommand][:l]
}
