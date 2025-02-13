package main

import (
	"layerg-poi-da/global"
	"layerg-poi-da/internal/initialize"
)

// Seeding default data
func main() {

	initialize.LoadConfig()
	initialize.InitLogger()
	initialize.InitDB()

	// Comment lines that you don't need
	seedUsers(global.Cdb)
}
