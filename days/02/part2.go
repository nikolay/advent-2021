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

	scanner := bufio.NewScanner(file)
	x, aim, depth := 0, 0, 0
	for scanner.Scan() {
		cmd := scanner.Text()
		matches := r.FindStringSubmatch(cmd)
		if len(matches) == 0 {
			log.Fatal(errors.New(fmt.Sprintf("invalid command: %v", cmd)))
		}
		parameter, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatal(err)
		}
		switch verb := matches[1]; verb {
		case "forward":
			x += parameter
			depth += aim * parameter
		case "up":
			aim -= parameter
		case "down":
			aim += parameter
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(x * depth)
}
