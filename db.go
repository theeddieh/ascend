package main

import "fmt"

// Command is a single database command
type Command []string

// Database is an in-memory KV store.
// It tracks the current state and the transaction history.
type Database struct {
	state   map[string]string
	history []Command
}

// Write the value `v` to the database under key `k`.
func (d Database) Write(k, v string) {
	d.state[k] = v
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
