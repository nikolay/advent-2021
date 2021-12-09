package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calcFuel(positions []int, position int) (fuel int) {
	fuel = 0
	for _, p := range positions {
		offset := p - position
		if offset < 0 {
			fuel -= offset
		} else {
			fuel += offset
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
	min, max := 0, 0
	for i, p := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			log.Fatal(err)
		}
		if i == 0 {
			min, max = num, num
		} else if num < min {
			min = num
		} else if num > max {
			max = num
		}
		positions = append(positions, num)
	}

	minFuel := -1
	for i := min; i <= max; i++ {
		fuel := calcFuel(positions, i)
		if minFuel == -1 || fuel < minFuel {
			minFuel = fuel
		}
	}
	fmt.Println(minFuel)
}
