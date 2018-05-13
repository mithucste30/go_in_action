package main

import (
	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"log"
	"os"
	"github.com/goinaction/code/chapter2/sample/search"
)

// init called prior to main
func init()  {
	// change the device for logging
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program
func main()  {
	// perform the search for the specified term
    search.Run("president")
}
