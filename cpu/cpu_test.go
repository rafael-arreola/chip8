package cpu

import (
	"testing"
)

func TestNew(t *testing.T) {
	cpu := New()
	if cpu.PC != 0x200 {
		t.Errorf("The initial value on PC registers is not 0x200")
	}
}
