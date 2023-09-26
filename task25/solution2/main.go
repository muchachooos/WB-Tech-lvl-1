package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func main() {
	fmt.Println("Start")

	mySleep(3 * time.Second)

	fmt.Println("End")
}

func mySleep(in time.Duration) {
	// Засекаем текущее время.
	startTime := time.Now()

	// Since возвращает время, прошедшее с момента startTime, пока оно меньше in, цикл не завершиться.
	for time.Since(startTime) < in {
	}
}
