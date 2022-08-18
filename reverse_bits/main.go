package main

import (
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

func reverseBits(num uint32) uint32 {
	return bits.Reverse32(num)
}

func main() {

	given, err := strconv.ParseUint("00000010100101000001111010011100", 2, 32)
	expected, _ := strconv.ParseUint("00111001011110000010100101000000", 2, 32)

	if err != nil {
		fmt.Println(err)
		fmt.Println(fmt.Errorf("got error "))
		os.Exit(1)
	}

	result := reverseBits(uint32(given))

	fmt.Printf("Given: %d\nResult: %d\nExpected: %d\n", given, result, expected)
}
