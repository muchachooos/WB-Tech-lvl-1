package main

// ReactorAdapter Структура, которая является адаптером между интерфейсом Computer и типом NuclearReactor.
type ReactorAdapter struct {
	*NuclearReactor
}

// TurnOn вызывает метод Power для вложенного объекта NuclearReactor и передаёт значение true.
func (r *ReactorAdapter) TurnOn() {
	r.Power(true)
}

// TurnOff вызывает метод Power для вложенного объекта NuclearReactor и передаёт значение false.
func (r *ReactorAdapter) TurnOff() {
	r.Power(false)
}
