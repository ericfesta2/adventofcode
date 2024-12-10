package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var graph map[string][]string = map[string][]string{}

func contains(list []string, target string) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == target {
			return true
		}
	}

	return false
}

func immutableAppendToVisited(mp map[string]bool, key string) map[string]bool {
	copy := map[string]bool{}

	for k, v := range(mp) {
		copy[k] = v
	}

	copy[key] = true

	return copy
}

func addToGraph(line string) {
	pageNums := strings.Split(line, "|")

	if _, fnd := graph[pageNums[0]]; !fnd {
		graph[pageNums[0]] = []string{pageNums[1]}
	} else {
		graph[pageNums[0]] = append(graph[pageNums[0]], pageNums[1])
	}
}

func verifyPageOrder(pageOrder []string, offsetFromEnd int, currGraphNode string, visited map[string]bool) bool {
	if offsetFromEnd == len(pageOrder) {
		return true
	}

	if visited[currGraphNode] {
		return false
	}

	newVisited := immutableAppendToVisited(visited, currGraphNode)
	adjacent := graph[currGraphNode]

	// Ensure that one of the pages after the current in pageOrder doesn't come before the current page
	for i := 0; i < offsetFromEnd; i++ {
		if currGraphNode == pageOrder[i] {
			return false
		}
	}

	newOffsetFromEnd := offsetFromEnd

	// If this isn't the target node, keep looking for the same one
	if currGraphNode == pageOrder[offsetFromEnd] {
		newOffsetFromEnd++
	}

	for i := 0; i < len(adjacent); i++ {
		if verifyPageOrder(pageOrder, newOffsetFromEnd, adjacent[i], newVisited) {
			return true
		}
	}

	return false
	// return verifyPageOrder(pageOrder, newOffsetFromEnd, pageOrder[newOffsetFromEnd], newVisited)
}

func main() {
	file, err := os.Open("input.txt")

    if err != nil {
        panic("Error opening file: " + err.Error())
    }

    defer file.Close()

	scanner := bufio.NewScanner(file)

	middlePageSum := 0

    for scanner.Scan() {
        line := scanner.Text()

		if strings.Contains(line, "|") {
			addToGraph(line)
		} else if strings.Contains(line, ",") {
			pageOrder := strings.Split(line, ",")

			// Assume all pages in each page order list exist in the built graph
			if verifyPageOrder(pageOrder, 1, pageOrder[len(pageOrder) - 1], map[string]bool{}) {
				if middlePage, convErr := strconv.Atoi(pageOrder[len(pageOrder) / 2]); convErr == nil {
					middlePageSum += middlePage
				} else {
					fmt.Println("Unexpected error with page number formatting: " + convErr.Error())
				}
			}
		}
    }

	fmt.Println(middlePageSum)
}
