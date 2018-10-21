package interpreter

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
	/*
	* Program
	* Contains the commands to be executed.
	*/
	program = make([]int, 0)

	/*
	* Program Counter
	* Tells what command in program should be executed
	*/
	pc int = 0

	mem = make([]int, 0) //Program Memory
	mem_pos = 0 //Memory Position

	regsVal int
	isRegsVal bool = false
)

func exec(instruction int) {
	reader := bufio.NewReader(os.Stdin)

	switch instruction {

	// hmm
	case 0:
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

	//hmmm
	case 1:
		if mem_pos == 0 {
			os.Exit(1)
		} else {
			mem_pos--
		}

	//hmmmm
	case 2:
		mem_pos++
		if mem_pos == len(mem) {
			mem = append(mem, 0)
			mem_pos = len(mem) - 1
		}

	// hmmmmm
	case 3:
		if mem[mem_pos] == 3 {
			fmt.Printf("Error.\n")
			os.Exit(1)
		}
		exec(mem[mem_pos])

	// hmmmmmm
	case 4:
		if mem[mem_pos] != 0 {
			fmt.Printf("%c", mem[mem_pos])
		} else {
			input, _ := reader.ReadByte()
			mem[mem_pos] = int(input)
		}

	// hmmmmmmm
	case 5:
		mem[mem_pos]--

	// hmmmmmmmm
	case 6:
		mem[mem_pos]++

	// hmmmmmmmmm
	case 7:
		if mem[mem_pos] == 0 {

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

	// hmmmmmmmmmm
	case 8:
		mem[mem_pos] = 0

	// hmmmmmmmmmmm
	case 9:
		if isRegsVal {
			mem[mem_pos] = regsVal
		} else {
			regsVal = mem[mem_pos]
		}
		isRegsVal = !isRegsVal

	// hmmmmmmmmmmmm
	case 10:
		fmt.Printf("%d\n", mem[mem_pos])

	// hmmmmmmmmmmmm
	case 11:
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
		}

		mem[mem_pos], err = strconv.Atoi(string(buf))

		if err != nil {
			fmt.Printf("Error getting input.\n")
			os.Exit(1)
		}

	default:
		fmt.Printf("Unrecognized command.\n")
		os.Exit(1)
	}

	pc++
}


/*
* Interpret
* Opens the file specified in the string, f.
* Transforms each command in to a number and puts it into program[]
* Exec() program until program counter reaches end of program.
*/
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

	for read {
		n, readerr := reader.ReadBytes('\n')
		if readerr == io.EOF {
			read = false
		} else if readerr != nil {
			panic(readerr.Error)
		}

		input = string(n)

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

	if input[0] != 'h' {
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
