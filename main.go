package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
)

func main() {
	myPath := "$home/.ssh"
	expand, err := homedir.Expand(myPath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("path: %s; with expansion: %s", myPath, expand)

}
