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

const winningScore = 21

type Memory struct {
	positions  [2]int
	score      [2]int
	multiplier uint64
	turn       byte
}

var rolls map[int]uint64

func calculateRollFrequencies() (rolls map[int]uint64) {
	rolls = make(map[int]uint64)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				rolls[i+j+k]++
			}
		}
	}
	return
}

func move(position, roll int) int {
	return (position+roll-1)%10 + 1
}

func multiverse(positions [2]int, scores [2]int, multiplier uint64, turn byte, memory map[Memory][2]uint64) (wins [2]uint64) {
	m := Memory{positions, scores, multiplier, turn}
	if w, ok := memory[m]; ok {
		wins = w
	} else {
		position, score := positions[turn], scores[turn]
		for roll, frequency := range rolls {
			p := move(position, roll)
			s := score + p
			positions[turn], scores[turn] = p, s
			if s >= winningScore {
				wins[turn] += multiplier * frequency
				continue
			}
			w := multiverse(positions, scores, multiplier*frequency, 1-turn, memory)
			wins[0] += w[0]
			wins[1] += w[1]
		}
		memory[m] = wins
	}
	return
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

	rolls = calculateRollFrequencies()
	memory := make(map[Memory][2]uint64)
	w := multiverse(players, [2]int{0, 0}, 1, 0, memory)

	if w[0] > w[1] {
		fmt.Println(w[0])
	} else {
		fmt.Println(w[1])
	}
}
