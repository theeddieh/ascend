# AscenDB Readme 

## Overview

AscenDB is a basic in-memory nosql database. 

## Implementation Details

```
type Database struct {
	state   map[string][]string
	history []string
}
```

Rollback implementation details:
1. Keys that are written or deleted are pushed to a the `history` stack
2. Values are also pushed to a stack, one for each key. 
    * deletes are handled internally as writing an empty string to a key's value stack.
        * unless the delete is on a key that was never set
3. On a rollback:
    1. pop the latest changes key name off `history`
    2. pop the latest value off that key's value stack
    3. since the input is whitespace delimited, we cannot have a legitimate empty value, so the empty strings serve as placeholders for deletes

## Usage

### Build
Checkout repo to `<where ever>/src/github.com/theeddieh/ascend`.

Build `ascendb` binary. The `Makefile` will update `$GOPATH`.
```
make
```

### Run
```
$ ./ascendb <path to input file>
```

## Prompt
We'd like to see how you would finish this coding exercise.
The goal is to implement a basic in-memory nosql database.

Each line of the input file contains an instruction followed by a set of *whitespace delimited arguments*.
Your program should run the commands in the file *line by line* and print the final database state to standard out.

You can choose to finish the exercise in any language you prefer.
We have started a portion of the script in golang, as this is the primary language you can expect to work on here at Ascend.

We have already implemented WRITE, DELETE, and PRINT for you.
It is up to you to implement ROLLBACK.

Feel free to change any of the code we have provided.
You will be expected to review your coding exercise at your on-site interview, so be ready to talk through design choices and tradeoffs that you made in your program.

## Example
Input:
```
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
```

Output:
```
key-1 val-3
key-2 val-8
```
