package main

// SysAdmin Тип клиента.
type SysAdmin struct {
}

// Computer интерфейс с методами TurnOn и TurnOff.
type Computer interface {
	TurnOn()
	TurnOff()
}

// TurnOn Метод SysAdmin, который вызывает метод TurnOn у переданного объекта типа Computer.
func (s *SysAdmin) TurnOn(c Computer) {
	c.TurnOn()
}

// TurnOff Метод SysAdmin, который вызывает метод TurnOn у переданного объекта типа Computer.
func (s *SysAdmin) TurnOff(c Computer) {
	c.TurnOff()
}
