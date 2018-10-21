package main

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
	"os"

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
		fmt.Printf("Interpreting file %s\n", *r)
		interpreter.Interpret(*r)
	default:
		fmt.Printf("Invalid Task\n")
	}
}
