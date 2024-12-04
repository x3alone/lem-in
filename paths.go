package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// goes thru the graph provided explore valid paths using recursion building paths from scratch, all paths get returned in a 2d string
func FindPaths(g *Graph, current string, end string, paths []string, fixedStart string) [][]string {
	var validPaths [][]string
	if current == end {
		validPaths = append(validPaths, append([]string{}, paths...))
	}
	for _, value := range g.room[current] {
		if !slices.Contains(paths, value) && value != fixedStart { // insuring that it isnt visited twice in a single path
			newpaths := append(paths, value)
			validPaths = append(validPaths, FindPaths(g, value, end, newpaths, fixedStart)...)

		}
	}
	return validPaths
}

// figuers out the path that have an end and distination returns them
func ValidPaths(Paths [][]string, start string, end string) [][]string {
	var result [][]string
	for i := 0; i < len(Paths); i++ {
		if Paths[i][len(Paths[i])-1] == end {
			result = append(result, Paths[i][:len(Paths[i])-1])
		}
	}
	return result
}

// responsable to track an ant and print its movment along the path till it reaches the end
func AppendPaths(ant int, path []string, end string, printres [][]string, index int) {
	j := 0
	path = append(path, end)
	for i := index; i < len(printres); i++ {
		if j < len(path) {
			printres[i] = append(printres[i], "L"+strconv.Itoa(ant)+"-"+(path[j]))
			j++
		}
	}
}

// distributes ants thru multiple paths traversing to the destination end
// calculates the steps each path requires using Steps, the total waves needed using CountSteps, and assigns ants to paths
// using the LowestCOunt function. the function tracks the movmment of ants using AppendPaths
// cmp stores index value for each ant that in a path, so we can know on what wave (turn) we should assign the next ant
func AntBalancing(solution [][]string, end string, antNb int) []int {
	cmp, steps := Steps(solution)
	waveNumber := CountSteps((steps), (antNb), (len(solution)))
	printResult := make([][]string, waveNumber)
	numberOfAnts := []int{}
	antDisrbution := []int{}
	antINdex := 1
	for i := 0; i < len(solution); i++ {
		if i == 0 {
			numberOfAnts = append(numberOfAnts, len(solution[i]))
			antDisrbution = append(antDisrbution, 1)

		} else {
			numberOfAnts = append(numberOfAnts, len(solution[i]))
			antDisrbution = append(antDisrbution, 0)
		}
	}
	for {
		lowest := LowestCOunt(numberOfAnts, antDisrbution)

		AppendPaths(antINdex, solution[lowest], end, printResult, cmp[lowest])
		cmp[lowest]++
		antDisrbution[lowest]++
		antNb--
		antINdex++
		if antNb == 0 {
			break
		}
	}
	for _, value := range printResult {
		fmt.Println(strings.Join(value, " "))
	}
	return antDisrbution
}
