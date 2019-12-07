package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func b() {
	// Fuel required to launch a given module is based on its mass.
	// Specifically, to find the fuel required for a module, take its mass,
	// divide by three, round down, and subtract 2.
	total := 0
	f, err := os.Open("./input-a")
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
		total += getFuel(n)
	}
	fmt.Println("B:", total)
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
