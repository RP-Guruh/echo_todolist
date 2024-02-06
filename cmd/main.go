package main

import (
	"todolist/db"
	"todolist/routes"
)

func main() {

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
