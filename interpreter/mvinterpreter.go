package interpreter

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"strings"

	"github.com/tmthrgd/go-memset"
)

/*
* Program
* Contains the commands to be executed.
*/
var program = make([]int, 0)

/*
* Program Counter
* Tells what command in program should be executed
*/
var pc int = 0

var mem = make([]byte, 0) //Program Memory
var mem_pos = 0 //Memory Position

func exec(instruction int) {
	switch instruction {
	case 0:
		if pc == 0 {
			fmt.Printf("Error.Don't use \"hmm\" as the first cmd.")
			os.Exit(1)
		}

		pc--

		level := 1
		for level > 0 {
			if pc == 0 {
				break
			}

			pc--

			if program[pc] == 0 {
				level++
			} else if program[pc] == 7 {
				level--
			}
		}
		if level != 0 {
			fmt.Printf("Error.\n")
			os.Exit(1)
		}
		exec(program[pc])
	case 1:
		if mem_pos == 0 {
			os.Exit()
		} else {
			mem_pos--
		}
	case 2:
		mem_pos++
		if mem_pos == len(mem) {
			mem = append(mem, 0)
			mem_pos = len(mem) - 1
		}
	case 3:
		if mem[mem_pos] == 3 {
			fmt.Printf("Error.\n")
			os.Exit(1)
		}
		exec(mem[mem_pos])
	case 4:

        }
}

func interpreter(f string) {
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("Cannot open source file")
		panic(err.Error)
	}

	//Close the file
	defer file.Close()

	reader := bufio.NewReader(file)

	buf := make([]byte, 14)
	Memset(buf, 0)

	/*
	* Interpret each command
	* Read until EOF
	*/
	read := true


	for read {
		n, readerr := reader.ReadBytes('\n')
		if readerr == io.EOF {
			read = false
		} else if readerr != nil {
			panic(readerr.Error)
		}

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
			fmt.Printf("ERROR. Invalid command.\n")
			os.Exit(1)
		}
	}

	if input[0] != "h" {
		fmt.Println("Syntax Error. Use h%s, not %s", input, input)
	}


	//Init main memory
	mem = append(mem, 0)
	mem_pos = mem[0]

	pc = program[0]

	for len(program) > 0 {
		exec(program[0])
	}
}
