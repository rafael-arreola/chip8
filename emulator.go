package main

import (
	"fmt"
	// "github.com/veandco/go-sdl2/sdl"
	"os"
)

const MEMSIZ = 4096

type Machine_t struct {
	mem    [MEMSIZ]uint8
	pc     uint16
	stack  [16]uint16
	sp     uint16
	v      [16]uint8
	i      uint16
	dt, st uint8
}

func readFile(m *Machine_t) {
	stream, err := os.Open("roms/Alien 8")
	if err != nil {
		panic(err)
	}
	stream.Read(&m.mem)

}

func main() {
	machine := Machine_t{}
	readFile(&machine)
	fmt.Println("Hello World")
}
