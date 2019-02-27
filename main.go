/*

Exercise:

We'd like to see how you would finish this coding exercise.
The goal is to implement a basic in-memory nosql database.

Each line of the input file contains an instruction followed by a set of whitespace delimited arguments.
Your program should run the commands in the file line by line and print the final database state to standard out.

You can choose to finish the exercise in any language you prefer.
We have started a portion of the script in golang, as this is the primary language you can expect to work on here at Ascend.

We have already implemented WRITE, DELETE, and PRINT for you.
It is up to you to implement ROLLBACK.

Feel free to change any of the code we have provided.
You will be expected to review your coding exercise at your on-site interview, so be ready to talk through design choices and tradeoffs that you made in your program.

input-1.log:
WRITE key-0 val-1
WRITE key-1 val-3
WRITE key-2 val-4
DELETE key-1
DELETE key-0
ROLLBACK
ROLLBACK
WRITE key-2 val-8
DELETE key-0
PRINT

Output:
key-1 val-3
key-2 val-8

*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

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
	myDb := make(map[string]string)
	for fileScanner.Scan() {
		if err != nil {
			break
		} else {
			switch command := strings.Fields(fileScanner.Text()); command[0] {
			case "WRITE":
				myDb[command[1]] = command[2]
			case "DELETE":
				delete(myDb, command[1])
			case "PRINT":
				for key, value := range myDb {
					fmt.Println(key, value)
				}
			default:
				fmt.Println(errors.New(fmt.Sprintf("unknown instruction `%s` found", command[0])))
				return
			}
		}
	}
	return
}
