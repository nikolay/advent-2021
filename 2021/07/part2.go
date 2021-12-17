package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
		num, _ := strconv.Atoi(strings.TrimSpace(p))
		positions = append(positions, num)
	}

	minFuel := math.MinInt
	for _, p := range positions {
		fuel := calcFuel(positions, p)
		if minFuel == math.MinInt || fuel < minFuel {
			minFuel = fuel
		}
	}
	fmt.Println(minFuel)
}
