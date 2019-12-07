package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/tommyblue/advent-of-code/2019/utils"
)

func a1() {
	// Fuel required to launch a given module is based on its mass.
	// Specifically, to find the fuel required for a module, take its mass,
	// divide by three, round down, and subtract 2.
	total := 0
	f, err := os.Open("./inputs/1")
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(f)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		n, err := strconv.Atoi(string(line))
		if err != nil {
			panic(err)
		}
		v := n / 3
		total += v - 2
	}
	fmt.Println("1A:", total)
}

func b1() {
	// Fuel required to launch a given module is based on its mass.
	// Specifically, to find the fuel required for a module, take its mass,
	// divide by three, round down, and subtract 2.
	total := 0
	r := utils.NewFileReader("./inputs/1")
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		n := utils.ToInt(string(line))
		total += getFuel(n)
	}
	fmt.Println("1B:", total)
}

func getFuel(mass int) int {
	fuel := (mass / 3) - 2
	if fuel > 0 {
		fuel += getFuel(fuel)
	} else if fuel <= 0 {
		return 0
	}
	return fuel
}
