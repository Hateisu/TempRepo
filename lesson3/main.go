package main

import (
	"fmt"
	"strings"
)

func main() {
	task1([]int{5, 3, 8, 6, 2})
	task2([]int{1, 2, 3, 4, 5})
	task3([]int{7, 1, 3, 5, 9})
	task4("algorithm")
	task5(56, 98)
}

// * Задание 1: Сортировка с подсчётом
// * Описание:
// * Напишите программу, которая сортирует массив чисел с помощью пузырьковой
// * сортировки (bubble sort), но добавляет счётчик, который отслеживает количество
// * перестановок элементов. Вспомогательный алгоритм должен возвращать количество
// * перестановок.

func task1(arr []int) int {
	fmt.Println("Задание 1")
	defer fmt.Println("Конец задания 1")
	counter := bubblesort(arr)
	fmt.Println("Итого перестановок ", counter)
	return counter
}

func bubblesort(arr []int) int {
	var counter int = -1
	fmt.Println("ВВодной массив ", arr)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				counter++ // ! СЧЕТЧИК ПЕРЕСТАНОВКИ
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println("Итоговый массив ", arr)
	return counter
}

// * Задание 2: Рекурсивный поиск суммы массива
// * Описание:
// * Реализуйте программу для поиска суммы всех элементов в массиве с использованием
// * рекурсии. Вспомогательный алгоритм должен реализовать рекурсивное сложение.

func task2(arr []int) {
	fmt.Println("Задание 2")
	defer fmt.Println("Конец задания 2")
	fmt.Println("Сумма ", task2_recursion(arr))
}

func task2_recursion(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + task2_recursion(arr[1:]) // ! РЕКУРСИЯ
}

// * Задание 3: Поиск среднего арифметического с исключением максимума
// * Описание:
// * Напишите программу, которая находит среднее арифметическое всех элементов
// * массива, кроме максимального. Вспомогательный алгоритм должен находить
// * максимальное значение.

func task3(arr []int) {
	fmt.Println("Задание 3")
	defer fmt.Println("Конец задания 3")

	_, v := task3_maxer(arr)
	fmt.Println("Максимальное число: ", v)

	var sum, popups int
	for _, num := range arr {
		if num != v {
			sum += num
		} else {
			popups++
		}
	}
	fmt.Println("Среднее арифметическое ", (float64(sum) / float64(len(arr)-popups)))
}

func task3_maxer(arr []int) (index, maximum int) {
	for v, num := range arr {
		if num > maximum {
			maximum = num
			index = v
		}
	}
	return index, maximum
}

// * Задание 4: Переворот строки с использованием стека
// * Описание:
// * Напишите программу, которая переворачивает строку с использованием стека.
// * Реализуйте вспомогательный алгоритм для работы со стеком: добавление элементов в
// * стек и извлечение их в обратном порядке.

func task4(line string) {
	fmt.Println("Задание 4")
	defer fmt.Println("Конец задания 4")
	arr := strings.Split(line, "")
	fmt.Println("Первая строка: ", strings.Join(arr, ""))
	stack := make([]string, 0, len(arr))
	iterator := len(arr)
	for i := 0; i < iterator; i++ {
		stack = task4_add(stack, task4_getLast(&arr))
	}
	fmt.Println("Перевернутая строка: ", strings.Join(stack, ""))

}
func task4_getLast(arr *[]string) string {
	tmp := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	return tmp
}
func task4_add(stack []string, el string) []string {
	stack = append(stack, el)
	return stack
}

// * Задание 5: НОД двух чисел с использованием алгоритма Евклида
// * Описание:
// * Напишите программу для нахождения наибольшего общего делителя (НОД) двух чисел
// * с использованием алгоритма Евклида как вспомогательного алгоритма.

func task5(a, b int) {
	fmt.Println("Задание 5")
	defer fmt.Println("Конец задания 5")
	fmt.Println("НОД", a, " и ", b, ": ", task5_gcd(a, b))
}

func task5_gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return task5_gcd(b, a%b)
	}
}
