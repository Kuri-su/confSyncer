package main

import (
	"log"

	"github.com/Kuri-su/confSyncer/pkg/confSyncer"
)

func main() {
	log.SetFlags(log.Ldate | log.Lshortfile)
	confSyncer.Execute()
}
