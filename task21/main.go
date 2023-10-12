package main

// Реализовать паттерн «адаптер» на любом примере.

func main() {
	// Создаём объект типа SysAdmin.
	oleg := SysAdmin{}

	// Вызов функции runComputer.
	runComputer(oleg)
	// Вызов функции runReactor.
	runReactor(oleg)
}

func runComputer(oleg SysAdmin) {
	// Создаем объект Mac.
	mac := &Mac{}

	// Вызываем методы TurnOn и TurnOff у объекта SysAdmin с объектом Mac.
	oleg.TurnOn(mac)
	oleg.TurnOff(mac)
}

func runReactor(oleg SysAdmin) {
	// Создаем объект NuclearReactor.
	reactor := &NuclearReactor{true}
	// Создаем адаптер ReactorAdapter для этого объекта.
	reactorAdapter := ReactorAdapter{reactor}

	// Вызываем методы TurnOn и TurnOff у объекта SysAdmin с созданным адаптером ReactorAdapter.
	oleg.TurnOn(&reactorAdapter)
	oleg.TurnOff(&reactorAdapter)
}
