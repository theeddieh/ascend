package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Command is a single database command
type Command []string

// Database is an in-memory KV store.
// It tracks the current state and the transaction history.
type Database struct {
	state   map[string]string
	history []Command
}

func main() {
	// this should accept a single argument on the command line
	// pointing to the path of the input file
	if len(os.Args) < 2 {
		fmt.Println("No filepath specified")
		return
	}
	infile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to open file: %s", os.Args[1])
		return
	}
	defer infile.Close()
	fileScanner := bufio.NewScanner(infile)

	d := &Database{
		state: make(map[string]string),
	}

	for fileScanner.Scan() {
		if err != nil {
			break
		} else {
			switch command := strings.Fields(fileScanner.Text()); command[0] {
			case "WRITE":
				d.Write(command[1], command[2])
			case "DELETE":
				d.Delete(command[1])
			case "PRINT":
				d.Print()
			case "ROLLBACK":
				d.Rollback()
			case "#":
			default:
				fmt.Println(fmt.Errorf("unknown instruction `%s` found", command[0]))
				return
			}
		}
	}
	return
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
