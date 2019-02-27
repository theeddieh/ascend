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
