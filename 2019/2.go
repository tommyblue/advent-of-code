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
	for i := 0; i+4 <= len(vals); i += 4 {
		opcode := vals[i]
		input1 := vals[i+1]
		input2 := vals[i+2]
		output := vals[i+3]
		if opcode == 1 {
			vals[output] = vals[input1] + vals[input2]

		} else if opcode == 2 {
			vals[output] = vals[input1] * vals[input2]
		}
	}
	fmt.Println("2A:", vals[0])
}

func b2() {

}
