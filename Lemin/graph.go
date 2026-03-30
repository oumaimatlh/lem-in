package lemin


var graph map[string][]string

func Graph() {
	graph = map[string][]string{}

	for _, r := range colony.rooms {
		for _, l := range colony.links {
			if r.name == l.room1 {
				graph[r.name] = append(graph[r.name], l.room2)
				continue
			}
			if r.name == l.room2 {
				graph[r.name] = append(graph[r.name], l.room1)
				continue
			}
		}
	}

	Lemin()
}