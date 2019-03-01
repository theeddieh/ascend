# Rollback approaches

1. Need to revert DB to state prior to last command
    1. backwards:
        1. keep a stack of "undo commands"
        2. undo last command by issuing the "opposite", e.g.
            1. new key:    write(k,v)   -> delete(k)
            2. update key: write(k, v1) -> write(k, v0)
            3. delete key: delete(k)    -> write(k, v0)
        3. for all, need to read(k) before each operation to capture state
        4. but what if the db tracked the history of values for each key?
            1. e.g. `map[string][]string`, where `[]string` is a stack of values
            2. by the definition of the input format, we can't have an empty value, so `""` can be used to represent the deleted state
            3. then just track the transactions as a history of keys, and pop as we go
            4. catches:
                1. what if we roll back to an empty slice?
                2. tricky to test, requires multiple commands to test a single case
                3. I did not include input file parsing as part of `db`, but that, or spinning it of into a separate package, would make feeding long/complex command sequence to the tests much easier
    2. forwards:
        1. keep a queue of commands issued
        2. on a rollback:
            1. replay the the commands, skipping the latest
        3. that's potentially a lot of writes!
        4. from the exercise brief: `Your program should run the commands in the file line by line`
            1. so we aren't:
                1. treating the input file as the transaction log
                2. reading the whole file in first so we can just identify and skip rolledback commands
                    ```
                        WRITE key-0 val-1
                        WRITE key-1 val-3
                        WRITE key-2 val-4
                      ┌─DELETE key-1
                      │┌DELETE key-0
                      │└ROLLBACK
                      └─ROLLBACK
                        WRITE key-2 val-8
                        DELETE key-0
                        PRINT
                    ```
                3. could be implemented as a stack, push commands on and pop when you get to a rollback
                4. if we _did_ go down this road, could also check for keys deleted and never re-written
                5. but this is a stream process, not batch!








# watch out!

* what if we start with a ROLLBACK?
* a DELETE on an empty should be fine, what if we followed that with a ROLLBACK?
    * well, we should only be able to rollback successful db writes/deletes, otherwise what does that even mean?
    * so rollback need to know the difference between a delete that was on a real key (which we track by considering its value to be `""`, and a DELETE that was _NOT_ preceeded by a WRITE)
* commands should be process as a stream, not a batch, so we can't optimize along those lines.
* input is "whitespace delimited" so `""` is not a valid key OR value

# Enhancements

* print out live updates to db
* save db to disk on a hard exit, give options to restore/resume
    * interrupt or input file typo
        1. log error
        2. write out "restore file" containing:
            * input file name
            * line number of error/last line processes
            * db dump
        3. on next run:
            1. check for new "-restore" flag pointing to backup file
            2. parse file to rebuild db
            3. open logged input file and pick up where left off
            4. note: no history is preserved, it's possible that subsequent commands could try to rollback to before the restore
            5. actually, go prints out structs very nicely, we _could_  write out the entire state and history of the db and then parse it back in
* setup CI/CD
* refactor to make it easy to test large inputs
