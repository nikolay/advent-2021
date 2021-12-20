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

const (
	enhancementsPart1 = 2
	enhancementsPart2 = 50
)

const (
	lightPixel = '#'
	darkPixel  = '.'
)

type Point struct {
	x, y int
}

type (
	Algo  [512]bool
	Image map[Point]bool
)

func getBounds(pixels Image) (minX, maxX, minY, maxY int) {
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

func bitmap(pixels Image, x0, y0 int, minX, maxX, minY, maxY int, void bool) (result uint) {
	var pixel bool
	for dy := -1; dy <= 1; dy++ {
		y := y0 + dy
		for dx := -1; dx <= 1; dx++ {
			x := x0 + dx
			if x < minX || x > maxX || y < minY || y > maxY {
				pixel = void
			} else {
				if v, ok := pixels[Point{x, y}]; ok {
					pixel = v
				} else {
					pixel = false
				}
			}
			result <<= 1
			if pixel {
				result |= 0b000000001
			}
		}
	}
	return
}

func enhance(pixels Image, algo Algo, void bool) (newPixels Image, result int) {
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
			algo[i] = r == lightPixel
		}
	}
	pixels := make(Image)
	y := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		for x, r := range line {
			if r == lightPixel {
				pixels[Point{x, y}] = true
			}
		}
		y++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var result int
	void := false
	for i := 0; i < enhancementsPart2; i++ {
		switch i {
		case enhancementsPart1:
			fmt.Println("Part 1", "=", result)
		}
		pixels, result = enhance(pixels, algo, void)
		if void {
			void = algo[0b111111111]
		} else {
			void = algo[0b000000000]
		}
	}
	fmt.Println("Part 2", "=", result)
}
