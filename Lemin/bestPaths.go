package lemin

func SelectBestPaths(paths [][]string) [][]string {
	best := [][]string{}
	bestTurns := -1

	Search(0, [][]string{}, map[string]bool{}, paths, &best, &bestTurns)
	return best
}

func Search(start int, current [][]string, usedRooms map[string]bool, paths [][]string, best *[][]string, bestTurns *int) {

	if len(current) > 0 {
		turns := calcTurns(current, colony.numberAnts)

		if *bestTurns == -1 || turns < *bestTurns {
			*bestTurns = turns

			tmp := make([][]string, len(current))
			copy(tmp, current)
			*best = tmp
		}
	}

	for i := start; i < len(paths); i++ {
		path := paths[i]

		conflict := false
		for _, room := range path[1 : len(path)-1] {
			if usedRooms[room] {
				conflict = true
				break
			}
		}
		if conflict {
			continue
		}

		newUsed := make(map[string]bool)
		for k, v := range usedRooms {
			newUsed[k] = v
		}

		for _, room := range path[1 : len(path)-1] {
			newUsed[room] = true
		}

		Search(i+1, append(current, path), newUsed, paths, best, bestTurns)
	}
}

func calcTurns(paths [][]string, n int) int {
	lo, hi := 1, n+len(paths[0])

	for lo < hi {
		mid := (lo + hi) / 2
		total := 0

		for _, p := range paths {
			moves := mid - len(p) + 2
			if moves > 0 {
				total += moves
			}
		}

		if total >= n {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}