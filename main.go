package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/rafalmp/task/cmd"
	"github.com/rafalmp/task/db"
)

func main() {
	home, err := homedir.Dir()
	checkErr(err)
	dbPath := filepath.Join(home, "tasks.db")
	checkErr(db.Init(dbPath))
	cmd.Execute()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
