package main

import (
	"fmt"
)

//3 since our instructions are
//not going to have more than 3 args
//in our simple addition program below

type instruction [3]interface{}
type instructions []instruction

//simple addition program
//push number 5 to the stack
//prog[0] = instruction{"PSH", 5}
//push number 6 to the stack
//prog[1] = instruction{"PSH", 6}
//add the values in stack
//prog[2] = instruction{"ADD"}
//pop the value from stack
//prog[3] = instruction{"POP"}
//halt the system
//prog[4] = instruction{"HLT"}

var prog = instructions{
	instruction{"PSH", 5},
	instruction{"PSH", 6},
	instruction{"ADD"},
	instruction{"POP"},
	instruction{"HLT"},
}

var stack [256]int

var sp int = 0

//Set running to true
var running = true
var ip = 0

//program [256]instructions, ip int
func fetch() instruction {
	// return ipth instruction of program
	return prog[ip]
}
func evalinstr(running *bool, instr instruction) {
	//
	switch instr[0] {
	case "PSH":
		//push value to the stack
		tmp := instr[1].(int)
		stack[sp] = tmp
		sp++
	case "ADD":
		//add the values in stack
		a := stack[sp-2]
		b := stack[sp-1]
		result := a + b
		sp++
		stack[sp] = result
		sp++
	case "POP":
		//pop the value from stack
		tmp := stack[sp-1]
		fmt.Println(tmp)
	case "HLT":
		//halt the system
		fmt.Println("HLT: system halted!")
		*running = false

	}
}

func main() {
	fmt.Println("Starting emulation")
	for running == true {
		evalinstr(&running, fetch())
		ip++
	}
}
