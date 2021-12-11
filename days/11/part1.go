package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func calcFlashes(matrix [10][10]int, loops int) (result int) {
	flashed := [10][10]bool{}
	for loop := 0; loop < loops; loop++ {
		for row := 0; row < 10; row++ {
			for col := 0; col < 10; col++ {
				matrix[row][col]++
			}
		}
		any := true
		for any {
			any = false
			for row := 0; row < 10; row++ {
				for col := 0; col < 10; col++ {
					if matrix[row][col] < 10 || flashed[row][col] {
						continue
					}

					flashed[row][col] = true

					if row > 0 {
						matrix[row-1][col]++
					}
					if row < 9 {
						matrix[row+1][col]++
					}
					if col > 0 {
						matrix[row][col-1]++
					}
					if col < 9 {
						matrix[row][col+1]++
					}

					if row > 0 && col > 0 {
						matrix[row-1][col-1]++
					}
					if row > 0 && col < 9 {
						matrix[row-1][col+1]++
					}
					if row < 9 && col > 0 {
						matrix[row+1][col-1]++
					}
					if row < 9 && col < 9 {
						matrix[row+1][col+1]++
					}

					any = true
				}
			}
		}

		for row := 0; row < 10; row++ {
			for col := 0; col < 10; col++ {
				if flashed[row][col] {
					matrix[row][col] = 0
					result++
					flashed[row][col] = false
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

	scanner := bufio.NewScanner(file)

	var matrix [10][10]int
	row := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		for col, b := range line {
			matrix[row][col] = int(b - '0')
		}
		row++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	flashes := calcFlashes(matrix, 100)
	fmt.Println(flashes)
}
