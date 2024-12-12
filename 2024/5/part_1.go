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
	for _, el := range list {
		if el == target {
			return true
		}
	}

	return false
}

func deepCopyVisited(mp map[string]bool) map[string]bool {
	copy := map[string]bool{}

	for k, v := range mp {
		copy[k] = v
	}

	return copy
}

func immutableAppendToVisited(mp map[string]bool, key string) map[string]bool {
	copy := deepCopyVisited(mp)

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

func dfs(pageOrder []string, currGraphNode string, target string, visited map[string]bool) bool {
	if currGraphNode == target {
		return true
	}

	adjacent := graph[currGraphNode]

	for i := 0; i < len(adjacent); i++ {
		if _, alreadyVisited := visited[adjacent[i]]; !alreadyVisited {
			if contains(pageOrder, adjacent[i]) && dfs(pageOrder, adjacent[i], target, immutableAppendToVisited(visited, adjacent[i])) {
				return true
			}
		}
	}

	return false
}

func verifyPageOrder(pageOrder []string, offset int) bool {
	for i := 0; i < offset; i++ {
		if dfs(pageOrder, pageOrder[offset], pageOrder[i], map[string]bool{}) {
			return false
		}
	}

	return true
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

			pageOrderCorrect := true

			// Assume all pages in each page order list exist in the built graph
			for i := len(pageOrder) - 1; i >= 0; i-- {
				if !verifyPageOrder(pageOrder, i) {
					pageOrderCorrect = false
					break
				}
			}

			if pageOrderCorrect {
				if middlePage, convErr := strconv.Atoi(pageOrder[len(pageOrder)/2]); convErr == nil {
					middlePageSum += middlePage
				} else {
					fmt.Println("Unexpected error with page number formatting: " + convErr.Error())
				}
			}
		}
	}

	fmt.Println(middlePageSum)
}
