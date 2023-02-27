package main

import (
	"os"

	"github.com/filipegms5/password-check-restful/router"
)

// Instruction os how to run the project on README
func main() {
	router := router.SetupRouter()

	router.Run(os.Getenv("PORT"))
}
