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
	"strings"
)

type Algo [512]bool

type Point struct {
	x, y int
}

func getBounds(pixels map[Point]bool) (minX, maxX, minY, maxY int) {
	first := true
	for pt, flag := range pixels {
		if flag {
			if first {
				minX, maxX = pt.x, pt.x
				minY, maxY = pt.y, pt.y
				first = false
			} else {
				if pt.x < minX {
					minX = pt.x
				} else if pt.x > maxX {
					maxX = pt.x
				}
				if pt.y < minY {
					minY = pt.y
				} else if pt.y > maxY {
					maxY = pt.y
				}
			}
		}
	}
	return
}

func bitmap(pixels map[Point]bool, x0, y0 int, minX, maxX, minY, maxY int, void bool) (result uint) {
	for dy := -1; dy <= 1; dy++ {
		y := y0 + dy
		for dx := -1; dx <= 1; dx++ {
			x := x0 + dx
			pixel := false
			if x < minX || x > maxX || y < minY || y > maxY {
				pixel = void
			} else {
				if v, ok := pixels[Point{x, y}]; ok {
					pixel = v
				}
			}
			result <<= 1
			if pixel {
				result |= 1
			}
		}
	}
	return
}

func enhance(pixels map[Point]bool, algo Algo, void bool) (newPixels map[Point]bool, result int) {
	minX, maxX, minY, maxY := getBounds(pixels)
	newPixels = make(map[Point]bool)
	for y := minY - 1; y <= maxY+1; y++ {
		for x := minX - 1; x <= maxX+1; x++ {
			if algo[bitmap(pixels, x, y, minX, maxX, minY, maxY, void)] {
				newPixels[Point{x, y}] = true
				result++
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

	var algo Algo
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		for i, r := range line {
			algo[i] = r == '#'
		}
	}
	pixels := make(map[Point]bool)
	y := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		for x, r := range line {
			if r == '#' {
				pixels[Point{x, y}] = true
			}
		}
		y++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var result int
	for i := 0; i < 50; i++ {
		switch i {
		case 2:
			fmt.Println("Part 1", "=", result)
		}
		var void bool
		if i%2 == 1 {
			void = algo[0]
		} else {
			void = false
		}
		pixels, result = enhance(pixels, algo, void)
	}
	fmt.Println("Part 2", "=", result)
}
