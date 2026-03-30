package lemin

import (
	"fmt"
	"strings"
)

func SimulateTurns(paths [][]string, numAnts int) {
	selected := SelectBestPaths(paths)
	distribution := DistributeAnts(selected, numAnts)

	queues := make([][]int, len(selected))
	antPath := map[int]int{}
	antStep := map[int]int{}
	antDone := map[int]bool{}

	antID := 1
	for pathIdx, count := range distribution {
		for i := 0; i < count; i++ {
			queues[pathIdx] = append(queues[pathIdx], antID)
			antPath[antID] = pathIdx
			antStep[antID] = 0 
			antDone[antID] = false
			antID++
		}
	}

	qPtr := make([]int, len(selected)) 
	active := []int{}                  
	done := 0

	for done < numAnts {
		turnMoves := []string{}
		occupied := map[string]bool{}

		stillActive := []int{}
		for _, id := range active {
			pi := antPath[id]
			path := selected[pi]
			nextStep := antStep[id] + 1

			if nextStep >= len(path) {
				antDone[id] = true
				done++
				continue
			}

			nextRoom := path[nextStep]
			isEnd := nextRoom == colony.end

			if !isEnd && occupied[nextRoom] {
				stillActive = append(stillActive, id)
				continue
			}

			if !isEnd {
				occupied[nextRoom] = true
			}
			antStep[id] = nextStep
			turnMoves = append(turnMoves, fmt.Sprintf("L%d-%s", id, nextRoom))

			if isEnd {
				antDone[id] = true
				done++
			} else {
				stillActive = append(stillActive, id)
			}
		}

		for pi := range selected {
			if qPtr[pi] >= len(queues[pi]) {
				continue
			}
			path := selected[pi]
			if len(path) < 2 {
				continue
			}
			firstRoom := path[1]
			isEnd := firstRoom == colony.end

			if !isEnd && occupied[firstRoom] {
				continue 
			}

			id := queues[pi][qPtr[pi]]
			qPtr[pi]++

			if !isEnd {
				occupied[firstRoom] = true
			}
			antStep[id] = 1
			turnMoves = append(turnMoves, fmt.Sprintf("L%d-%s", id, firstRoom))

			if isEnd {
				antDone[id] = true
				done++
			} else {
				stillActive = append(stillActive, id) 
			}
		}

		active = stillActive

		if len(turnMoves) > 0 {
			fmt.Println(strings.Join(turnMoves, " "))
		}
	}
}

