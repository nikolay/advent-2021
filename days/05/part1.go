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

func (l Line) IsVertical() bool {
	return l.a.x == l.b.x
}

func (l Line) IsHorizontal() bool {
	return l.a.y == l.b.y
}

func (l *Line) Normalize() *Line {
	if l.IsVertical() && (l.a.y > l.b.y) || l.IsHorizontal() && (l.a.x > l.b.x) {
		return &Line{l.b, l.a}
	} else {
		return l
	}
}

func (l Line) Cross(m Line) (result []Point) {
	l1, m1 := l.Normalize(), m.Normalize()
	if l1.IsVertical() && m1.IsVertical() {
		if l1.a.y > m1.a.y {
			l1, m1 = m1, l1
		}
		if l1.a.x == m1.a.x {
			y1, y2 := max(l1.a.y, m1.a.y), min(l1.b.y, m1.b.y)
			if y1 <= y2 {
				result = make([]Point, y2-y1+1)
				for y := y1; y <= y2; y++ {
					result[y-y1] = Point{l1.a.x, y}
				}
				return
			}
		}
	} else if l1.IsHorizontal() && m1.IsHorizontal() {
		if l1.a.x > m1.a.x {
			l1, m1 = m1, l1
		}
		if l1.a.y == m1.a.y {
			x1, x2 := max(l1.a.x, m1.a.x), min(l1.b.x, m1.b.x)
			if x1 <= x2 {
				result = make([]Point, x2-x1+1)
				for x := x1; x <= x2; x++ {
					result[x-x1] = Point{x, l1.a.y}
				}
				return
			}
		}
	} else {
		if l1.IsVertical() {
			l1, m1 = m1, l1
		}
		if l1.a.x <= m1.a.x && l1.b.x >= m1.a.x && m1.a.y <= l1.a.y && m1.b.y >= l1.b.y {
			return []Point{{m1.a.x, l1.a.y}}
		}
	}
	return nil
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
		line := Line{Point{x1, y1}, Point{x2, y2}}
		if line.IsHorizontal() || line.IsVertical() {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	points := make([]Point, 0, 0)
	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			cross := lines[i].Cross(lines[j])
			if cross != nil && len(cross) > 0 {
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
