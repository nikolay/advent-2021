package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	mx = 5
	my = 5
)

type Coord struct {
	x, y int
}

type Item struct {
	value    Coord
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Update(item *Item, value Coord, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func solve(x1, y1, x2, y2 int, risks [][]int) int {
	steps := []Coord{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	width, height := len(risks[0]), len(risks)

	queue := make(PriorityQueue, 0, width*height)
	heap.Init(&queue)

	distances := make([][]int, height)
	for y := 0; y < height; y++ {
		row := make([]int, width)
		for x := 0; x < width; x++ {
			row[x] = math.MaxInt
		}
		distances[y] = row
	}

	distances[y1][x1] = 0
	queue.Push(&Item{Coord{x1, y1}, 0, 0})
	for len(queue) > 0 {
		item := heap.Pop(&queue).(*Item)
		node, distance := item.value, item.priority
		if distance < distances[node.y][node.x] {
			continue
		}
		for _, step := range steps {
			x, y := node.x+step.x, node.y+step.y
			if x < 0 || x >= width || y < 0 || y >= height {
				continue
			}
			d := distance + risks[y][x]
			if d < distances[y][x] {
				distances[y][x] = d
				queue.Push(&Item{Coord{x, y}, d, 0})
			}
		}
	}
	return distances[y2][x2]
}

func transform(matrix [][]int, mx, my int) [][]int {
	width, height := len(matrix[0]), len(matrix)
	newWidth, newHeight := width*mx, height*my
	newMatrix := make([][]int, newHeight)
	for y := 0; y < newHeight; y++ {
		newMatrix[y] = make([]int, newWidth)
		for x := 0; x < newWidth; x++ {
			v := matrix[y%height][x%width] + (x / width) + (y / height)
			for v > 9 {
				v -= 9
			}
			newMatrix[y][x] = v
		}
	}
	return newMatrix
}

var part int

func init() {
	if len(os.Args) > 0 {
		if p, err := strconv.Atoi(os.Args[1]); err != nil {
			log.Fatal(err)
		} else if p < 1 || p > 2 {
			log.Fatal(errors.New(fmt.Sprintf("invalid part: %v", p)))
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

	risks := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		row := make([]int, 0, len(line))
		for _, b := range line {
			risk := int(b - '0')
			row = append(row, risk)
		}
		risks = append(risks, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if part == 2 {
		risks = transform(risks, mx, my)
	}
	width, height := len(risks[0]), len(risks)
	fmt.Println(solve(0, 0, width-1, height-1, risks))
}
