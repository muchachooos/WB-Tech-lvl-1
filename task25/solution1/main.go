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

func mySleep(d time.Duration) {
	// time.After ожидает истечения продолжительности d, а затем отправляет текущее время по возвращаемому каналу.
	<-time.After(d)
}
