package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
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
		v := n / 3
		total += v - 2
	}
	fmt.Println(total)
}
