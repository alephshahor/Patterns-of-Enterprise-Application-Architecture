package main

import (
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/cmd"
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/db"
)

func main() {
	cmd.Execute()

	var err error
	if _, err = db.DB().Exec("SELECT 1"); err != nil {
		panic(err)
	}

}
