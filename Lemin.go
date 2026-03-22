package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
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
	// ordrer les paths  selon distance
	for i := 0; i < len(paths); i++ {
		m := paths[i]
		for j := i + 1; j < len(paths); j++ {
			if len(paths[i]) > len(paths[j]) {
				paths[i] = paths[j]
				paths[j] = m
				m = paths[i]
			}
		}

	}

	fmt.Println(paths)
	fmt.Println()
	// //Maintenant on doit Chercher meilleur Solution afin dim le nombre de Tours

	path := strings.Join(paths[0], "-")
	nameAnt := "L1"
	PathAnts := map[string][]string{
		path: {nameAnt},
	}

	for n := 2; n <= colony.numberAnts; n++ {
		nameAnt = "L" + strconv.Itoa(n)
		for _, path := range paths {
			item := PathAnts[strings.Join(path, "-")]
			pathTime := len(path) + len(item)
			check := false
			for _, p := range paths {
				timeP := len(p) + len(PathAnts[strings.Join(p, "-")])
				if pathTime > timeP {
					//Verefication d paths
					found := false
					for c, antsOnC := range PathAnts {
						it1 := p[1 : len(p)-1]
						it2 := strings.Split(c, "-")
						it2 = it2[1 : len(it2)-1]

						rangNouvelleFourmi := len(PathAnts[strings.Join(p, "-")]) + 1

						for rangC := range antsOnC {
							rangFourmiC := rangC + 1
							for i, n1 := range it1 {
								tourN1 := i + rangNouvelleFourmi
								for j, n2 := range it2 {
									tourN2 := j + rangFourmiC
									if n1 == n2 && tourN1 == tourN2 {
										found = true
									}
								}
							}
						}
					}
					if found {
						continue
					}
					check = true
					PathAnts[strings.Join(p, "-")] = append(PathAnts[strings.Join(p, "-")], nameAnt)
					break
				}
			}

			if !check {
				cle := strings.Join(paths[0], "-")
				min := len(paths[0]) + len(PathAnts[cle])
				for c, v := range PathAnts {
					time := len(strings.Split(c, "-")) + len(v)
					if min > time {
						cle = c
						min = len(strings.Split(c, "-")) + len(v)
					}
				}
				PathAnts[cle] = append(PathAnts[cle], nameAnt)
			}

			break
		}
	}

	fmt.Print(PathAnts)

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
