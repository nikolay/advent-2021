package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func testLine(line []rune) (result int) {
	score := make(map[rune]int)
	score[')'] = 1
	score[']'] = 2
	score['}'] = 3
	score['>'] = 4
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
			return
		}
		if r != stack[len(stack)-1] {
			return
		}
		stack = stack[:len(stack)-1]
	}
	for i := len(stack) - 1; i >= 0; i-- {
		result = result*5 + score[stack[i]]
	}
	return
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scores := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		score := testLine([]rune(line))
		if score != 0 {
			scores = append(scores, score)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Sort(sort.IntSlice(scores))
	fmt.Println(scores[(len(scores)+1)/2-1])
}
