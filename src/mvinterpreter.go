package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"string"
	"github.com/tmthrgd/go-memset"
)

var program []int

func interpret(input string) {
	mcount := strings.Count(input, "m")
	switch m {
	case 2:

	case 3:

	case 4:

	case 5:

	case 6:

	case 7:

	case 8:

	case 9:

	case 10:

	case 11:

	case 12:

	case 13:

	default:
		fmt.Printf("ERROR. Invalid command.\n")
	}
	return false
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
	memset.Memset(buf, 0)

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
			break
		}

		interpret(string(buf))
	}
}
