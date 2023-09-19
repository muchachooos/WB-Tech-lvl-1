package main

import "fmt"

// Дана последовательность температурных колебаний: -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func main() {
	// Определяем последовательность температурных колебаний.
	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 1, 6, -2}

	fmt.Println(distribution(temperatures))
}

func distribution(temperatures []float32) map[int][]float32 {
	// Создаем map для группировки значений.
	tempGroup := make(map[int][]float32)

	// Проходим по каждому значению в последовательности.
	for _, valTemp := range temperatures {
		valTempInt := int(valTemp) // Конвертируем в инт.
		rem := valTempInt % 10     // Получаем остаток от деления.
		key := valTempInt - rem    // Вычитаем остаток от значения тем самым определяем множество.

		// Добавляем значение в соответствующее множество.
		tempGroup[key] = append(tempGroup[key], valTemp)

	}

	return tempGroup
}
