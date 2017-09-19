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
	pc := CPU{
		PC: 0x200,
	}
	return pc
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
	fmt.Printf("%d bytes readed in the file [%s] ..\n", b_rom, src)
	copy(cpu.ROM[0x200:(0x200+b_rom)], rom)
}

/*
	get Opcode from current PC address

	nnn or addr - A 12-bit value, the lowest 12 bits of the instruction
	n or nibble - A 4-bit value, the lowest 4 bits of the instruction
	x - A 4-bit value, the lower 4 bits of the high byte of the instruction
	y - A 4-bit value, the upper 4 bits of the low byte of the instruction
	kk or byte - An 8-bit value, the lowest 8 bits of the instruction

	Example:
	    |------ nnn ------|
	                |- n -|
	    |- x -|
	    	  |- y -|
	4444  3333  2222  1111

	nnn = 3333 2222 1111
	n = 1111
	x = 3333
	y = 2222
	kk = 2222 1111
	i = 4444

*/
func (cpu *CPU) DecodeOp() (uint16, uint16, uint8, uint8, uint8, uint8, uint8) {
	opcode := uint16(cpu.ROM[cpu.PC])<<8 | uint16(cpu.ROM[cpu.PC+1])
	nnn := uint16(opcode & 0x0FFF)
	n := uint8(opcode & 0x000F)
	x := uint8((opcode >> 8) & 0xF)
	y := uint8((opcode >> 4) & 0xF)
	kk := uint8(opcode & 0x00FF)
	i := uint8((opcode >> 12))
	return opcode, nnn, n, x, y, kk, i
}

func (cpu *CPU) Next() {
	cpu.PC = (cpu.PC + 2) & 0xFFF
}

/*
	----------------- Standard Chip-8 Instructions -----------------
	S0nnn - SYS addr
	00E0 - CLS
	00EE - RET
>	1nnn - JP addr
>	2nnn - CALL addr
>	3xkk - SE Vx, byte
>	4xkk - SNE Vx, byte
>	5xy0 - SE Vx, Vy
>	6xkk - LD Vx, byte
>	7xkk - ADD Vx, byte
>	8xy0 - LD Vx, Vy
>	8xy1 - OR Vx, Vy
>	8xy2 - AND Vx, Vy
>	8xy3 - XOR Vx, Vy
>	8xy4 - ADD Vx, Vy
>	8xy5 - SUB Vx, Vy
>	8xy6 - SHR Vx {, Vy}
>	8xy7 - SUBN Vx, Vy
>	8xyE - SHL Vx {, Vy}
	9xy0 - SNE Vx, Vy
	Annn - LD I, addr
	Bnnn - JP V0, addr
	Cxkk - RND Vx, byte
	Dxyn - DRW Vx, Vy, nibble
	Ex9E - SKP Vx
	ExA1 - SKNP Vx
	Fx07 - LD Vx, DT
	Fx0A - LD Vx, K
	Fx15 - LD DT, Vx
	Fx18 - LD ST, Vx
	Fx1E - ADD I, Vx
	Fx29 - LD F, Vx
	Fx33 - LD B, Vx
	Fx55 - LD [I], Vx
	Fx65 - LD Vx, [I]
*/

func (cpu *CPU) JP(nnn uint16) {
	cpu.PC = nnn
}

func (cpu *CPU) SE(x uint8, kk uint8) {
	if cpu.V[x] == kk {
		cpu.Next()
	}
}
func (cpu *CPU) SEy(x uint8, y uint8) {
	if cpu.V[x] == cpu.V[y] {
		cpu.Next()
	}
}

func (cpu *CPU) SNE(x uint8, kk uint8) {
	if cpu.V[x] != kk {
		cpu.Next()
	}
}

func (cpu *CPU) LD(x uint8, kk uint8) {
	cpu.V[x] = kk
	cpu.Next()
}

func (cpu *CPU) ADD(x uint8, kk uint8) {
	cpu.V[x] = cpu.V[x] + kk
	cpu.Next()
}

func (cpu *CPU) LDy(x uint8, y uint8) {
	cpu.V[x] = cpu.V[y]
	cpu.Next()
}

func (cpu *CPU) OR(x uint8, y uint8) {
	cpu.V[x] = cpu.V[x] | cpu.V[y]
	cpu.Next()
}

func (cpu *CPU) AND(x uint8, y uint8) {
	cpu.V[x] = cpu.V[x] & cpu.V[y]
	cpu.Next()
}

func (cpu *CPU) XOR(x uint8, y uint8) {
	cpu.V[x] = cpu.V[x] ^ cpu.V[y]
	cpu.Next()
}

func (cpu *CPU) ADDy(x uint8, y uint8) {
	var cf uint8
	if (uint16(cpu.V[x]) + uint16(cpu.V[y])) > uint16(cpu.V[x]) {
		cf = 1
	}
	cpu.V[0xF] = cf
	cpu.V[x] = uint8((uint16(cpu.V[x]) + uint16(cpu.V[y])) & 0xFF)
	cpu.Next()
}

func (cpu *CPU) SUB(x uint8, y uint8) {
	var cf uint8
	if cpu.V[x] > cpu.V[y] {
		cf = 1
	}
	cpu.V[0xF] = cf
	cpu.V[x] = cpu.V[x] - cpu.V[y]
	cpu.Next()
}

func (cpu *CPU) SHR(x uint8, y uint8) {
	var cf uint8
	if (cpu.V[x] & 0x01) == 0x01 {
		cf = 1
	}
	cpu.V[0xF] = cf
	cpu.V[x] = cpu.V[x] / 2
	cpu.Next()
}

func (cpu *CPU) SUBN(x uint8, y uint8) {
	var cf uint8
	if cpu.V[y] > cpu.V[x] {
		cf = 1
	}
	cpu.V[0xF] = cf
	cpu.V[x] = cpu.V[y] - cpu.V[x]
	cpu.Next()
}

func (cpu *CPU) SHL(x uint8, y uint8) {
	var cf uint8
	if (cpu.V[x] & 0x80) == 0x80 {
		cf = 1
	}
	cpu.V[0xF] = cf

	cpu.V[x] = cpu.V[x] * 2
	cpu.Next()
}

func (cpu *CPU) SNE(x uint8, y uint8) {
	if cpu.V[x] != cpu.V[x] {
		cpu.Next()
	}
	cpu.Next()
}
