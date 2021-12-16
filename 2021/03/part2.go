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
	bitMask := uint64(1) << (bitSize - 1)
	for bitMask > 0 && count > 1 {
		ones, zeroes := 0, 0
		for i, n := range nums {
			if flag[i] {
				continue
			}
			if n&bitMask != 0 {
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
		targetBitMask := bitMask * bit
		for i, n := range nums {
			if flag[i] {
				continue
			}
			if n&bitMask != targetBitMask {
				flag[i] = true
				count--
				if count == 1 {
					break
				}
			}
		}
		bitMask >>= 1
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
	bitSize := int(0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bits := scanner.Text()
		if bitSize == 0 {
			bitSize = len(bits)
		}
		decimal, err := strconv.ParseUint(bits, 2, bitSize)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, decimal)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(calc(nums, bitSize, false) * calc(nums, bitSize, true))
}
