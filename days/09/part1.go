package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(prev, curr, next []rune) (result int) {
	for i, c := range curr {
		checkUp := prev == nil || prev[i] > c
		checkLeft := i == 0 || curr[i-1] > c
		checkRight := i == len(curr)-1 || curr[i+1] > c
		checkDown := next == nil || next[i] > c
		if checkUp && checkLeft && checkRight && checkDown {
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

	var line string
	var buffer []string = make([]string, 0, 3)
	risk := 0
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
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
		risk += check(prev, curr, next)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	l := len(buffer)
	prev := []rune(buffer[l-2])
	curr := []rune(buffer[l-1])
	risk += check(prev, curr, nil)
	fmt.Println(risk)
}
