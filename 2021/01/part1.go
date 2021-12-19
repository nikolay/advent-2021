/*
 *   Copyright (c) 2021
 *   All rights reserved.
 */
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	last := math.MinInt
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		if last != math.MinInt && num > last {
			count++
		}
		last = num
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
