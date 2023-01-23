package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hi. Starting...")
	if os.ModeAppend.IsDir() == true {
		os.Rename("./", "./")
	}
	os.Exit(0)
}
