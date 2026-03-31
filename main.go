package main

import (
	"fmt"
	"os"
	"strings"

	lemin "LEM-IN/Lemin"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filename>")
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR: cannot read file:", err)
		return
	}

	content := string(data)

	errMsg, ok := lemin.Parsing(content)
	if !ok {
		fmt.Println(errMsg)
		return
	}

	lines := strings.Split(strings.TrimSpace(content), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Println(line)
	}
	fmt.Println()

	lemin.Graph()
}
