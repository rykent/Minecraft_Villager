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
* Program Top
* Contains the commands to be executed.
* exec() always runs the command specified in program[0]
* Commands are popped and pushed to program bottom when executed.
*/
var prog_top = make([]int, 0)

/*
* Program bottom
* Contains the commands that have been executed
* Commands popped and pushed here can be moved back to prog_top to create a loop
*/
var prog_bottom = make([]int, 0)

var mem = make([]byte, 0) //Program Memory
var mem_pos = 0 //Memory Position

func prog_top_pop() {
	prog_bottom = append(prog_bottom, 0)
	copy(prog_bottom[1:], prog_bottom[0:])
	prog_bottom = prog_top[0]
}

func prog_top_push() {
	prog_top = append(prog_top, 0)
	copy(prog_top[1:], prog_top[0:])
	prog_top = prog_bottom[0]
}

func exec(instruction int) {
	switch instruction {
	case 0:
		if prog_top[0] == 0 {
			fmt.Printf("Error.
				Don't use \"hmm\" as the first command.")
			os.Exit(1)
		}

		level := 0
		for level > 0 {

			/*
			* If the bottom half (already executed part) has a
			* length of 0, then prog_top[0] is the program start.
			*/
			if len(prog_bottom) == 0 {
				break
			}

			pc--

			if prog_top[0] == 0 {
				level++
			} else if prog_top[0] == 7 {
				level--
			}
		}
		if level != 0 {
			fmt.Printf("Error.")
			os.Exit(1)
		}
		exec(pc)
	case 1:
		if mem_pos == mem[0] {
			os.Exit()

		}
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

	if input[0] != "h" {
		fmt.Println("Syntax Error. Use h%s, not %s", input, input)
	}
	mcount := strings.Count(input, "m")
	switch mcount {
	case 2:
		prog_top = append(prog_top, 0)
	case 3:
		prog_top = append(prog_top, 1)
	case 4:
		prog_top = append(prog_top, 2)
	case 5:
		prog_top = append(prog_top, 3)
	case 6:
		prog_top = append(prog_top, 4)
	case 7:
		prog_top = append(prog_top, 5)
	case 8:
		prog_top = append(prog_top, 6)
	case 9:
		prog_top = append(prog_top, 7)
	case 10:
		prog_top = append(prog_top, 8)
	case 11:
		prog_top = append(prog_top, 9)
	case 12:
		prog_top = append(prog_top, 10)
	case 13:
		prog_top = append(prog_top, 11)
	default:
		fmt.Printf("ERROR. Invalid command.\n")
		os.Exit(1)
	}

	for read {
		n, readerr := reader.ReadBytes('\n')
		if readerr == io.EOF {
			//do nothing yet
			read = false
		} else if readerr != nil {
			panic(readerr.Error)
		}

		interpret(string(buf))
	}

	//Init main memory
	mem = append(mem, 0)
	mem_pos = mem[0]

	pc = prog_top[0]

	for len(prog_top) > 0 {
		exec(prog_top[0])
	}
}
