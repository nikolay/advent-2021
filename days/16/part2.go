package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func hexToBin(s string) (result string) {
	tans := map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}
	for _, r := range s {
		result += tans[r]
	}
	return
}

type Literal struct {
	literal int64
}

type Operator struct {
	typeID   int64
	operands []*Packet
}

type Packet struct {
	version  int64
	literal  *Literal
	operator *Operator
}

func parseLiteral(s string, pos int) (result Literal, i int) {
	i = pos
	n := ""
	for {
		bits := s[i : i+5]
		i += 5
		n += bits[1:]
		if bits[0] == '0' {
			break
		}
	}
	result.literal, _ = strconv.ParseInt(n, 2, 64)
	return
}

func parseOperator(s string, pos int, typeID int64) (result Operator, i int) {
	i = pos
	result.typeID = typeID
	lengthTypeID := s[i : i+1]
	i++
	if lengthTypeID == "0" {
		totalLength, _ := strconv.ParseInt(s[i:i+15], 2, 64)
		i += 15
		endPos := i + int(totalLength)
		result.operands = make([]*Packet, 0)
		for i < endPos {
			var packet Packet
			packet, i = parsePacket(s, i)
			result.operands = append(result.operands, &packet)
		}
	} else if lengthTypeID == "1" {
		numberOfSubpackets, _ := strconv.ParseInt(s[i:i+11], 2, 64)
		i += 11
		result.operands = make([]*Packet, numberOfSubpackets)
		for j := int64(0); j < numberOfSubpackets; j++ {
			var subpacket Packet
			subpacket, i = parsePacket(s, i)
			result.operands[j] = &subpacket
		}
	}
	return
}

func parsePacket(s string, pos int) (result Packet, i int) {
	i = pos
	result.version, _ = strconv.ParseInt(s[i:i+3], 2, 64)
	i += 3
	typeID, _ := strconv.ParseInt(s[i:i+3], 2, 64)
	i += 3
	if typeID == 4 {
		var literal Literal
		result.literal = &literal
		literal, i = parseLiteral(s, i)
	} else {
		var operator Operator
		result.operator = &operator
		operator, i = parseOperator(s, i, typeID)
	}
	return
}

func calc(p *Packet) (result int64) {
	if p == nil {
		return
	}
	if p.literal != nil {
		result = p.literal.literal
	} else if p.operator != nil {
		for i, op := range p.operator.operands {
			value := calc(op)
			switch p.operator.typeID {
			case 0:
				result += value
			case 1:
				if i == 0 {
					result = value
				} else {
					result *= value
				}
			case 2:
				if i == 0 {
					result = value
				} else if value < result {
					result = value
				}
			case 3:
				if i == 0 {
					result = value
				} else if value > result {
					result = value
				}
			case 5:
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
			case 6:
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
			case 7:
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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := int64(0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		trans := hexToBin(line)
		packet, _ := parsePacket(trans, 0)
		s := calc(&packet)
		sum += s
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	println(sum)
}
