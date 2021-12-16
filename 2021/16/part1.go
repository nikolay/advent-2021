package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func hexToBin(s string) (result string) {
	tans := map[rune]string{
		'0': "0000", '1': "0001", '2': "0010", '3': "0011",
		'4': "0100", '5': "0101", '6': "0110", '7': "0111",
		'8': "1000", '9': "1001", 'A': "1010", 'B': "1011",
		'C': "1100", 'D': "1101", 'E': "1110", 'F': "1111",
	}
	for _, r := range s {
		result += tans[r]
	}
	return
}

type Literal struct {
	value int64
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

func chomp(s string, pos *int, size int) string {
	i := *pos
	(*pos) += size
	return s[i:*pos]
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
	number, _ := strconv.ParseInt(bits, 2, 64)
	return &Literal{number}
}

func parseOperator(s string, pos *int, typeID int64) *Operator {
	operator := Operator{typeID, nil}
	if chomp(s, pos, 1) == "0" {
		totalLength, _ := strconv.ParseInt(chomp(s, pos, 15), 2, 64)
		endPos := *pos + int(totalLength)
		operator.operands = make([]*Packet, 0)
		for *pos < endPos {
			operator.operands = append(operator.operands, parsePacket(s, pos))
		}
	} else {
		numberOfSubpackets, _ := strconv.ParseInt(chomp(s, pos, 11), 2, 64)
		operator.operands = make([]*Packet, numberOfSubpackets)
		for i := range operator.operands {
			operator.operands[i] = parsePacket(s, pos)
		}
	}
	return &operator
}

func parsePacket(s string, pos *int) *Packet {
	packet := Packet{}
	packet.version, _ = strconv.ParseInt(chomp(s, pos, 3), 2, 64)
	typeID, _ := strconv.ParseInt(chomp(s, pos, 3), 2, 64)
	if typeID == 4 {
		packet.literal = parseLiteral(s, pos)
	} else {
		packet.operator = parseOperator(s, pos, typeID)
	}
	return &packet
}

func sumVersions(p *Packet) (result int64) {
	if p == nil {
		return
	}
	result = p.version
	if p.operator != nil {
		for _, op := range p.operator.operands {
			result += sumVersions(op)
		}
	}
	return
}

func calc(p *Packet) (result int64) {
	if p == nil {
		return
	}
	if p.literal != nil {
		result = p.literal.value
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
	part := 1
	if len(os.Args) > 0 {
		if p, err := strconv.Atoi(os.Args[1]); err != nil {
			log.Fatal(err)
		} else {
			part = p
		}
	}

	file, err := os.Open("sample2.txt")
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
		var result int64
		packet := parsePacket(hexToBin(line), new(int))
		if part == 1 {
			result = sumVersions(packet)
		} else {
			result = calc(packet)
		}
		fmt.Println(result)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
