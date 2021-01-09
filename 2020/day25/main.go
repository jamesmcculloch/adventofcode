package main

import (
	"fmt"

	"github.com/jamesmcculloch/adventofcode/utils"
)

func encryptionKey(cardPublicKey, doorPublicKey int) int {
	cardLoopSize := loopSize(cardPublicKey)
	doorLoopSize := loopSize(doorPublicKey)

	cardEncryptionKey := encryptionKeyValue(doorLoopSize, cardPublicKey)
	doorEncryptionKey := encryptionKeyValue(cardLoopSize, doorPublicKey)
	if cardEncryptionKey != doorEncryptionKey {
		return -1
	}
	return cardEncryptionKey
}

func loopSize(key int) int {
	subjectNumber := 7
	size := 0
	value := 1
	for {
		size++
		value = value * subjectNumber
		value = value % 20201227

		if value == key {
			break
		}
	}
	return size
}

func encryptionKeyValue(loopSize, subjectNumber int) int {
	value := 1
	for i := 1; i <= loopSize; i++ {
		value = value * subjectNumber
		value = value % 20201227
	}
	return value
}

func main() {
	keys, err := utils.LoadNumbersFromInput("input")
	if err != nil {
		panic(err)
	}

	fmt.Printf("part 1: %d\n", encryptionKey(keys[0], keys[1]))
}
