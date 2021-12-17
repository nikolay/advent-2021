package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const bufferSize = 3

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	last := math.MinInt
	count := 0
	buffer := make([]int, 0, bufferSize)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		buffer = append(buffer, num)
		if len(buffer) < bufferSize {
			continue
		}
		sum := 0
		for _, v := range buffer {
			sum += v
		}
		buffer = buffer[1:]
		if last != math.MinInt && sum > last {
			count++
		}
		last = sum
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
