package main

import (
	"POS-SRI/db"
	"POS-SRI/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":3333"))
}
