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
	x, y, z int64
}

type Cuboid struct {
	state bool
	a, b  Point
}

type Space struct {
	cuboids []Cuboid
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func abs(x int64) int64 {
	if x > 0 {
		return x
	}
	return -x
}

func (this Cuboid) Volume() int64 {
	return (abs(this.b.x-this.a.x) + 1) * (abs(this.b.y-this.a.y) + 1) * (abs(this.b.z-this.a.z) + 1)
}

func (this Cuboid) IsValid() bool {
	return this.a.x <= this.b.x && this.a.y <= this.b.y && this.a.z <= this.b.z
}

func (this Cuboid) Intersect(that Cuboid, state bool) Cuboid {
	return Cuboid{
		state,
		Point{max(this.a.x, that.a.x), max(this.a.y, that.a.y), max(this.a.z, that.a.z)},
		Point{min(this.b.x, that.b.x), min(this.b.y, that.b.y), min(this.b.z, that.b.z)},
	}
}

func (this *Space) Volume() (volume int64) {
	list := make([]Cuboid, 0)
	for _, c1 := range this.cuboids {
		addons := make([]Cuboid, 0)
		if c1.state {
			addons = append(addons, c1)
		}
		for _, c2 := range list {
			if intersection := c1.Intersect(c2, !c2.state); intersection.IsValid() {
				addons = append(addons, intersection)
			}
		}
		list = append(list, addons...)
	}
	for _, c := range list {
		if c.state {
			volume += c.Volume()
		} else {
			volume -= c.Volume()
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

	r := regexp.MustCompile(`^(on|off) x=([-]?\d+)..([-]?\d+),y=([-]?\d+)..([-]?\d+),z=([-]?\d+)..([-]?\d+)$`)

	space := Space{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		matches := r.FindStringSubmatch(line)
		if len(matches) > 0 {
			state := false
			if matches[1] == "on" {
				state = true
			}

			x1, _ := strconv.ParseInt(matches[2], 10, 64)
			x2, _ := strconv.ParseInt(matches[3], 10, 64)

			y1, _ := strconv.ParseInt(matches[4], 10, 64)
			y2, _ := strconv.ParseInt(matches[5], 10, 64)

			z1, _ := strconv.ParseInt(matches[6], 10, 64)
			z2, _ := strconv.ParseInt(matches[7], 10, 64)

			x1, x2 = min(x1, x2), max(x1, x2)
			y1, y2 = min(y1, y2), max(y1, y2)
			z1, z2 = min(z1, z2), max(z1, z2)

			space.cuboids = append(space.cuboids, Cuboid{state, Point{x1, y1, z1}, Point{x2, y2, z2}})
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(space.Volume())
}
