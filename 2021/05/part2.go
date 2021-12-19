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
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Point struct {
	x, y int
}

type Line struct {
	a, b Point
}

func (p Point) Equal(o Point) bool {
	return p.x == o.x && p.y == o.y
}

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

func cmp(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func (l Line) Rasterize() (result []Point) {
	x1, y1 := l.a.x, l.a.y
	x2, y2 := l.b.x, l.b.y
	dx, dy := cmp(x2, x1), cmp(y2, y1)
	for x1 != x2 || y1 != y2 {
		result = append(result, Point{x1, y1})
		x1 += dx
		y1 += dy
	}
	return append(result, Point{x1, y1})
}

func (l Line) Cross(m Line) (result []Point) {
	points := append(l.Rasterize(), m.Rasterize()...)
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			if points[i].Equal(points[j]) {
				result = append(result, points[i])
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

	r := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)

	lines := make([]Line, 0, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cmd := scanner.Text()
		matches := r.FindStringSubmatch(cmd)
		if len(matches) == 0 {
			log.Fatal(errors.New(fmt.Sprintf("invalid command: %v", cmd)))
		}
		x1, _ := strconv.Atoi(matches[1])
		y1, _ := strconv.Atoi(matches[2])
		x2, _ := strconv.Atoi(matches[3])
		y2, _ := strconv.Atoi(matches[4])
		lines = append(lines, Line{Point{x1, y1}, Point{x2, y2}})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	points := make([]Point, 0, 0)
	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			cross := lines[i].Cross(lines[j])
			if len(cross) > 0 {
				points = append(points, cross...)
			}
		}
	}
	sort.Slice(points, func(a, b int) bool {
		return points[a].x < points[b].x || points[a].x == points[b].x && points[a].y < points[b].y
	})

	count := 1
	for i := 1; i < len(points); i++ {
		if !points[i].Equal(points[i-1]) {
			count++
		}
	}
	fmt.Println(count)
}
