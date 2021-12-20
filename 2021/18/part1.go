/*
 *   Copyright (c) 2021
 *   All rights reserved.
 */
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	owner, tree    *Pair
	value          int
	next, previous *Item
}

type Pair struct {
	parent      *Pair
	left, right *Item
}

func (item *Item) IsNumber() bool {
	return !item.IsPair()
}

func (item *Item) IsPair() bool {
	return item.tree != nil
}

func (item *Item) Split() bool {
	if item.value < 10 {
		return false
	}
	previous, next := item.previous, item.next
	left := &Item{
		value:    item.value / 2,
		previous: previous,
	}
	right := &Item{
		value:    (item.value + 1) / 2,
		previous: left,
		next:     next,
	}
	left.next = right

	if previous != nil {
		previous.next = left
	}
	if next != nil {
		next.previous = right
	}

	pair := &Pair{
		parent: item.owner,
		left:   left,
		right:  right,
	}
	left.owner = pair
	right.owner = pair

	item.tree = pair
	item.value = 0

	return true
}

func (pair *Pair) Explode() {
	previous, next := pair.left.previous, pair.right.next
	item := &Item{
		owner:    pair.parent,
		value:    0,
		previous: previous,
		next:     next,
	}
	if previous != nil {
		previous.value += pair.left.value
		previous.next = item
	}
	if next != nil {
		next.value += pair.right.value
		next.previous = item
	}

	if pair == pair.parent.left.tree {
		pair.parent.left = item
	} else if pair == pair.parent.right.tree {
		pair.parent.right = item
	} else {
		log.Fatal(errors.New("pair not in parent"))
	}
}

func (pair *Pair) FindLeaves(leftmost **Item, rightmost **Item) {
	if pair.left.tree != nil {
		pair.left.tree.FindLeaves(leftmost, rightmost)
	} else {
		if leftmost != nil && *leftmost == nil {
			*leftmost = pair.left
		}
	}
	if pair.right.tree != nil {
		pair.right.tree.FindLeaves(leftmost, rightmost)
	} else {
		if rightmost != nil {
			*rightmost = pair.right
		}
	}
}

func (pair *Pair) Append(appendix *Pair) *Pair {
	var leftmost, rightmost *Item = nil, nil
	pair.FindLeaves(nil, &rightmost)
	appendix.FindLeaves(&leftmost, nil)

	rightmost.next = leftmost
	leftmost.previous = rightmost

	newPair := &Pair{
		left:  &Item{tree: pair},
		right: &Item{tree: appendix},
	}

	pair.parent = newPair
	appendix.parent = newPair

	newPair.left.owner = newPair
	newPair.right.owner = newPair

	return newPair
}

func (pair *Pair) ProcesSplits() bool {
	left, right := pair.left, pair.right

	if left.IsNumber() {
		if left.Split() {
			return true
		}
	} else if left.tree.ProcesSplits() {
		return true
	}

	if right.IsNumber() {
		if right.Split() {
			return true
		}
	} else if right.tree.ProcesSplits() {
		return true
	}

	return false
}

func (pair *Pair) ProcessExpodes(depth int) bool {
	left, right := pair.left.tree, pair.right.tree
	if left == nil && right == nil {
		if depth >= 4 {
			pair.Explode()
			return true
		}
	} else {
		if left != nil && left.ProcessExpodes(depth+1) {
			return true
		}
		if right != nil && right.ProcessExpodes(depth+1) {
			return true
		}
	}
	return false
}

func (pair *Pair) Reduce() {
	for {
		if pair.ProcessExpodes(0) {
			continue
		}
		if pair.ProcesSplits() {
			continue
		}
		break
	}
}

func (pair *Pair) CalcMagnitude() int {
	var left, right int
	if pair.left.tree == nil {
		left = pair.left.value
	} else {
		left = pair.left.tree.CalcMagnitude()
	}
	if pair.right.tree == nil {
		right = pair.right.value
	} else {
		right = pair.right.tree.CalcMagnitude()
	}
	return 3*left + 2*right
}

func parseItem(s string, pos *int, parent *Pair, lastItem **Item) *Item {
	item := &Item{owner: parent}

	if s[*pos] == '[' {
		item.tree = parsePair(s, pos, parent, lastItem)
	} else {
		if *lastItem != nil {
			item.previous = *lastItem
			(*lastItem).next = item
		}
		*lastItem = item

		startPos := *pos
		for s[*pos] >= '0' && s[*pos] <= '9' {
			(*pos)++
		}
		item.value, _ = strconv.Atoi(s[startPos:*pos])
	}

	return item
}

func parsePair(s string, pos *int, parent *Pair, lastItem **Item) *Pair {
	pair := &Pair{parent: parent}

	if s[*pos] != '[' {
		log.Fatal(errors.New("missing '['"))
	}
	(*pos)++

	pair.left = parseItem(s, pos, pair, lastItem)

	if s[*pos] != ',' {
		log.Fatal(errors.New("missing ','"))
	}
	(*pos)++

	pair.right = parseItem(s, pos, pair, lastItem)

	if s[*pos] != ']' {
		log.Fatal(errors.New("missing ']'"))
	}
	(*pos)++

	return pair
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var pair *Pair
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
		pos := 0
		var lastItem *Item = nil
		p := parsePair(line, &pos, nil, &lastItem)
		if pair == nil {
			pair = p
		} else {
			pair = pair.Append(p)
			pair.Reduce()
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(lines) == 1 {
		pair.Reduce()
	}
	fmt.Println("Part 1", "=", pair.CalcMagnitude())

	maxMagnitude := math.MinInt
	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			for d := 0; d < 2; d++ {
				pos1, pos2 := 0, 0
				var lastItem1, lastItem2 *Item = nil, nil
				p1 := parsePair(lines[i], &pos1, nil, &lastItem1)
				p2 := parsePair(lines[j], &pos2, nil, &lastItem2)
				var p *Pair
				if d == 0 {
					p = p1.Append(p2)
				} else {
					p = p2.Append(p1)
				}
				p.Reduce()
				m := p.CalcMagnitude()
				if m > maxMagnitude {
					maxMagnitude = m
				}
			}
		}
	}
	fmt.Println("Part 2", "=", maxMagnitude)
}
