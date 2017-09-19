package main

import (
	"fmt"
	// "github.com/veandco/go-sdl2/sdl"
	"github.com/rafael-arreola/chip8/cpu"
)

func main() {
	running := true
	pc := cṕu.New()
	pc.Load("roms/rom")

	for running {

		opcode, nnn, n, x, y, kk, i := pc.DecodeOp()
		switch i {
		case 0x0:
			switch nnn {
			case 0x0E0:
				fmt.Println("CLS")
			case 0x00EE:
				fmt.Println("RET")
			default:
				fmt.Printf("SYS %X\n", nnn)
			}
		case 0x1:
			pc.PC = nnn
		case 0x2:
			fmt.Printf("CALL %X\n", nnn)
		case 0x3:
			if pc.V[x] == kk {

			}
		case 0x4:
			if pc.V[x] != kk {

			}
		case 0x5:
			if pc.V[x] == pc.V[y] {

			}
		case 0x6:
			pc.V[x] = kk

		case 0x7:
			pc.V[x] = pc.V[x] + kk

		case 0x8:
			switch n {
			case 0x0:
				pc.V[x] = pc.V[y]

			case 0x1:
				pc.V[x] = pc.V[x] | pc.V[y]

			case 0x2:
				pc.V[x] = pc.V[x] & pc.V[y]

			case 0x3:
				pc.V[x] = pc.V[x] ^ pc.V[y]

			case 0x4:
				var cf uint8
				if (uint16(pc.V[x]) + uint16(pc.V[y])) > uint16(pc.V[x]) {
					cf = 1
				}
				pc.V[0xF] = cf
				pc.V[x] = uint8((uint16(pc.V[x]) + uint16(pc.V[y])) & 0xFF)

			case 0x5:
				var cf uint8
				if pc.V[x] > pc.V[y] {
					cf = 1
				}
				pc.V[0xF] = cf
				pc.V[x] = pc.V[x] - pc.V[y]

			case 0x6:
				var cf uint8
				if (pc.V[x] & 0x01) == 0x01 {
					cf = 1
				}
				pc.V[0xF] = cf
				pc.V[x] = pc.V[x] / 2

			case 0x7:
				var cf uint8
				if pc.V[y] > pc.V[x] {
					cf = 1
				}
				pc.V[0xF] = cf
				pc.V[x] = pc.V[y] - pc.V[x]

			case 0xE:
				var cf uint8
				if (pc.V[x] & 0x80) == 0x80 {
					cf = 1
				}
				pc.V[0xF] = cf

				pc.V[x] = pc.V[x] * 2
			}
		case 0x9:
			switch n {
			case 0x0:
				if pc.V[x] != pc.V[x] {

				}
			}
		case 0xA:
			pc.I = nnn

		case 0xB:
			pc.PC = nnn + uint16(pc.V[0])
		case 0xC:
			pc.V[x] = kk + randByte()
		case 0xD:
			fmt.Printf("DRW V[%X], V%X, %X\n", x, y, n)
		case 0xE:
			switch nnn & 0xFF {
			case 0x9E:
				fmt.Printf("SKP %X\n", x)
			case 0xA1:
				fmt.Printf("SKNP %X\n", x)
			}
		case 0xF:
			switch nnn & 0XFF {
			case 0x07:
				fmt.Printf("LD V[%X], DT\n", x)
			case 0x0A:
				fmt.Printf("LD V[%X], K\n", x)
			case 0x15:
				fmt.Printf("LD DT, V[%X]\n", x)
			case 0x18:
				fmt.Printf("LD ST, V[%X]\n", x)
			case 0x1E:
				fmt.Printf("ADD I, V[%X]\n", x)
			case 0x29:
				fmt.Printf("LD F, V[%X]\n", x)
			case 0x33:
				fmt.Printf("LD B, V[%X]\n", x)
			case 0x55:
				fmt.Printf("LD [I], V[%X]\n", x)
			case 0x65:
				fmt.Printf("LD V[%X], [I]\n", x)
			}
		}

		pc.Next()

	}
}
