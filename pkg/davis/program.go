package davis

import (
	"fmt"
	"strconv"
	"strings"
)

func NewProgram(instructions []Instruction, inputs []int) (Program, error) {
	labels := make(map[Label]int)

	maxVar := 0
	for index, instruction := range instructions {
		if int(instruction.Var) > maxVar {
			maxVar = int(instruction.Var)
		}

		if instruction.Label == NoLabel {
			continue
		}

		previousIndex, ok := labels[instruction.Label]
		if ok {
			// if you want to support a single label on multiple lines, comment out this section
			return Program{}, fmt.Errorf("label %s is referrenced twice at instructions %d and %d",
				instruction.Label.String(), previousIndex, index)
		}
		labels[instruction.Label] = index

	}
	variables := make([]int, max(maxVar, len(inputs)*2, 1))

	for inputIndex, value := range inputs {
		variables[2*(inputIndex+1)-1] = value
	}

	return Program{
		instructions:            instructions,
		labels:                  labels,
		variables:               variables,
		currentInstructionIndex: 0,
	}, nil
}

type Program struct {
	instructions []Instruction
	// labels contains a map from Label to the number of its instruction
	labels map[Label]int
	// variables contains a map from each variable to its value
	variables               []int
	currentInstructionIndex int
}

func (p *Program) String() string {
	fmt.Println("Instructions")
	for i, instruction := range p.instructions {
		fmt.Println(i, instruction.String())
	}
	fmt.Println("Labels")
	for label, instructionNumber := range p.labels {
		fmt.Println(label.String(), "->", instructionNumber)
	}

	fmt.Println("variables")
	for index, value := range p.variables {
		fmt.Println(Variable(index+1).String(), "->", value)
	}
	return ""
}

func (p *Program) Snapshot() string {
	var builder, inputs, variables strings.Builder
	builder.WriteString(fmt.Sprintf("%d ", p.currentInstructionIndex+1))
	for i, value := range p.variables[1:] {
		switch i % 2 {
		case 0:
			inputs.WriteString(fmt.Sprintf("%d ", value))
		case 1:
			variables.WriteString(fmt.Sprintf("%d ", value))
		}
	}
	builder.WriteString(inputs.String())
	builder.WriteString(variables.String())
	builder.WriteString(strconv.Itoa(p.variables[0]))
	return builder.String()
}

// Step runs one step of the program
func (p *Program) Step() (terminated bool) {
	if p.IsTerminated() {
		return true
	}
	currentInstruction := p.instructions[p.currentInstructionIndex]
	switch currentInstruction.Type {
	case InstructionTypeNoOp:
		p.currentInstructionIndex += 1
	case InstructionTypeIncr:
		p.variables[currentInstruction.Var-1] += 1
		p.currentInstructionIndex += 1
	case InstructionTypeDecr:
		p.variables[currentInstruction.Var-1] -= 1
		p.currentInstructionIndex += 1
	default: // IF V != 0 GOTO L instruction
		value := p.variables[currentInstruction.Var-1]
		if value == 0 {
			p.currentInstructionIndex += 1
			return p.IsTerminated()
		}
		instructionIndex, ok := p.labels[Label(currentInstruction.Type-2)]
		if !ok {
			// label doesn't exist, terminating
			p.currentInstructionIndex = len(p.instructions)
			return true
		}
		p.currentInstructionIndex = instructionIndex
	}
	return p.IsTerminated()
}

func (p *Program) IsTerminated() bool {
	return p.currentInstructionIndex >= len(p.instructions)
}
