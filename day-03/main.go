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
	count := uint(0)
	freq := make(map[uint]uint, 5)
	for scanner.Scan() {
		bits, err := strconv.ParseUint(scanner.Text(), 2, 5)
		if err != nil {
			log.Fatal(err)
		}
		count++
		if bits&1 != 0 {
			freq[0]++
		}
		if bits&2 != 0 {
			freq[1]++
		}
		if bits&4 != 0 {
			freq[2]++
		}
		if bits&8 != 0 {
			freq[3]++
		}
		if bits&16 != 0 {
			freq[4]++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	gamma := 0
	if freq[0]<<1 > count {
		gamma |= 1
	}
	if freq[1]<<1 > count {
		gamma |= 2
	}
	if freq[2]<<1 > count {
		gamma |= 4
	}
	if freq[3]<<1 > count {
		gamma |= 8
	}
	if freq[4]<<1 > count {
		gamma |= 16
	}
	epsilon := gamma ^ 0x1f
	fmt.Println(gamma * epsilon)
}
