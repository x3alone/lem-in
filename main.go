package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(" Error: wrong number of arguments")
		return
	}
	fileName := os.Args[1]
	colony, err := ReadMap(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	Graph := &Graph{room: make(map[string][]string)}
	for _, room := range colony.rooms {
		Graph.room[room.name] = []string{}
		for _, link := range room.links {
			Graph.room[room.name] = append(Graph.room[room.name], link.name)
		}
	}

	start := colony.start.name
	end := colony.end.name
	antNum := colony.ants

	paths := []string{}
	allPath := FindPaths(Graph, start, end, paths, start)
	solutions := Solutions(ValidPaths(allPath, start, end), antNum)
	solution := solutions[0]

	AntBalancing(solution, end, antNum)
}
