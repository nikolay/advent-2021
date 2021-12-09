package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, "|")
		digits := strings.Fields(parts[1])
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
			}
			if d != -1 {
				count++
			}
		}
	}
	fmt.Println(count)
}
