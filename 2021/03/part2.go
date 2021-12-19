/*
 *   Copyright (c) 2021
 *   All rights reserved.
 */
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calc(nums []uint64, bitSize int, reverse bool) uint64 {
	count := len(nums)
	flag := make([]bool, count)
	bitmask := uint64(1) << (bitSize - 1)
	for bitmask > 0 && count > 1 {
		ones, zeroes := 0, 0
		for i, n := range nums {
			if flag[i] {
				continue
			}
			if n&bitmask != 0 {
				ones++
			} else {
				zeroes++
			}
		}
		bit := uint64(0)
		if reverse {
			if ones < zeroes {
				bit = 1
			}
		} else if ones >= zeroes {
			bit = 1
		}
		targetBitmask := bitmask * bit
		for i, n := range nums {
			if flag[i] {
				continue
			}
			if n&bitmask != targetBitmask {
				flag[i] = true
				count--
				if count == 1 {
					break
				}
			}
		}
		bitmask >>= 1
	}
	for i, f := range flag {
		if !f {
			return nums[i]
		}
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nums := make([]uint64, 0, 64)
	bitsize := int(0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bits := scanner.Text()
		if bitsize == 0 {
			bitsize = len(bits)
		}
		decimal, _ := strconv.ParseUint(bits, 2, bitsize)
		nums = append(nums, decimal)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(calc(nums, bitsize, false) * calc(nums, bitsize, true))
}
