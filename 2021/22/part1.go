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

type Cuboid struct {
	state bool
	a, b  Point
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

func cubes(cuboids []Cuboid) (count int) {
	for x := -50; x <= 50; x++ {
		for y := -50; y <= 50; y++ {
			for z := -50; z <= 50; z++ {
				cube := 0
				for _, cuboid := range cuboids {
					if x >= cuboid.a.x && x <= cuboid.b.x && y >= cuboid.a.y && y <= cuboid.b.y && z >= cuboid.a.z && z <= cuboid.b.z {
						if cuboid.state {
							cube = 1
						} else {
							cube = 0
						}
					}
				}
				count += cube
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

	cuboidRegexp := regexp.MustCompile(`^(on|off) x=([-]?\d+)..([-]?\d+),y=([-]?\d+)..([-]?\d+),z=([-]?\d+)..([-]?\d+)$`)

	cuboids := make([]Cuboid, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		matches := cuboidRegexp.FindStringSubmatch(line)
		if len(matches) > 0 {
			state := false
			if matches[1] == "on" {
				state = true
			}

			x1, _ := strconv.Atoi(matches[2])
			x2, _ := strconv.Atoi(matches[3])

			y1, _ := strconv.Atoi(matches[4])
			y2, _ := strconv.Atoi(matches[5])

			z1, _ := strconv.Atoi(matches[6])
			z2, _ := strconv.Atoi(matches[7])

			x1, x2 = min(x1, x2), max(x1, x2)
			y1, y2 = min(y1, y2), max(y1, y2)
			z1, z2 = min(z1, z2), max(z1, z2)

			if x1 < -50 || y1 < -50 || z1 < -50 || x2 > 50 || y2 > 50 || z2 > 50 {
				continue
			}

			cuboids = append(cuboids, Cuboid{state, Point{x1, y1, z1}, Point{x2, y2, z2}})
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(cubes(cuboids))
}
