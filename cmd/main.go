package main

import (
	"log"

	"asmkit"
)

func main() {
	err := asmkit.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
