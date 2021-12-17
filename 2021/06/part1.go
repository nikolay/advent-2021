package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var part int

func init() {
	if len(os.Args) > 0 {
		if p, err := strconv.Atoi(os.Args[1]); err != nil {
			log.Fatal(err)
		} else if p < 1 || p > 2 {
			log.Fatal(errors.New(fmt.Sprintf("invalid part: %v", p)))
		} else {
			part = p
		}
	}
}

func main() {
	days := [2]int{80, 256}[part-1]

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(line, ",")
	fish := make(map[int]int)
	for _, p := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(p))
		fish[num] += 1
	}

	for day := 0; day < days; day++ {
		moms := fish[0]
		for i := 0; i < 8; i++ {
			fish[i] = fish[i+1]
		}
		fish[6] += moms
		fish[8] = moms
	}

	count := 0
	for i := 0; i <= 8; i++ {
		count += fish[i]
	}
	fmt.Println(count)
}
