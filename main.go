package main

import (
	"fmt"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/wasd424/Minecraft_Villager/interpreter"
)

func main() {
	parser := argparse.NewParser("Minecraft_Villager interpreter",
		"Allows minecraft villagers to code.")

	t := parser.String("t", "Task", &argparse.Options{Required: true,
		Help: "Task to run"})
	r := parser.String("r", "File", &argparse.Options{Required: true,
		Help: "Input file"})

	err := parser.Parse(os.Args)
	if err != nil {
		panic(err.Error)
	}

	switch *t {
	case "interpret":
		fmt.Printf("Interpreting file %s", r)
		interpreter.interpreter(r)
	default:
		fmt.Println("Invalid Task")
	}
}
