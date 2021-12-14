package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func polymer(formula string, rules map[string]string, steps int) int64 {
	freq := make(map[string]int64)
	for i := 0; i < len(formula)-1; i++ {
		freq[formula[i:i+2]]++
	}
	for steps > 0 {
		newFreq := make(map[string]int64)
		for k, v := range rules {
			if f, ok := freq[k]; ok {
				newFreq[k[0:1]+v] += f
				newFreq[v+k[1:2]] += f
			}
		}
		freq = newFreq
		steps--
	}
	elementFreq := make(map[byte]int64)
	for k, v := range freq {
		elementFreq[k[0]] += v
	}
	most, least := int64(-1), int64(-1)
	for _, v := range elementFreq {
		if least == -1 || v < least {
			least = v
		}
		if most == -1 || v > most {
			most = v
		}
	}
	return most - least + 1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	formula := ""
	rules := make(map[string]string)
	if scanner.Scan() {
		formula = strings.TrimSpace(scanner.Text())
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if len(line) == 0 {
				continue
			}
			parts := strings.Split(line, " -> ")
			rules[parts[0]] = parts[1]
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1 Result:", polymer(formula, rules, 10))
	fmt.Println("Part 2 Result:", polymer(formula, rules, 40))
}
