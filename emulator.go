package main

import (
	"fmt"
	// "github.com/veandco/go-sdl2/sdl"
	"os"
)

const MEMSIZ = 4096

type CPU struct {
	ROM    [MEMSIZ]uint8
	PC     uint16
	STACK  [16]uint16
	V      [16]uint8
	I, SP  uint16
	DT, ST uint8
}

func main() {
	mustQuit := false
	cpu := CPU{}

	init_machine(&cpu)
	init_machine_rom(&cpu)

	for !mustQuit {
		var opcode uint16 = uint16(cpu.rom[cpu.PC])<<8 | uint16(cpu.rom[cpu.PC+1])
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
}

func init_machine(cpu *CPU) {
	cpu.PC = 0x200
}

func init_machine_rom(cpu *CPU) {
	stream, err := os.Open("roms/rom")
	if err != nil {
		panic(err)
	}

	SIZE, errsize := stream.Stat()
	if errsize != nil {
		panic(errsize)
	}

	rom := make([]byte, SIZE.Size())
	b_rom, err := stream.Read(rom)
	if err != nil {
		panic(err)
	}
	if (b_rom + 0x200) > MEMSIZ {
		panic("The ROM exceeds the maximum memory size")
	}
	fmt.Printf("%d bytes readed..\n", b_rom)
	copy(machine.rom[0x200:(0x200+b_rom)], rom)
}
