package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
)

type Room struct {
	name    string
	x       string
	y       string
	links   []*Room
	isStart bool
	isEnd   bool
}

type Colony struct {
	ants  int
	rooms map[string]*Room
	start *Room
	end   *Room
}

type Graph struct {
	room map[string][]string
}

func (Graph *Graph) Edges(room1 string, room2 string) {
	Graph.room[room1] = append(Graph.room[room1], string(room2))
	Graph.room[room2] = append(Graph.room[room2], room1)
}

func NotVisited(str string, slice [][]string) bool {
	for _, value := range slice {
		if slices.Contains(value, str) {
			return true
		}
	}

	return false
}
//this function is the responsable for making all the solutions possible in a 3d table a slice of slices of strings containing all solutions 
func Solutions(validPaths [][]string, antNb int) [][][]string {
	noDup := make(map[string]struct{})
	solutions := [][][]string{}
	if len(validPaths) == 1 {
		return [][][]string{validPaths}
	}
	for i := 0; i < len(validPaths); i++ {
		solution := [][]string{}
		solution = append(solution, validPaths[i])
		for j := 0; j < len(validPaths); j++ {
			notCommun := true
			if i != j {
				for k := 0; k < len(validPaths[j]); k++ {
					if slices.Contains(validPaths[i], validPaths[j][k]) || NotVisited(validPaths[j][k], solution) {
						notCommun = false
						break
					} else {
						continue
					}
				}
				if notCommun {
					solution = append(solution, validPaths[j])
					slices.SortFunc(solution,
						func(a, b []string) int {
							return cmp.Compare(len(a), len(b))
						})
					fn := fmt.Sprintf("%v", solution)

					if _, found := noDup[fn]; !found {

						solutions = append(solutions, solution)
						noDup[fn] = struct{}{}
					}

				}

			}
		}
	}
	if len(solutions) == 0 {
		return [][][]string{append([][]string{}, validPaths[0])}
	}
	slices.SortFunc(solutions,
		func(a, b [][]string) int {
			_, stp1 := Steps(a)
			_, stp2 := Steps(b)
			return cmp.Compare(CountSteps(stp1, antNb, len(a)), CountSteps(stp2, antNb, len(b)))
		})
	return solutions
}
//handels which paths has the lowest Q by taking the rooms and ants that exist on a single path and comparing it with other paths
func LowestCOunt(rooms []int, ints []int) int {
	var lowestindex int
	lowest := 0
	for i := 0; i < len(rooms); i++ {
		if i == 0 {
			lowest = rooms[i] + ints[i]
			lowestindex = i
		}
		if rooms[i]+ints[i] < lowest {
			lowestindex = i
			lowest = rooms[i] + ints[i]
		}
	}
	if lowest == rooms[0]+ints[0] {
		return 0
	}
	return lowestindex
}
// handels how many steps exist to the end in each solution and ints cmp which w'll use later on  solution function
func Steps(solution [][]string) ([]int, int) {
	steps := 0
	cmp := []int{}
	for i := 0; i < len(solution); i++ {
		steps += len(solution[i])
		cmp = append(cmp, 0)
	}
	return cmp, steps
}
//contains the ocuation that figueres out how many steps are in a Q, by taking the sum of steps and ants devid by path
func CountSteps(steps int, antNb int, SolLen int) int {
	return int(math.Ceil(((float64(steps) + float64(antNb)) / float64(SolLen))))
}