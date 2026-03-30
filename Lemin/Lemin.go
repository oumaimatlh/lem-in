package lemin

import (
	"fmt"
)

var paths [][]string

func Lemin() {
	paths = [][]string{}
	for _, edge := range graph[colony.start] {
		path := []string{colony.start, edge}
		visited := map[string]int{colony.start: 0, edge: 1}
		LoopCheckPaths(edge, path, visited)
	}
	if len(paths) == 0 {
		fmt.Println("ERROR: invalid data format, no path between start and end")
		return
	}
	SimulateTurns(paths, colony.numberAnts)
}

func LoopCheckPaths(node string, path []string, visited map[string]int) {
	if node == colony.end {
		finalPath := append([]string{}, path...)
		paths = append(paths, finalPath)
		return
	}
	for _, next := range graph[node] {
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
