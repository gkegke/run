/*

  Command related functions and variables.

  - Vars

    COMMANDS []string - list of valid commands found in the command directory

  - Functions

    getCommands(path string) => []string
      - searches a path, and returns all the valid command directories in that path

*/

package main

import (
    "fmt"
    "os"
)

var COMMANDS []string = getCommands(COMMAND_DIRECTORY)

func getCommands(path string) []string {

    f, err := os.Open(path)

    if err != nil {
        fmt.Println(err)
    }

    names, err := f.Readdirnames(0)
    if err != nil {
        fmt.Println(err)
    }

    return names
}