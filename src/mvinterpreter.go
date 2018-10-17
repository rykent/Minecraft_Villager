package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"strings"
)

var (
    program []byte
    pc int
)

func Memset(a []byte, v byte) {
    if len(a) == 0 {
        return
    }
    a[0] = v
    for bp := 1; bp < len(a); bp *= 2 {
        copy(a[bp:], a[:bp])
    }
}

func exec() {
    for i := int; i < len(program); i++ {
        switch program[i] {
            
        }
    }
}

func interpret(input string) {
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

func main() {
	file, err := os.Open(os.Args[0])
	if err != nil {
		fmt.Println("Cannot open source file")
		os.Exit(1)
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
			//do nothing yet
			read = false
		} else if readerr != nil {
			fmt.Printf("Read Error, %d amount of bytes read.\n", n)
			os.Exit(1)
		}

		interpret(string(buf))
		exec()
	}
}
