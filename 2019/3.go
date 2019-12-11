package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tommyblue/advent-of-code/2019/utils"
)

func a3() {
	r := utils.NewFileReader("./inputs/3")
	line1, _, err := r.ReadLine()
	if err != nil {
		panic(err)
	}
	path1 := strings.Split(string(line1), ",")
	w1 := &wire{}
	w1.run(path1)

	line2, _, err := r.ReadLine()
	if err != nil {
		panic(err)
	}
	path2 := strings.Split(string(line2), ",")
	w2 := &wire{}
	w2.run(path2)

	positions := findCommonPositions(w1, w2)
	p := evaluateDistances(positions)
	fmt.Println("3A:", p)
}
func b3() {
	r := utils.NewFileReader("./inputs/3")
	line1, _, err := r.ReadLine()
	if err != nil {
		panic(err)
	}
	path1 := strings.Split(string(line1), ",")
	w1 := &wire{}
	w1.run(path1)

	line2, _, err := r.ReadLine()
	if err != nil {
		panic(err)
	}
	path2 := strings.Split(string(line2), ",")
	w2 := &wire{}
	w2.run(path2)

	minSteps := findCommonPositionsWithSteps(w1, w2)
	fmt.Println("3B:", minSteps)
}

func evaluateDistances(positions []string) int {
	minDistance := 0
	for _, position := range positions {
		xy := strings.Split(position, ":")
		x, err := strconv.Atoi(xy[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(xy[1])
		if err != nil {
			panic(err)
		}
		d := calculateDistance(x, y)
		if minDistance == 0 {
			minDistance = d
		}
		if d < minDistance {
			minDistance = d

		}
	}
	return minDistance
}

func calculateDistance(x, y int) int {
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}
	return x + y
}

type wire struct {
	positions map[string]int
}

func findCommonPositionsWithSteps(w1, w2 *wire) int {
	minSteps := 0
	for position, s1 := range w1.positions {
		if s2, ok := w2.positions[position]; ok {
			if minSteps == 0 {
				minSteps = s1 + s2
			}
			if s1+s2 < minSteps {
				minSteps = s1 + s2
			}
		}
	}
	return minSteps
}

func findCommonPositions(w1, w2 *wire) []string {
	common := []string{}
	for position := range w1.positions {
		if _, ok := w2.positions[position]; ok {
			common = append(common, position)
		}
	}
	return common
}

func (w *wire) add(x, y, steps int) {
	w.positions[fmt.Sprintf("%d:%d", x, y)] = steps
}

func (w *wire) run(path []string) {
	w.positions = make(map[string]int)
	currentX := 0
	currentY := 0
	steps := 0

	for _, v := range path {
		action := string(v[0])
		moveBy, err := strconv.Atoi(v[1:])
		if err != nil {
			panic(err)
		}

		if action == "R" { // right
			for i := 0; i < moveBy; i++ {
				currentX++
				steps++
				w.add(currentX, currentY, steps)
			}
		} else if action == "L" { // left
			for i := 0; i < moveBy; i++ {
				currentX--
				steps++
				w.add(currentX, currentY, steps)
			}
		} else if action == "U" { //up
			for i := 0; i < moveBy; i++ {
				currentY++
				steps++
				w.add(currentX, currentY, steps)
			}
		} else if action == "D" { //down
			for i := 0; i < moveBy; i++ {
				currentY--
				steps++
				w.add(currentX, currentY, steps)
			}
		}
	}
}
