package main

import "fmt"

// Дана переменная int64.
// Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	var num int64 = 123 // Исходное число. (1111011)
	bitIndex := 2       // Индекс бита который мы хотим поменять.

	// Вызываем функцию setBit().
	num = setBit(num, uint64(bitIndex))
	fmt.Println(num) // 127 (1111111)

	// Вызываем функцию clearBit().
	num = clearBit(num, uint64(bitIndex))
	fmt.Println(num) // 123 (1111011)
}

// Устанавливает бит в позиции pos в числе n.
func setBit(num int64, pos uint64) int64 {
	var mask int64 = 1 << pos
	// 0001 << 3 = 1000
	// 1001 | 0010 = 1011
	// 1001 | 0100 = 1101
	return num | mask
}

// Очищает бит в позиции pos в числе n.
func clearBit(num int64, pos uint64) int64 {
	var mask int64 = ^(1 << pos)
	return num & mask
}
