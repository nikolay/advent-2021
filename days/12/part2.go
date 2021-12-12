package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func isLower(s string) bool {
	return strings.ToLower(s) == s
}

func isUpper(s string) bool {
	return strings.ToUpper(s) == s
}

func canBeVisited(arr []string, s string) bool {
	if s == "start" || s == "end" {
		return false
	}
	if isUpper(s) {
		return true
	}
	counts := map[string]int{}
	maxCount := 0
	for _, v := range arr {
		if isLower(v) {
			counts[v]++
			if counts[v] > maxCount {
				maxCount = counts[v]
			}
		}
	}
	if maxCount == 2 {
		if counts[s] == 0 {
			return true
		}
		return false
	}
	return true
}

func findPaths(caves map[string][]string, queue []string, current string, end string) (result [][]string) {
	for _, next := range caves[current] {
		if next == end {
			path := append(queue, next)
			//			fmt.Println(strings.Join(path, ","))
			result = append(result, path)
		} else if isUpper(next) || next != "start" && next != "end" && isLower(next) && canBeVisited(queue, next) {
			result = append(result, findPaths(caves, append(queue, next), next, end)...)
		}
	}
	return
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	caves := map[string][]string{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		vertices := strings.Split(line, "-")
		caves[vertices[0]] = append(caves[vertices[0]], vertices[1])
		caves[vertices[1]] = append(caves[vertices[1]], vertices[0])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(findPaths(caves, []string{"start"}, "start", "end")))
}
