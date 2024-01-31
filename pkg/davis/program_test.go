package davis

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestProgram(t *testing.T) {
	g := NewWithT(t)

	instructions := []Instruction{
		DecodeInstruction(45),
		DecodeInstruction(34),
		DecodeInstruction(350),
		DecodeInstruction(2),
		DecodeInstruction(46),
	}
	inputs := []int{2, 1}
	program, err := NewProgram(instructions, inputs)
	g.Expect(err).ToNot(HaveOccurred())

	lines := []string{
		"1 2 1 0 0 0 0",
		"2 1 1 0 0 0 0",
		"3 1 1 0 0 1 0",
		"4 1 1 0 0 1 0",
		"5 1 1 0 0 1 1",
		"1 1 1 0 0 1 1",
		"2 0 1 0 0 1 1",
		"3 0 1 0 0 2 1",
		"4 0 1 0 0 2 1",
		"5 0 1 0 0 2 2",
	}
	for _, l := range lines[:len(lines)-1] {
		g.Expect(program.Snapshot()).To(Equal(l))
		g.Expect(program.Step()).To(BeFalse())
	}
	g.Expect(program.Snapshot()).To(Equal(lines[len(lines)-1]))
	g.Expect(program.Step()).To(BeTrue())

	g.Expect(program.Step()).To(BeTrue())
}
