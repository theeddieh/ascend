package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/theeddieh/ascend/db"
)

func main() {
	// this should accept a single argument on the command line
	// pointing to the path of the input file
	if len(os.Args) < 2 {
		fmt.Printf("No filepath specified\n")
		fmt.Printf("Usage:\n")
		fmt.Printf("  ./memdb <path to input file>\n")
		return
	}
	infile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", os.Args[1])
		return
	}
	defer infile.Close()
	fileScanner := bufio.NewScanner(infile)

	d := db.New()
	for fileScanner.Scan() {
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
			fmt.Printf("unknown instruction `%s` found\n", command[0])
			return
		}
	}
	return
}
