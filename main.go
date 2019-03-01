package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/theeddieh/ascend/db"
)

// accepts a single argument pointing to the path of the input file
// reads file a line at a time
// executes database commands as they are read
func main() {

	if len(os.Args) < 2 {
		fmt.Printf("No filepath specified\n")
		fmt.Printf("Usage:\n")
		fmt.Printf("  ./memdb <path to input file>\n")
		return
	}

	var debug bool
	if len(os.Args) > 2 && os.Args[2] == "-v" {
		debug = true
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
			if debug {
				fmt.Println("-----------")
			}
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
