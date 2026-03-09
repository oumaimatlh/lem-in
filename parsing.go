package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type AntsFarm struct {
	numberAnts int
	rooms      []Room
	links      []Link
	start      string
	end        string
}
type Room struct {
	name string
	x    int
	y    int
}
type Link struct {
	room1 string
	room2 string
}

var colony AntsFarm
var room Room
var link Link

func Parsing(content string) (string, bool) {

	lines := strings.Split(content, "\r\n")

	countStart := 0
	countEnd := 0
	for k := 0; k < len(lines)-1; k++ {
		line := strings.TrimSpace(lines[k])
		if line == "##start" {
			countStart++
		}
		if line == "##end" {
			countEnd++
		}
	}
	if countStart != 1 || countEnd != 1 {
		return "Two Tags  Start or End", false
	}
	if countStart == 0 || countEnd == 0 {
		return "NoT Start || End Tag", false
	}

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
				return "ERROR: invalid data format, Tag Start/End before Number of Ants", false
			}
			continue
		}
		numberAnts, er := strconv.Atoi(line)
		if er != nil || numberAnts <= 0 {
			return "ERROR: invalid data format, , invalid number of Ants", false
		}
		colony.numberAnts = numberAnts
		lines = lines[i+1:]
		break
	}

	//------------------------------------------//

	roomRegex := "^[a-zA-Z0-9]+ [0-9]+ [0-9]+$"
	linkRegex := "^[a-zA-Z0-9]+-[a-zA-Z0-9]+$"

	for j := 0; j < len(lines); j++ {
		line := strings.TrimSpace(lines[j])

		if line == "" {
			continue
		}

		if line[0] == '#' && line[1:]!= "#start" && line[1:] != "#end"{
			continue
		}
		matchRoom, _ := regexp.MatchString(roomRegex, line)
		if !matchRoom {
			matchLink, _ := regexp.MatchString(linkRegex, line)
			if !matchLink {
				return "Error Room", false
			}

			if j+1 < len(lines)  && (line == "##start" || line == "##end"){
				check, _ := regexp.MatchString(roomRegex, lines[j+1])
				if check {
					r := strings.Split(lines[j+1], " ")
					if line == "##start" {
						colony.start = r[0]
					}
					if line == "##end" {
						colony.end = r[0]
					}
				}else{
					return "No room for Start/End", false
				}

			}
			//"The rooms => Selections"
			break
		}
		r := strings.Split(line, " ")
		coorX, _ := strconv.Atoi(r[1])
		coorY, _ := strconv.Atoi(r[2])
		room = Room{r[0], coorX, coorY}

		for _, r := range colony.rooms {
			if room.name == r.name || (room.x == r.x && room.y == r.y) {
				return "this room is double coord or name ", false
			}
		}

		colony.rooms = append(colony.rooms, room)
	}

	fmt.Println(colony.numberAnts, colony.rooms, colony.start, colony.end)

	return "", true
}
