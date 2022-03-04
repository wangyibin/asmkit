package main

import (
	"log"

	"github.com/wangyibin/asmkit"
)

func main() {
	err := asmkit.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
