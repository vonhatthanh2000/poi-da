package main

import (
	"layerg-poi-da/internal/initialize"
)

// Seeding default data
func main() {

	initialize.LoadConfig()
	initialize.InitLogger()
	initialize.InitDB()

}
