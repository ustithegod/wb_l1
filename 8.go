package main

func SetBit(n int64, i uint, bit int) int64 {
	if bit == 1 {
		return n | (1 << i)
	}
	return n &^ (1 << i)
}
