package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func testLine(line []rune) int {
	score := make(map[rune]int)
	score[')'] = 3
	score[']'] = 57
	score['}'] = 1197
	score['>'] = 25137
	stack := make([]rune, 0)
	for _, r := range line {
		switch r {
		case '(':
			stack = append(stack, ')')
			continue
		case '[':
			stack = append(stack, ']')
			continue
		case '{':
			stack = append(stack, '}')
			continue
		case '<':
			stack = append(stack, '>')
			continue
		}
		if len(stack) == 0 {
			return score[r]
		}
		if r == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			continue
		}
		return score[r]
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		score += testLine([]rune(line))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(score)
}
