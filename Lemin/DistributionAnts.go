package lemin

func DistributeAnts(paths [][]string, n int) []int {
	dist := make([]int, len(paths))
	for a := 0; a < n; a++ {
		best := -1
		bestVal := -1
		for i, p := range paths {
			val := len(p) - 1 + dist[i] + 1
			if best == -1 || val < bestVal {
				bestVal = val
				best = i
			}
		}
		dist[best]++
	}
	return dist
}