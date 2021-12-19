/*
 *   Copyright (c) 2021
 *   All rights reserved.
 */
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

func contains(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

func findPaths(caves map[string][]string, queue []string, current string, end string) (result [][]string) {
	for _, next := range caves[current] {
		if next == end {
			result = append(result, append(queue, next))
		} else if isUpper(next) || isLower(next) && !contains(queue, next) {
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

	caves := map[string][]string{}
	scanner := bufio.NewScanner(file)
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
