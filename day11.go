package main


import (
	"fmt"
	"strings"
)

// paths := [["you"]]
// graph["you"] = ("bbb", "ccc")
// vertex := "you"
//
// paths := [["you", "bbb"], ["you", "bbb"]]



func walkToNextVertex(graph map[string][]string, paths [][]string) ([][]string) {
	for _, path := range paths {
		vertex := path[len(path)-1] // take last element
		for _, nextVertex := graph[vertex] {
			paths = append(paths, nextVertex)
		}
	}
	return paths
}

func day11part1(filename string) {
	input := FileToArray(filename)

	// make a map of input
	graph := make(map[string][]string)

	for _, line := range input {
		parts := strings.Split(line, " ")
		vertex := parts[0][:len(parts[0])-1] // cut off the ":"
		graph[vertex] = parts[1:]
	}
	fmt.Println(graph)

	// find paths
	countPath := 0
	visited := map[string]bool
	for _, vertex := range graph["you"] {
		fmt.Println(vertex)
		walkToNextVertex(graph, vertex)
	}


	fmt.Println(countPath)
}
