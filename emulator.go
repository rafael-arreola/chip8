package main

import (
	"fmt"
	// "github.com/veandco/go-sdl2/sdl"
	"github.com/rafael-arreola/chip8/cpu"
)

func main() {
	//mustQuit := false

	pc := cpu.New()
	pc.Load("roms/rom")
	/*
		init_machine(&cpu)
		init_machine_rom(&cpu)

		for !mustQuit {
			var opcode uint16 =
			var nnn uint16 = opcode & 0xFFF
			var kk uint8 = uint8(opcode & 0xFF)
			var n uint8 = uint8(opcode & 0xF)
			var x uint8 = uint8(opcode >> 8 & 0xF)
			var y uint8 = uint8(opcode >> 4 & 0xF0)

			fmt.Printf("opcode: %X \n", opcode)
			fmt.Printf("nnn: %X \n", nnn)
			fmt.Printf("kk: %X \n", kk)
			fmt.Printf("n: %X \n", n)
			fmt.Printf("x: %X \n", x)
			fmt.Printf("y: %X \n", y)
			fmt.Printf("AH: %X \n", cpu.rom[cpu.PC])
			fmt.Printf("AL: %X \n", cpu.rom[cpu.PC+1])

			mustQuit = true
			if cpu.PC+2 == MEMSIZ {
				cpu.PC = 0
				mustQuit = true
			} else {
				cpu.PC += 2
			}

		}
	*/
}
