package cpu

import (
	"fmt"
	"os"
)

const MEMSIZ = 4096

/*
	+---------------+= 0xFFF (4095) End of Chip-8 RAM
	|               |
	|               |
	|               |
	|               |
	|               |
	| 0x200 to 0xFFF|
	|     Chip-8    |
	| Program / Data|
	|     Space     |
	|               |
	|               |
	|               |
	+- - - - - - - -+= 0x600 (1536) Start of ETI 660 Chip-8 programs
	|               |
	|               |
	|               |
	+---------------+= 0x200 (512) Start of most Chip-8 programs
	| 0x000 to 0x1FF|
	| Reserved for  |
	|  interpreter  |
	+---------------+= 0x000 (0) Start of Chip-8 RAM
*/
type CPU struct {
	ROM    [MEMSIZ]uint8
	PC     uint16
	STACK  [16]uint16
	V      [16]uint8
	I, SP  uint16
	DT, ST uint8
}

func New() CPU {
	return CPU{
		PC: 0x200,
	}
}

func (cpu *CPU) Load(src string) {
	stream, err := os.Open(src)
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
	copy(cpu.ROM[0x200:(0x200+b_rom)], rom)
}

/*
func (cpu *CPU) DecodeOp() uint16 {
	return uint16(cpu.rom[cpu.PC])<<8 | uint16(cpu.rom[cpu.PC+1])
}
*/
