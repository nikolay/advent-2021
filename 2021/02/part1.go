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
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := regexp.MustCompile(`^(forward|up|down) (\d+)$`)

	x, depth := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cmd := scanner.Text()
		matches := r.FindStringSubmatch(cmd)
		if len(matches) == 0 {
			log.Fatal(errors.New(fmt.Sprintf("invalid command: %v", cmd)))
		}
		parameter, _ := strconv.Atoi(matches[2])
		switch matches[1] {
		case "forward":
			x += parameter
		case "up":
			depth -= parameter
		case "down":
			depth += parameter
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(x * depth)
}
