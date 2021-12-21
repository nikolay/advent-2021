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

type Point struct {
	x, y, z int
}

type Beacons []Point

type PointTransformation func(p Point) Point

var rotations = [24]PointTransformation{
	func(p Point) Point { return Point{p.x, p.y, p.z} },
	func(p Point) Point { return Point{p.y, p.z, p.x} },
	func(p Point) Point { return Point{p.z, p.x, p.y} },
	func(p Point) Point { return Point{-p.x, p.z, p.y} },
	func(p Point) Point { return Point{p.z, p.y, -p.x} },
	func(p Point) Point { return Point{p.y, -p.x, p.z} },
	func(p Point) Point { return Point{p.x, p.z, -p.y} },
	func(p Point) Point { return Point{p.z, -p.y, p.x} },
	func(p Point) Point { return Point{-p.y, p.x, p.z} },
	func(p Point) Point { return Point{p.x, -p.z, p.y} },
	func(p Point) Point { return Point{-p.z, p.y, p.x} },
	func(p Point) Point { return Point{p.y, p.x, -p.z} },
	func(p Point) Point { return Point{-p.x, -p.y, p.z} },
	func(p Point) Point { return Point{-p.y, p.z, -p.x} },
	func(p Point) Point { return Point{p.z, -p.x, -p.y} },
	func(p Point) Point { return Point{-p.x, p.y, -p.z} },
	func(p Point) Point { return Point{p.y, -p.z, -p.x} },
	func(p Point) Point { return Point{-p.z, -p.x, p.y} },
	func(p Point) Point { return Point{p.x, -p.y, -p.z} },
	func(p Point) Point { return Point{-p.y, -p.z, p.x} },
	func(p Point) Point { return Point{-p.z, p.x, -p.y} },
	func(p Point) Point { return Point{-p.x, -p.z, -p.y} },
	func(p Point) Point { return Point{-p.z, -p.y, -p.x} },
	func(p Point) Point { return Point{-p.y, -p.x, -p.z} },
}

func (p1 Point) Equal(p2 Point) bool {
	return p1.x == p2.x && p1.y == p2.y && p1.z == p2.z
}

func (p Point) Rotate(index int) Point {
	return rotations[index](p)
}

func (p1 Point) Add(p2 Point) Point {
	return Point{p1.x + p2.x, p1.y + p2.y, p1.z + p2.z}
}

func (p1 Point) Subtract(p2 Point) Point {
	return Point{p1.x - p2.x, p1.y - p2.y, p1.z - p2.z}
}

func (b Beacons) Rotate(rotation int) (result Beacons) {
	result = make(Beacons, len(b))
	for i, p := range b {
		result[i] = p.Rotate(rotation)
	}
	return
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func (pt1 Point) ManhattanDistance(pt2 Point) int {
	return abs(pt1.x-pt2.x) + abs(pt1.y-pt2.y) + abs(pt1.z-pt2.z)
}

func (base Beacons) Match(candidate Beacons) (bool, Beacons, Point) {
	maxOverlap := 0
	var maxMerger Beacons
	var maxDelta Point
	for rotation := range rotations {
		c := candidate.Rotate(rotation)

		freq := map[Point]int{}
		for _, p1 := range base {
			for _, p2 := range c {
				delta := p1.Subtract(p2)
				freq[delta]++
			}
		}

		maxFreq, delta := 0, Point{}
		for d, f := range freq {
			if f > maxFreq {
				maxFreq, delta = f, d
			}
		}

		if maxFreq >= 12 && maxFreq > maxOverlap {
			maxOverlap = maxFreq
			maxMerger = make(Beacons, len(base))
			copy(maxMerger, base)
			maxDelta = delta
			for _, p := range c {
				p2 := p.Add(delta)
				dupe := false
				for _, p1 := range base {
					if p2.Equal(p1) {
						dupe = true
						break
					}
				}
				if !dupe {
					maxMerger = append(maxMerger, p2)
				}
			}
		}
	}
	if maxOverlap >= 12 {
		return true, maxMerger, maxDelta
	}
	return false, nil, Point{}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scannerRegexp := regexp.MustCompile(`^--- scanner (\d+) ---$`)
	beaconRegexp := regexp.MustCompile(`^([-]?\d+),([-]?\d+),([-]?\d+)$`)

	scanners := make([]Beacons, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		matches := scannerRegexp.FindStringSubmatch(line)
		if len(matches) > 0 {
			_, _ = strconv.Atoi(matches[1])
			scanners = append(scanners, make(Beacons, 0))
		} else {
			matches := beaconRegexp.FindStringSubmatch(line)
			if len(matches) > 0 {
				x, _ := strconv.Atoi(matches[1])
				y, _ := strconv.Atoi(matches[2])
				z, _ := strconv.Atoi(matches[3])
				scanners[len(scanners)] = append(scanners[len(scanners)], Point{x, y, z})
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	locations := []Point{{0, 0, 0}}

loop:
	for {
		for i := 0; i < len(scanners)-1; i++ {
			base := scanners[i]
			for j := i + 1; i < len(scanners); i++ {
				candidate := scanners[j]
				if match, merger, location := base.Match(candidate); match {
					locations = append(locations, location)
					scanners[i] = merger
					scanners = append(scanners[:j], scanners[j+1:]...)
					continue loop
				}
			}
		}
		break
	}

	count := 0
	for _, beacons := range scanners {
		count += len(beacons)
	}
	fmt.Println("Part 1", "=", count)

	maxDistance := 0
	for i := 0; i < len(locations)-1; i++ {
		for j := i + 1; j < len(locations); j++ {
			distance := locations[i].ManhattanDistance(locations[j])
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}
	fmt.Println("Part 2", "=", maxDistance)
}
