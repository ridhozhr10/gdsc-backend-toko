package main

import (
	"log"
	"os"

	"github.com/ridhozhr10/gdsc-backend-toko/engine"
)

func main() {
	if err := engine.Run(":3000"); err != nil {
		log.New(os.Stdout, "Main: ", log.Ldate|log.Ltime|log.Lshortfile).Print(err)
	}
}
