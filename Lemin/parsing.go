package lemin

import (
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

var (
	colony AntsFarm
	room   Room
	link   Link
)

func Parsing(content string) (string, bool) {
	content = strings.ReplaceAll(content, "\r\n", "\n")
	lines := strings.Split(content, "\n")
	// START . END :
	countStart := 0
	countEnd := 0
	for a := 0; a < len(lines)-1; a++ {
		line := strings.TrimSpace(lines[a])
		if line == "##start" {
			countStart++
		}
		if line == "##end" {
			countEnd++
		}
	}
	if countStart != 1 || countEnd != 1 {
		return "ERROR: invalid data format, Problem with ##start or ##end tag", false
	}
	//----------------------------------------------------

	// NUMBER OF ANTS :
	for n := 0; n < len(lines)-1; n++ {
		line := strings.TrimSpace(lines[n])

		if line == "##start" || line == "##end" {
			return "ERROR: invalid data format, ##start or ##end tag appears before the number of ants", false
		}
		if line == "" || line[0] == '#' {
			continue
		}
		numberAnts, er := strconv.Atoi(line)
		if er != nil || numberAnts <= 0 {
			return "ERROR: invalid data format, , invalid number of Ants", false
		}
		colony.numberAnts = numberAnts
		lines = lines[n+1:]
		break
	}

	//---------------------------------------------------------------
	//ROOMS:
	roomRegex := "^[^-\\s][^-]* [0-9]+ [0-9]+$"
	linkRegex := "^[^-]+-[^-]+$"

	index := 0
	for j := 0; j < len(lines); j++ {
		line := strings.TrimSpace(lines[j])

		if line == "##start" || line == "##end" {
			if j+1 < len(lines) {
				check, _ := regexp.MatchString(roomRegex, lines[j+1])
				if check {
					r := strings.Split(lines[j+1], " ")
					if line == "##start" {
						colony.start = r[0]
					}
					if line == "##end" {
						colony.end = r[0]
					}
					continue
				} else {
					return "ERROR: invalid data format, no room defined for #start or #end", false
				}
			} else {
				return "ERROR: invalid data format, no room defined for #start or #end", false
			}
		}
		if line == "" || line[0] == '#' {
			continue
		}
		if line[0] == '#' || line[0] == 'L' {
			return "ERROR: invalid data format, this is begin with L or #", false
		}
		matchLink, _ := regexp.MatchString(linkRegex, line)
		if matchLink {
			index = j
			break
		}
		matchRoom, _ := regexp.MatchString(roomRegex, line)
		if !matchRoom {
			return "ERROR: invalid data format, invalid room", false
		}

		r := strings.Split(line, " ")
		name := r[0]
		coorX, _ := strconv.Atoi(r[1])
		coorY, _ := strconv.Atoi(r[2])

		for _, r := range colony.rooms {
			if r.name == name || (r.x == coorX && r.y == coorY) {
				return "ERROR: invalid data format, room name or coordinates already exist", false
			}
		}
		room = Room{name, coorX, coorY}
		colony.rooms = append(colony.rooms, room)
	}
	//--------------------------------------------------------------------
	//LINKS
	lines = lines[index:]

	for q := 0; q < len(lines); q++ {
		line := strings.TrimSpace(lines[q])
		if line == "##start" || line == "##end" {
			return "Error: invalid data format,  ##start or ##end tag appears in Links part", false
		}
		if line == "" || line[0] == '#' {
			continue
		}
		matchLink, _ := regexp.MatchString(linkRegex, line)
		if !matchLink {
			return "ERROR: invalid data format, invalid link", false
		}
		l := strings.Split(line, "-")

		R1 := l[0]
		R2 := l[1]

		// check rooms exist
		roomR1 := false
		roomR2 := false
		for _, r := range colony.rooms {
			if r.name == R1 {
				roomR1 = true
			}
			if r.name == R2 {
				roomR2 = true
			}
		}
		if !roomR1 || !roomR2 {
			return "ERROR: invalid data format, invalid link room does not exist", false
		}

		// check duplicate link (R1-R2 or R2-R1)
		for _, lnk := range colony.links {
			if (lnk.room1 == R1 && lnk.room2 == R2) || (lnk.room1 == R2 && lnk.room2 == R1) {
				return "ERROR: duplicate link", false
			}
		}

		// si pas d'erreur → continuer le parsing
		colony.links = append(colony.links, Link{R1, R2})

	}

	return "", true
}
 
