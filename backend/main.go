package main

import (
	"fmt"
	"server/db"
)

func main() {
	db.InitDB()
  fmt.Println("Initialization complete")
}
