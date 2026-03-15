package main

import (
	"fmt"
	"slices"
)

var paths [][]string

func Lemin() {
	paths = [][]string{}

	for _, link := range graph[colony.start] {
		path := []string{colony.start, link}
		check := []string{colony.start, link}
		LoopCheckPaths(link, path, check)
	}

	if len(paths) == 0 {
		fmt.Println("ERROR: invalid data format, no path between start and end")
		return
	}
	// // Je veux ordrer les paths  selon distance

	fmt.Println(paths)

}

func LoopCheckPaths(node string, path []string, check []string) {
	//Backtracking / Récursivité
	if node == colony.end {
		newPath := make([]string, len(path))
		copy(newPath, path)
		paths = append(paths, newPath)
		return
	}

	for _, l := range graph[node] {
		if slices.Contains(check, l) {
			continue
		}
		newPath := make([]string, len(path)+1)
		copy(newPath, path)
		newPath[len(path)] = l

		newCheck := make([]string, len(check)+1)
		copy(newCheck, check)
		newCheck[len(check)] = l

		LoopCheckPaths(l, newPath, newCheck)
	}
}
