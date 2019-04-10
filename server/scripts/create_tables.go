package main

import (
	"io/ioutil"
	"os"

	"github.com/jmoiron/sqlx"
	// sqlite necessary
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	sqlite, err := sqlx.Connect("sqlite3", "../circles.db")
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 1 && os.Args[1] != "all" {
		file, err := ioutil.ReadFile("data/" + os.Args[1])
		if err != nil {
			panic(err)
		}
		sqlite.MustExec(string(file))

	} else {
		dir, err := ioutil.ReadDir("data")
		if err != nil {
			panic(err)
		}

		for _, fileInfo := range dir {
			file, err := ioutil.ReadFile("data/" + fileInfo.Name())
			if err != nil {
				panic(err)
			}
			sqlite.MustExec(string(file))
		}
	}
}
