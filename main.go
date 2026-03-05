package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// if !(Parsing(content)){
	// 	fmt.Println("ERROR: invalid data format")
	// 	return
	// }

	Edited, err := Parsing(content)

	if !err {
		fmt.Println("ERROR: invalid data format")
		return
	}
}
