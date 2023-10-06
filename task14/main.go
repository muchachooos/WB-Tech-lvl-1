package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

func main() {
	// Создаём переменную i типа interface{}.
	var i interface{}

	// Присваеваем ей разные значения и передаём в ф-цию для определения типа.
	i = 33
	typeDefinition(i)

	i = "III"
	typeDefinition(i)

	i = false
	typeDefinition(i)

	i = make(chan int64)
	typeDefinition(i)
}

// typeDefinition спользует пакет reflect для определения типа переменной in.
func typeDefinition(in interface{}) {
	fmt.Println(reflect.TypeOf(in))
}
