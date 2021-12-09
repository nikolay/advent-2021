package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func calcRisk(prev, curr, next []rune) (result int) {
	for i, c := range curr {
		up := prev == nil || prev[i] > c
		left := i == 0 || curr[i-1] > c
		right := i == len(curr)-1 || curr[i+1] > c
		down := next == nil || next[i] > c
		if up && left && right && down {
			result += int(c-'0') + 1
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

	var buffer []string = make([]string, 0, 3)
	risk := 0
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
