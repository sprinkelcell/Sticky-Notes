package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"

	"github.com/sprinkelcell/Sticky-Notes/cmd"
	"github.com/sprinkelcell/Sticky-Notes/db"
)

func main() {
	dir, err := homedir.Dir()
	if err != nil {
		handleError(err, "Something Went wrong please try again")
		os.Exit(1)
	}
	dir = dir + "/notes.db"
	handleError(db.Init(dir), "Something Went wrong please try again")
	handleError(cmd.RootCmd.Execute(), "Something Went wrong please try again")
	defer handleError(db.Exit(), "")
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		log.Fatal(err)
	}
}
