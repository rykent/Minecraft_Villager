package interpreter

/*
Copyright 2018 Ryken Thompson (wasd424)

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"strings"
	"strconv"

	"github.com/tmthrgd/go-memset"
)

var (
	// Program contains the cmds to be executed.
	program = make([]int, 0)

	// Program counter is the current position in program[].
	pc int

	mem = make([]int, 0) //Program Memory
	memPos int //Memory Position

	regsVal int
	isRegsVal bool = false
)


func cmd0() {
	if pc == 0 {
		fmt.Printf("Don't use \"hmm\" as the first cmd.\n")
		os.Exit(1)
	}

	pc--

	lvl := 1
	for lvl > 0 {
		if pc == 0 {
			break
		}

		pc--

		if program[pc] == 0 {
			lvl++
		} else if program[pc] == 7 {
			lvl--
		}
	}
	if lvl != 0 {
		fmt.Printf("Error.\n")
		os.Exit(1)
	}
	exec(program[pc])
}

func cmd1() {
	if memPos == 0 {
		os.Exit(1)
	} else {
		memPos--
	}
}

func cmd2() {
	memPos++
	if memPos == len(mem) {
		mem = append(mem, 0)
		memPos = len(mem) - 1
	}
}

func cmd3() {
	if mem[memPos] == 3 {
		fmt.Printf("Error.\n")
		os.Exit(1)
	}
	exec(mem[memPos])
}

func cmd4(reader *bufio.Reader) {
	if mem[memPos] != 0 {
		fmt.Printf("%c", mem[memPos])
	} else {
		input, _ := reader.ReadByte()
		mem[memPos] = int(input)
	}
}

// func cmd5 and cmd6 wouldn't add any readability

func cmd7() {
	if mem[memPos] == 0 {

			lvl := 1
			prev := 0
			pc++

			for lvl > 0 {
				prev = program[pc]
				pc++
				if pc == len(program) {
					break
				}

				if program[pc] == 7 {
					lvl++
				} else if program[pc] == 0 {
					lvl--
					if prev == 7 {
						lvl--
					}
				}
			}
			if lvl != 0 {
				fmt.Printf("Error. Check your code.\n")
				os.Exit(1)
			}
	}
}

// func cmd8 wouldn't add any readability

func cmd9() {
	if isRegsVal {
		mem[memPos] = regsVal
	} else {
		regsVal = mem[memPos]
	}
	isRegsVal = !isRegsVal
}

// func cmd10 wouldn't add any readability

func cmd11(reader *bufio.Reader) {
	buf := make([]byte, 100)
	c := 0

	var err error

	for c < len(buf)-1 {
		input, _ := reader.ReadByte()
		buf[c] = input
		c++
		buf[c] = 0

		if buf[c-1] == '\n' {
			break
		}
	}

	if c == len(buf) {
		_, err = reader.ReadBytes('\n')
		if err != nil {
			fmt.Printf("Error getting input.\n")
			os.Exit(1)
		}
	}

	mem[memPos], err = strconv.Atoi(string(buf))

	if err != nil {
		fmt.Printf("Error converting input to string.\n")
		os.Exit(1)
	}
}

func exec(instruction int) {
	reader := bufio.NewReader(os.Stdin)

	switch instruction {

	// hmm
	case 0:
		cmd0()
	//hmmm
	case 1:
		cmd1()

	//hmmmm
	case 2:
		cmd2()

	// hmmmmm
	case 3:
		cmd3()

	// hmmmmmm
	case 4:
		cmd4(reader)

	// hmmmmmmm
	case 5:
		mem[memPos]--

	// hmmmmmmmm
	case 6:
		mem[memPos]++

	// hmmmmmmmmm
	case 7:
		cmd7()

	// hmmmmmmmmmm
	case 8:
		mem[memPos] = 0

	// hmmmmmmmmmmm
	case 9:
		cmd9()

	// hmmmmmmmmmmmm
	case 10:
		fmt.Printf("%d\n", mem[memPos])

	// hmmmmmmmmmmmmm
	case 11:
		cmd11(reader)

	default:
		fmt.Printf("Unrecognized command.\n")
		os.Exit(1)
	}

	pc++
}

func strToCmd(input string, line int) {
	mcount := strings.Count(input, "m")
	switch mcount {
	case 2:
		program = append(program, 0)
	case 3:
		program = append(program, 1)
	case 4:
		program = append(program, 2)
	case 5:
		program = append(program, 3)
	case 6:
		program = append(program, 4)
	case 7:
		program = append(program, 5)
	case 8:
		program = append(program, 6)
	case 9:
		program = append(program, 7)
	case 10:
		program = append(program, 8)
	case 11:
		program = append(program, 9)
	case 12:
		program = append(program, 10)
	case 13:
		program = append(program, 11)
	default:
		fmt.Printf("ERROR. Invalid command. At Line %d\n", line)
		os.Exit(1)
	}
}

// Interpret reads from file f and puts the cmds into program[].
func Interpret(f string) {
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("Cannot open source file")
		panic(err.Error)
	}

	//Close the file
	defer file.Close()

	reader := bufio.NewReader(file)

	buf := make([]byte, 14)
	memset.Memset(buf, 0)

	/*
	* Interpret each command
	* Read until EOF
	*/
	read := true

	var input string
	line := 0

	for read {
		n, readerr := reader.ReadBytes('\n')
		if readerr == io.EOF {
			read = false
		} else if readerr != nil {
			panic(readerr.Error)
		}

		input = string(n)

		hCount := strings.Count(input, "h")

		if hCount < 1 {
			fmt.Printf("Syntax Error, at line %d. Only use hmm...",
				line)
			os.Exit(1)
		}

		strToCmd(input, line)
		line++
	}


	//Init main memory
	mem = append(mem, 0)
	memPos = 0

	pc = 0

	for pc != len(program) {
		exec(program[pc])
		//fmt.Printf("Currently at program[%d]\n", pc)
	}
	os.Exit(0)
}
