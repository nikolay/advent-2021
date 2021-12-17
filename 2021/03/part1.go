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

	count := uint(0)
	freq := make([]uint, 0, 64)
	bitsize := int(0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bits := scanner.Text()
		if bitsize == 0 {
			bitsize = len(bits)
			freq = freq[0:bitsize]
		}
		decimal, _ := strconv.ParseUint(bits, 2, bitsize)
		for i := 0; i < bitsize; i++ {
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
	for i := 0; i < bitsize; i++ {
		if freq[i]<<1 > count {
			gamma |= 1 << i
		} else {
			epsilon |= 1 << i
		}
	}
	fmt.Println(gamma * epsilon)
}
