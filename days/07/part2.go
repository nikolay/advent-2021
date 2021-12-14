package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calcFuel(positions []int, position int) (fuel int) {
	fuel = 0
	for _, p := range positions {
		distance := abs(p - position)
		fuel += distance * (distance + 1) / 2
	}
	return
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(line, ",")
	positions := make([]int, 0, len(parts))
	for _, p := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			log.Fatal(err)
		}
		positions = append(positions, num)
	}

	minFuel := -1
	for _, p := range positions {
		fuel := calcFuel(positions, p)
		if minFuel == -1 || fuel < minFuel {
			minFuel = fuel
		}
	}
	fmt.Println(minFuel)
}
