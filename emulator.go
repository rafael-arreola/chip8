package main

import (
	"fmt"
	// "github.com/veandco/go-sdl2/sdl"
	"github.com/rafael-arreola/chip8/cpu"
)

func main() {
	running := true
	pc := cpu.New()
	pc.Load("roms/rom")

	for running {

		opcode, nnn, n, x, y, kk, i := pc.DecodeOp()
		pc.Next()
		// fmt.Printf("opcode: %X, nnn: %X, n: %X, x: %X, y: %X, kk: %X , i: %X \n", opcode, nnn, n, x, y, kk, i)
		opcode++
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
			pc.JP(nnn)
		case 0x2:
			fmt.Printf("CALL %X\n", nnn)
		case 0x3:
			pc.SE(x, kk)
		case 0x4:
			pc.SNE(x, kk)
		case 0x5:
			pc.SEy(x, y)
		case 0x6:
			pc.LD(x, kk)
		case 0x7:
			pc.ADD(x, kk)
		case 0x8:
			switch n {
			case 0x0:
				pc.LDy(x, y)
			case 0x1:
				pc.OR(x, y)
			case 0x2:
				pc.AND(x, y)
			case 0x3:
				pc.XOR(x, y)
			case 0x4:
				pc.ADD(x, y)
			case 0x5:
				pc.SUB(x, y)
			case 0x6:
				pc.SHR(x, y)
			case 0x7:
				pc.SUBN(x, y)
			case 0xE:
				pc.SHL(x, y)
			}
		case 0x9:
			switch n {
			case 0x0:
				pc.SNE(x, y)
			}
		case 0xA:
			fmt.Printf("LD I, %X\n", nnn)
		case 0xB:
			fmt.Printf("JP V[0], %X\n", nnn)
		case 0xC:
			fmt.Printf("RND V[%X], byte\n", x, kk)
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

	}
}
