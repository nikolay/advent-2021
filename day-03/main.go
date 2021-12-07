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
	freq := make([]uint, 0, 64)
	bitSize := int(0)
	for scanner.Scan() {
		bits := scanner.Text()
		if bitSize == 0 {
			bitSize = len(bits)
			freq = freq[0:bitSize]
		}
		decimal, err := strconv.ParseUint(bits, 2, bitSize)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < bitSize; i++ {
			if decimal&(1<<i) != 0 {
				freq[i]++
			}
		}
		count++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	gamma, epsilon := 0, 0
	for i := 0; i < bitSize; i++ {
		if freq[i]<<1 > count {
			gamma |= 1 << i
		} else {
			epsilon |= 1 << i
		}
	}
	fmt.Println(gamma * epsilon)
}
