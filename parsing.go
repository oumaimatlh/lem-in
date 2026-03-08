package main

import (
	//	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Parsing(content string) (string, bool) {

	lines := strings.Split(content, "\r\n")

	//Check a number of Ants :
	for i := 0; i < len(lines)-1; i++ {
		line := strings.TrimSpace(lines[i])
		//empty line
		if line == "" {
			continue
		}
		//Comment
		if line[0] == '#' {
			if line == "##start" || line == "##end" {
				return "ERROR: invalid data format, Tag Start before Number of Ants", false
			}
			continue
		}
		numberAnts, er := strconv.Atoi(line)
		if er != nil || numberAnts <= 0 {
			return "ERROR: invalid data format, , invalid number of Ants", false
		}
		lines = lines[i+1:]
		break
	}

	//------------------------------------------//

	room := "^[a-zA-Z0-9]+ [0-9]+ [0-9]+$"
	//rooms := []string{}

	for j := 0; j < len(lines); j++ {
		line := strings.TrimSpace(lines[j])

		if line == "" { 
			continue
		}
		if line[0] == '#' {
			continue
		}
		match, _ := regexp.MatchString(room, line)
		if !match {
			
		}
		
		

	}
	return "", true
}
