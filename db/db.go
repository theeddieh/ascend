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

// ErrMissingKey is returned when a db read fails
const ErrMissingKey = DatabaseError("key not present")

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
		// key never existed
		return "", ErrMissingKey
	}

	l := len(vals)
	if l == 0 {

		return "", ErrMissingKey
	}

	v = vals[len(vals)-1]
	if v == "" {
		// key has been deleted
		err = ErrMissingKey
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
