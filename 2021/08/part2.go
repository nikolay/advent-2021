/*
 *   Copyright (c) 2021
 *   All rights reserved.
 */
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func match(digit string, patttern string) bool {
	for _, rune := range patttern {
		if !strings.ContainsRune(digit, rune) {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, "|")

		patterns := strings.Fields(parts[0])
		digits := strings.Fields(parts[1])

		pattern1, pattern4 := "", ""
		for _, pattern := range append(patterns, digits...) {
			switch len(pattern) {
			case 2:
				pattern1 = pattern
			case 4:
				pattern4 = pattern
			}
		}
		if pattern1 == "" || pattern4 == "" {
			log.Fatal(errors.New(fmt.Sprintf("no ones and fours in line: %v", line)))
		}
		patternBD := strings.ReplaceAll(pattern4, pattern1[0:1], "")
		patternBD = strings.ReplaceAll(patternBD, pattern1[1:2], "")

		num := 0
		for _, digit := range digits {
			d := -1
			switch len(digit) {
			case 2:
				d = 1
			case 3:
				d = 7
			case 4:
				d = 4
			case 7:
				d = 8
			case 5:
				if match(digit, pattern1) {
					d = 3
				} else if match(digit, patternBD) {
					d = 5
				} else {
					d = 2
				}
			case 6:
				if !match(digit, patternBD) {
					d = 0
				} else if match(digit, pattern1) {
					d = 9
				} else {
					d = 6
				}
			}
			if d < 0 {
				log.Fatal(errors.New(fmt.Sprintf("cannot determine digit: %v", digit)))
			}
			num = num*10 + d
		}
		sum += num
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
