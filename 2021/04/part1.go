/*
 *   Copyright (c) 2021
 *   All rights reserved.
 */
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const width, height = 5, 5

type Coord struct {
	row, col int
}

type Card map[int]Coord

type Result struct {
	colHits [width]int
	rowHits [height]int
}

func hasBingo(result Result) bool {
	for row := 0; row < height; row++ {
		if result.rowHits[row] == width {
			return true
		}
	}
	for col := 0; col < width; col++ {
		if result.colHits[col] == height {
			return true
		}
	}
	return false
}

func scanCard(scanner *bufio.Scanner) (card Card, sum int, err error) {
	card, sum, err = make(Card), 0, nil
	for i := 0; i < height; i++ {
		if scanner.Scan() {
			line := scanner.Text()
			parts := strings.Fields(line)
			if len(parts) == 0 {
				continue
			}
			if len(parts) != width {
				return nil, 0, errors.New(fmt.Sprintf("invalid card line: %v", line))
			}
			for j := 0; j < width; j++ {
				num, _ := strconv.Atoi(parts[j])
				card[num] = Coord{i, j}
				sum += num
			}
		} else {
			return nil, 0, scanner.Err()
		}
	}
	return
}

func scanDraws(scanner *bufio.Scanner) (draws []int, err error) {
	draws, err = make([]int, 0), nil
	var line string
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			break
		}
	}
	parts := strings.Split(line, ",")
	for _, p := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(p))
		draws = append(draws, num)
	}
	return
}

var part int

func init() {
	if len(os.Args) > 0 {
		if p, err := strconv.Atoi(os.Args[1]); err != nil {
			log.Fatal(err)
		} else if p < 1 || p > 2 {
			log.Fatal(errors.New(fmt.Sprintf("invalid part: %v", p)))
		} else {
			part = p
		}
	}
}

func main() {
	firstBingo := part == 1

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	draws, err := scanDraws(scanner)
	if err != nil {
		log.Fatal(err)
	}

	cards := make([]Card, 0, 0)
	scores := make([]int, 0, 0)
	for scanner.Scan() {
		card, score, err := scanCard(scanner)
		if err != nil {
			log.Fatal(err)
		}
		cards = append(cards, card)
		scores = append(scores, score)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	score := math.MinInt
	results := make([]Result, len(cards))
	bingo := make(map[int]bool)
	for _, draw := range draws {
		for j, card := range cards {
			if bingo[j] {
				continue
			}
			coord, hit := card[draw]
			if hit {
				scores[j] -= draw
				results[j].rowHits[coord.row]++
				results[j].colHits[coord.col]++
				if hasBingo(results[j]) {
					bingo[j] = true
					if firstBingo || len(bingo) == len(cards) {
						score = scores[j] * draw
					}
				}
			}
		}
		if score != math.MinInt {
			break
		}
	}
	fmt.Println(score)
}
