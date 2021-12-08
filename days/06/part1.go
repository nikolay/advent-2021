package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const days = 80

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var line string
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			break
		}
	}

	parts := strings.Split(line, ",")
	fish := make(map[int]int)
	for _, p := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			log.Fatal(err)
		}
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
