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
	"regexp"
	"strconv"
	"strings"
)

const winningScore = 1000

var lastRoll = 1

func roll() (result int) {
	result = lastRoll
	lastRoll = lastRoll%100 + 1
	return
}

func move(position, roll int) int {
	return (position+roll-1)%10 + 1
}

func solve(position1, position2 int) int {
	score1, score2 := 0, 0
	rolls := 0
	for score1 < 1000 && score2 < 1000 {
		roll1 := roll() + roll() + roll()
		rolls += 3
		position1 = move(position1, roll1)
		score1 += position1

		if score1 >= winningScore {
			break
		}

		roll2 := roll() + roll() + roll()
		rolls += 3
		position2 = move(position2, roll2)
		score2 += position2
	}
	if score1 >= winningScore {
		return rolls * score2
	}
	return rolls * score1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	playerRegexp := regexp.MustCompile(`^Player (\d+) starting position: (\d+)$`)
	var players [2]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		matches := playerRegexp.FindStringSubmatch(line)
		if len(matches) > 0 {
			number, _ := strconv.Atoi(matches[1])
			position, _ := strconv.Atoi(matches[2])
			players[number-1] = position
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(players[0], players[1]))
}
