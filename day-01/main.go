package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	last := -1
	count := 0
	for scanner.Scan() {
		scan, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if last != -1 && scan > last {
			count++
		}
		last = scan
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
