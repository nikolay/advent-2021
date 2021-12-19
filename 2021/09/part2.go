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
	"sort"
	"strings"
)

type Coord struct {
	row, col int
}

func findLows(matrix [][]int) (result []Coord) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			height := matrix[row][col]
			up := row == 0 || matrix[row-1][col] > height
			left := col == 0 || matrix[row][col-1] > height
			right := col == len(matrix[row])-1 || matrix[row][col+1] > height
			down := row == len(matrix)-1 || matrix[row+1][col] > height
			if up && left && right && down {
				result = append(result, Coord{row, col})
			}
		}
	}
	return
}

func findBasinSize(matrix [][]int, low Coord) (surface int) {
	visited := make([][]bool, len(matrix))
	for row := range visited {
		visited[row] = make([]bool, len(matrix[row]))
	}
	visited[low.row][low.col] = true
	surface++
	any := true
	for any {
		any = false
		for row := 0; row < len(visited); row++ {
			for col := 0; col < len(visited[row]); col++ {
				if visited[row][col] || matrix[row][col] == 9 {
					continue
				}
				up := row > 0 && visited[row-1][col]
				left := col > 0 && visited[row][col-1]
				right := col < len(visited[row])-1 && visited[row][col+1]
				down := row < len(visited)-1 && visited[row+1][col]
				if up || left || right || down {
					visited[row][col] = true
					surface++
					any = true
				}
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

	matrix := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		digits := make([]int, 0, len(line))
		for _, r := range []rune(line) {
			digits = append(digits, int(r-'0'))
		}
		matrix = append(matrix, digits)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	lows := findLows(matrix)
	basins := make([]int, len(lows))
	for i, low := range lows {
		basins[i] = findBasinSize(matrix, low)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	fmt.Println(basins[0] * basins[1] * basins[2])
}
