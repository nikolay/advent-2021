package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func check(vx, vy, x1, y1, x2, y2 int) int {
	x, y := 0, 0
	peak := math.MinInt
	for {
		if y > peak {
			peak = y
		}
		if x >= x1 && x <= x2 && y >= y1 && y <= y2 {
			return peak
		}

		x += vx
		y += vy

		if vx < 0 {
			vx++
		} else if vx > 0 {
			vx--
		}
		vy--

		if vy < 0 && y < y1 || vx == 0 && (x < x1 || x > x2) {
			return math.MinInt
		}
	}
}

func solve(x1, y1, x2, y2 int) (highestPeak, hits int) {
	highestPeak = math.MinInt
	startX, endX := min(0, x2), max(0, x2)
	startY, endY := min(-y1, y1), max(y1, -y1)
	for vx := startX; vx <= endX; vx++ {
		for vy := startY; vy <= endY; vy++ {
			peak := check(vx, vy, x1, y1, x2, y2)
			if peak != math.MinInt {
				hits++
			}
			if peak > highestPeak {
				highestPeak = peak
			}
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

	r := regexp.MustCompile(`^target area: x=([-]?\d+)..([-]?\d+), y=([-]?\d+)..([-]?\d+)$`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		matches := r.FindStringSubmatch(line)
		x1, _ := strconv.Atoi(matches[1])
		x2, _ := strconv.Atoi(matches[2])
		y1, _ := strconv.Atoi(matches[3])
		y2, _ := strconv.Atoi(matches[4])
		highestPeak, hits := solve(x1, y1, x2, y2)
		fmt.Println("Part 1", "=", highestPeak)
		fmt.Println("Part 2", "=", hits)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
