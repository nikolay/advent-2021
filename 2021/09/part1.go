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

func calcRisk(prev, curr, next []rune) (result int) {
	for i, r := range curr {
		up := prev == nil || prev[i] > r
		left := i == 0 || curr[i-1] > r
		right := i == len(curr)-1 || curr[i+1] > r
		down := next == nil || next[i] > r
		if up && left && right && down {
			result += int(r-'0') + 1
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

	var buffer []string = make([]string, 0, 3)
	risk := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		buffer = append(buffer, line)
		if len(buffer) == 4 {
			buffer = buffer[1:]
		}
		l := len(buffer)
		if l == 1 {
			continue
		}
		var prev []rune
		if l == 3 {
			prev = []rune(buffer[0])
		} else {
			prev = nil
		}
		curr := []rune(buffer[l-2])
		next := []rune(buffer[l-1])
		risk += calcRisk(prev, curr, next)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	l := len(buffer)
	prev := []rune(buffer[l-2])
	curr := []rune(buffer[l-1])
	risk += calcRisk(prev, curr, nil)
	fmt.Println(risk)
}
