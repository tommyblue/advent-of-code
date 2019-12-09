package main

import (
	"fmt"
	"strings"

	"github.com/tommyblue/advent-of-code/2019/utils"
)

func a2() {
	r := utils.NewFileReader("./inputs/2")
	line, _, err := r.ReadLine()
	if err != nil {
		panic(err)
	}
	strVals := strings.Split(string(line), ",")
	vals := []int{}
	for _, v := range strVals {
		vals = append(vals, utils.ToInt(v))
	}
	vals[1] = 12
	vals[2] = 2
	ret := calc(vals)
	fmt.Println("2A:", ret)
}

func b2() {
	noun := 0
	verb := 0

	r := utils.NewFileReader("./inputs/2")
	line, _, err := r.ReadLine()
	if err != nil {
		panic(err)
	}
	strVals := strings.Split(string(line), ",")
	source := []int{}
	for _, v := range strVals {
		source = append(source, utils.ToInt(v))
	}

	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			vals := copy(source)
			vals[1] = n
			vals[2] = v
			ret := calc(vals)
			if ret == 19690720 {
				noun = n
				verb = v
				break
			}
		}
	}
	fmt.Println("2B:", 100*noun+verb)

}

func calc(source []int) int {
	vals := copy(source)
	index := 0
	l := len(vals)
	for index < l-1 {
		opcode := vals[index]
		input1 := vals[index+1]
		input2 := vals[index+2]
		output := vals[index+3]
		if opcode == 1 {
			vals[output] = vals[input1] + vals[input2]
		} else if opcode == 2 {
			vals[output] = vals[input1] * vals[input2]
		}
		index += 4
	}
	return vals[0]
}

func copy(s []int) []int {
	r := []int{}
	r = append(r, s...)
	return r
}
