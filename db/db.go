package db

import (
	"fmt"
)

// Command is a single database command
type command []string

// Database is an in-memory KV store.
// It tracks the current state and the transaction history.
type Database struct {
	state   map[string]string
	history []command
}

// DatabaseError handles errors specific to db operations.
type DatabaseError string

func (e DatabaseError) Error() string {
	return string(e)
}

// ErrMissingKey is returned when
const ErrMissingKey = DatabaseError("key not present")

// New in-memory database.
func New() (d *Database) {
	return &Database{
		state:   make(map[string]string),
		history: make([]command, 0),
	}
}

// Write the value `v` to the database under key `k`.
func (d Database) Write(k, v string) {
	d.state[k] = v
}

// Read the value stored under key `k`. Returns an error message if the key is missing.
func (d Database) Read(k string) (v string, err error) {
	v, ok := d.state[k]
	if !ok {
		err = ErrMissingKey
	}
	return v, err
}

// Delete the key `k` from the database.
func (d Database) Delete(k string) {
	delete(d.state, k)
}

// Print the current state of the database.
func (d Database) Print() {
	for k, v := range d.state {
		fmt.Println(k, v)
	}
}

// Rollback the database back to its state prior to the most recent command.
func (d Database) Rollback() {
}
