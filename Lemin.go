package main

import (
	"fmt"
	"slices"
)

var paths [][]string

func Lemin() {
	paths = [][]string{}

	for _, edge := range graph[colony.start] {
		path := []string{colony.start, edge}
		check := []string{colony.start, edge}
		LoopCheckPaths(edge, path, check)
	}

	if len(paths) == 0 {
		fmt.Println("ERROR: invalid data format, no path between start and end")
		return
	}
// // Je veux ordrer les paths  selon distance
	for i := 0; i < len(paths); i++ {
		m := paths[i]
		for j := i+1; j < len(paths); j++{
			if len(paths[i]) > len(paths[j]){
				paths[i]=paths[j]; paths[j]=m
				m = paths[i]
			}
		}
		
	}

//Maintenant on doit Chercher meilleur Solution afin dim le nombre de Tours 

}

func LoopCheckPaths(node string, path []string, check []string) {
	//Backtracking / Récursivité
	if node == colony.end {
		nvPath := make([]string, len(path))
		copy(nvPath, path)
		paths = append(paths, nvPath)
		return
	}

	for _, edge := range graph[node] {
		if slices.Contains(check, edge) {
			continue
		}
		nvPath := make([]string, len(path)+1)
		copy(nvPath, path)
		nvPath[len(path)] = edge

		nvCheck := make([]string, len(check)+1)
		copy(nvCheck, check)
		nvCheck[len(check)] = edge

		LoopCheckPaths(edge, nvPath, nvCheck)
	}
}
