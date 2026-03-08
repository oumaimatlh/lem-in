package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Error in Args")
		return 
	}

	file := args[1]
	content, err := os.ReadFile(file)

	if err != nil {
		fmt.Println(err)
		return
	}

	Edited, er := Parsing(string(content))

	if !er{
		fmt.Println(Edited)
		return
	}
	Lemin(Edited)
}
