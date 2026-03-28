package lemin

import (
	"fmt"
)

var paths [][]string
var explorationCount int

const (
	MAX_PATHS        = 100
	MAX_PATH_LENGTH  = 50
	MAX_EXPLORATIONS = 100000
)

func Lemin() {
	paths = [][]string{}
	explorationCount = 0

	firstPath := FindFirstPathBFS(colony.start, colony.end)
	if firstPath == nil {
		fmt.Println("ERROR: invalid data format, no path between start and end")
		return
	}
	paths = append(paths, firstPath)

	for _, edge := range graph[colony.start] {
		if explorationCount >= MAX_EXPLORATIONS || len(paths) >= MAX_PATHS {
			break
		}
		path := []string{colony.start, edge}
		visited := map[string]int{colony.start: 0, edge: 1}
		LoopCheckPaths(edge, path, visited)
	}

	for i := 0; i < len(paths); i++ {
		for j := i + 1; j < len(paths); j++ {
			if len(paths[i]) > len(paths[j]) {
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
	}

	SimulateTurns(paths, colony.numberAnts)
}

func FindFirstPathBFS(start, end string) []string {
	if start == end {
		return []string{start}
	}
	queue := [][]string{{start}}
	visited := map[string]bool{start: true}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		node := path[len(path)-1]

		if len(path) > MAX_PATH_LENGTH {
			continue
		}
		for _, next := range graph[node] {
			if visited[next] {
				continue
			}
			newPath := append([]string{}, path...)
			newPath = append(newPath, next)

			if next == end {
				return newPath
			}
			visited[next] = true
			queue = append(queue, newPath)
		}
	}

	return nil
}

func LoopCheckPaths(node string, path []string, visited map[string]int) {
	explorationCount++

	if explorationCount >= MAX_EXPLORATIONS || len(paths) >= MAX_PATHS || len(path) > MAX_PATH_LENGTH {
		return
	}

	if node == colony.end {
		finalPath := append([]string{}, path...)
		if !pathExists(finalPath) {
			paths = append(paths, finalPath)
		}
		return
	}

	for _, next := range graph[node] {
		if explorationCount >= MAX_EXPLORATIONS {
			return
		}

		if _, seen := visited[next]; seen {
			continue
		}

		visited[next] = len(path)
		path = append(path, next)

		LoopCheckPaths(next, path, visited)
		path = path[:len(path)-1]
		delete(visited, next)
	}
}

func pathExists(newPath []string) bool {
	for _, p := range paths {
		if len(p) != len(newPath) {
			continue
		}
		same := true
		for i := range p {
			if p[i] != newPath[i] {
				same = false
				break
			}
		}
		if same {
			return true
		}
	}
	return false
}
