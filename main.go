package main

import (
	"log"

	"github.com/Kuri-su/confSyncer/cmd"
)

func main() {
	log.SetFlags(log.Ldate | log.Lshortfile)
	cmd.Execute()
}
