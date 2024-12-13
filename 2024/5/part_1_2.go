package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const NoIncorrectlyOrderedPages = -1

var graph map[string][]string = map[string][]string{}

func contains(list []string, target string) bool {
	for _, el := range list {
		if el == target {
			return true
		}
	}

	return false
}

func addToGraph(line string) {
	pageNums := strings.Split(line, "|")

	if _, fnd := graph[pageNums[0]]; !fnd {
		graph[pageNums[0]] = []string{pageNums[1]}
	} else {
		graph[pageNums[0]] = append(graph[pageNums[0]], pageNums[1])
	}
}

func dfs(pageOrder []string, currGraphNode, target string, visited map[string]bool) bool {
	if currGraphNode == target {
		return true
	}

	adjacent := graph[currGraphNode]

	for i := 0; i < len(adjacent); i++ {
		if _, alreadyVisited := visited[adjacent[i]]; !alreadyVisited {
			visited[adjacent[i]] = true

			// Don't consider ordering rules containing nonexistent pages
			if contains(pageOrder, adjacent[i]) && dfs(pageOrder, adjacent[i], target, visited) {
				return true
			}
		}
	}

	return false
}

// Starting from the page at index _offset_ in the page order, if any of the pages that come before it in the order
// (0..<offset) can be reached, these pages are not in the right order.

// Return the index of the page that should come before the page at index _offset_ but instead comes after
// (for part 2), or -1 if this page is in the right order.
func verifyPageOrder(pageOrder []string, offset int) bool {
	for i := 0; i < offset; i++ {
		if dfs(pageOrder, pageOrder[offset], pageOrder[i], map[string]bool{}) {
			return false
		}
	}

	return true
}

func parseMiddlePage(pageOrder []string) int {
	if middlePage, convErr := strconv.Atoi(pageOrder[len(pageOrder)/2]); convErr == nil {
		return middlePage
	} else {
		panic("Unexpected error with page number formatting: " + convErr.Error())
	}
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic("Error opening file: " + err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1MiddlePageSum := 0
	part2MiddlePageSum := 0

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
				part1MiddlePageSum += parseMiddlePage(pageOrder)
			} else {
				// Topologically sort the page order based on each page's relative position in the graph.
				// An alternative could be to swap indices in violation after calling verifyPageOrder(),
				// but this gets the job done.
				sort.Slice(pageOrder, func(i, j int) bool {
					return dfs(pageOrder, pageOrder[i], pageOrder[j], map[string]bool{})
				})

				part2MiddlePageSum += parseMiddlePage(pageOrder)
			}
		}
	}

	fmt.Println(part1MiddlePageSum)
	fmt.Println(part2MiddlePageSum)
}
