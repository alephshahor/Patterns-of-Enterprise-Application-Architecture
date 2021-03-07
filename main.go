package main

import (
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/cmd"
)

func main() {
	cmd.Execute()
	db = DB()

	var err error
	if _, err = db.Exec("SELECT 1"); err != nil {
		panic(err)
	}

}
