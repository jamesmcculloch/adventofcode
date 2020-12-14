package main

import (
	"fmt"
	"strings"

	"github.com/jamesmcculloch/adventofcode/utils"
)

type seaPortComputer struct {
	memory   map[int64]int64
	mask     string
	version2 bool
}

func new(version2 bool) *seaPortComputer {
	return &seaPortComputer{
		memory:   map[int64]int64{},
		version2: version2,
	}
}

func (c *seaPortComputer) loadProgram(program []string) {
	for _, instruction := range program {
		if strings.HasPrefix(instruction, "mask") {
			c.loadMask(instruction)
		}

		if strings.HasPrefix(instruction, "mem") {
			c.writeMemory(instruction)
		}
	}
}

func (c *seaPortComputer) loadMask(maskInstruction string) {
	var mask string

	fmt.Sscanf(maskInstruction, "mask = %s", &mask)

	c.mask = mask
}

func (c *seaPortComputer) writeMemory(memoryInstruction string) {
	var memoryAddress int64
	var value int64

	fmt.Sscanf(memoryInstruction, "mem[%d] = %d", &memoryAddress, &value)

	if !c.version2 {
		value = c.applyMaskToValue(value)
		c.memory[memoryAddress] = value
	} else {
		memoryAddresses := c.applyMaskToMemory(memoryAddress)
		for _, memoryAddress := range memoryAddresses {
			c.memory[memoryAddress] = value
		}
	}
}

func (c *seaPortComputer) applyMaskToValue(value int64) int64 {
	for i := 0; i < len(c.mask); i++ {
		if c.mask[i] == '1' {
			value = c.setBit(value, int64(len(c.mask)-1-i))
		}
		if c.mask[i] == '0' {
			value = c.clearBit(value, int64(len(c.mask)-1-i))
		}
	}

	return value
}

func (c seaPortComputer) setBit(value int64, pos int64) int64 {
	mask := int64((1 << pos))
	value |= mask
	return value
}

func (c *seaPortComputer) clearBit(value int64, pos int64) int64 {
	mask := int64(^(1 << pos))
	value &= mask
	return value
}

func (c *seaPortComputer) applyMaskToMemory(memoryAddress int64) []int64 {
	addresses := []int64{}

	for i := 0; i < len(c.mask); i++ {
		if c.mask[i] == '1' {
			memoryAddress = c.setBit(memoryAddress, int64(len(c.mask)-1-i))
		}
	}

	addresses = append(addresses, c.applyFloatingBits(memoryAddress, 0)...)

	return addresses
}

func (c *seaPortComputer) applyFloatingBits(memoryAddress int64, currentOffset int) []int64 {
	addresses := []int64{memoryAddress}

	for i := currentOffset; i < len(c.mask); i++ {
		if c.mask[i] == 'X' {
			onesAddress := c.setBit(memoryAddress, int64(len(c.mask)-1-i))
			addresses = append(addresses, c.applyFloatingBits(onesAddress, i+1)...)
			zerosAddress := c.clearBit(memoryAddress, int64(len(c.mask)-1-i))
			addresses = append(addresses, c.applyFloatingBits(zerosAddress, i+1)...)
		}
	}

	return addresses
}

func (c *seaPortComputer) sumMemory() int64 {
	var total int64
	for _, value := range c.memory {
		total += value
	}
	return total
}

func main() {
	initialisationProgram, err := utils.LoadStringsFromFile("input")
	if err != nil {
		panic(err)
	}

	isVerison2 := false
	computer := new(isVerison2)

	computer.loadProgram(initialisationProgram)

	fmt.Printf("part 1: %d\n", computer.sumMemory())

	isVerison2 = true
	computer = new(isVerison2)

	computer.loadProgram(initialisationProgram)

	fmt.Printf("part 2: %d\n", computer.sumMemory())
}
