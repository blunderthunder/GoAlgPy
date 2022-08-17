package main

func reverseBits(num uint64) uint64 {

	var rev uint64 = 0

	for num > 0 {
		rev = rev << 1

		if (num & 1) == 1 {
			rev = rev ^ 1
		}

		num = num >> 1
	}
	return rev
}

func main() {
	// var given uint64 = strconv.ParseUint("00000010100101000001111010011100")
	// var expected uint64 = 00111001011110000010100101000000
	// fmt.Println("given: {} , expected result : {} \n\t result : {}", given, expected, reverseBits(given))
}
