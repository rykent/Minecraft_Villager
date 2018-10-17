package villager

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"github.com/tmthrgd/go-memset"
)

var program []int

func interpret(input string) bool {

	return false
}

func main() {
	file, err := os.Open(os.Args[1])
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
		n, readerr := reader.ReadBytes()
		if readerr == io.EOF {
			//do nothing yet
			read = false
		} else if readerr != nil {
			fmt.Printf("Read Error")
			break
		}
		var cmd int = 0

		interpret(string(buf[:n]))
	}
}
