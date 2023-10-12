package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

// Human создаём структуру.
type Human struct {
	Name      string
	HairColor string
}

// Action структура которая включает в себя объект типа Human.
type Action struct {
	Human
	Type  string
	Count int
}

func main() {
	// Создаём объект action типа Action который содержит объект типа Human, поля Type и Count.
	action := Action{
		Human: Human{
			Name:      "Oleg",
			HairColor: "Pink",
		},
		Type:  "Jump",
		Count: 4,
	}

	// Вызываем метод UpAge для объекта типа Human внутри объекта типа Action.
	action.ChangeColor("Red")
	// Вызываем метод Do для объекта типа Action.
	action.Do()
}

// ChangeColor Метод для объекта типа Human, который изменяет ему цвет волос.
func (h *Human) ChangeColor(hairColor string) {
	// Меняем значение поля HairColor на переданное и выводим его.
	h.HairColor = hairColor

	fmt.Println("New color: ", h.HairColor)
}

// Do метод для совершения действия Type объектом типа Action заданное количество раз Count.
func (a *Action) Do() {
	for i := 0; i < a.Count; i++ {
		fmt.Println("do " + a.Type + " by " + a.Human.Name)
	}
}
