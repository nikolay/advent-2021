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

func hexToBin(s string) (result string) {
	trans := map[rune]string{
		'0': "0000", '1': "0001", '2': "0010", '3': "0011",
		'4': "0100", '5': "0101", '6': "0110", '7': "0111",
		'8': "1000", '9': "1001", 'A': "1010", 'B': "1011",
		'C': "1100", 'D': "1101", 'E': "1110", 'F': "1111",
	}
	for _, r := range s {
		result += trans[r]
	}
	return
}

type Literal struct {
	value int
}

type Operator struct {
	typeID   int
	operands []*Packet
}

type Packet struct {
	version  int
	literal  *Literal
	operator *Operator
}

func chomp(s string, pos *int, size int) string {
	i := *pos
	(*pos) += size
	return s[i:*pos]
}

func parseNumber(s string) int {
	number, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return -1
	}
	return int(number)
}

func chompNumber(s string, pos *int, size int) int {
	return parseNumber(chomp(s, pos, size))
}

func parseLiteral(s string, pos *int) *Literal {
	bits := ""
	for {
		flag := chomp(s, pos, 1)
		bits += chomp(s, pos, 4)
		if flag == "0" {
			break
		}
	}
	return &Literal{parseNumber(bits)}
}

func parseOperator(s string, pos *int, typeID int) *Operator {
	operator := Operator{typeID, nil}
	if chomp(s, pos, 1) == "0" {
		totalLength := chompNumber(s, pos, 15)
		endPos := *pos + totalLength
		operator.operands = make([]*Packet, 0)
		for *pos < endPos {
			operator.operands = append(operator.operands, parsePacket(s, pos))
		}
	} else {
		numberOfSubpackets := chompNumber(s, pos, 11)
		operator.operands = make([]*Packet, numberOfSubpackets)
		for i := range operator.operands {
			operator.operands[i] = parsePacket(s, pos)
		}
	}
	return &operator
}

func parsePacket(s string, pos *int) *Packet {
	const LITERAL = 4
	packet := Packet{}
	packet.version = chompNumber(s, pos, 3)
	typeID := chompNumber(s, pos, 3)
	if typeID == LITERAL {
		packet.literal = parseLiteral(s, pos)
	} else {
		packet.operator = parseOperator(s, pos, typeID)
	}
	return &packet
}

func calcVersionSum(p *Packet) (result int) {
	if p == nil {
		return
	}
	result = p.version
	if p.operator != nil {
		for _, op := range p.operator.operands {
			result += calcVersionSum(op)
		}
	}
	return
}

func calc(p *Packet) (result int) {
	const (
		SUM          = 0
		PRODUCT      = 1
		MINIMUM      = 2
		MAXIMUM      = 3
		GREATER_THAN = 5
		LESS_THAN    = 6
		EQUAL        = 7
	)
	if p == nil {
		return
	}
	if p.literal != nil {
		result = p.literal.value
	} else if p.operator != nil {
		for i, operand := range p.operator.operands {
			value := calc(operand)
			switch p.operator.typeID {
			case SUM:
				result += value
			case PRODUCT:
				if i == 0 {
					result = value
				} else {
					result *= value
				}
			case MINIMUM:
				if i == 0 {
					result = value
				} else if value < result {
					result = value
				}
			case MAXIMUM:
				if i == 0 {
					result = value
				} else if value > result {
					result = value
				}
			case GREATER_THAN:
				if i == 0 {
					result = value
				} else {
					if result > value {
						result = 1
					} else {
						result = 0
					}
					break
				}
			case LESS_THAN:
				if i == 0 {
					result = value
				} else {
					if result < value {
						result = 1
					} else {
						result = 0
					}
					break
				}
			case EQUAL:
				if i == 0 {
					result = value
				} else {
					if result == value {
						result = 1
					} else {
						result = 0
					}
					break
				}
			}
		}
	}
	return
}

var part int

func init() {
	if len(os.Args) > 0 {
		if p, err := strconv.Atoi(os.Args[1]); err != nil {
			log.Fatal(err)
		} else if p < 1 || p > 2 {
			log.Fatal(errors.New(fmt.Sprint("invalid part: %v", p)))
		} else {
			part = p
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		var result int
		packet := parsePacket(hexToBin(line), new(int))
		if part == 1 {
			result = calcVersionSum(packet)
		} else {
			result = calc(packet)
		}
		fmt.Println(result)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
