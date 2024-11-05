package main

import "fmt"

func main() {
	task1([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	task2([]int{1, 2, 3, 4, 55, 6})
	task3([]int{1, 2, 3, 4, 55, 6})
	task4([]int{1, 2, 3, 4, 55, 909, 198}, 55)
	task5([]int{3, 1, 13, 546, 12, 321, 1, 2, 4})
	task6([]int{3, 1, 13, 546, 12, 321, 1, 2, 4})
	task7([]int{3, 1, 13, 546, 12, 321, 1, 2, 4}, 1)
	task8([]int{3, 1, 13, 546, 12, 321, 1, 2, 4})
	task9([]int{1, 2, 3, 4})
	task10([]int{1, 2, 3, 4})
	task11([]int{1, 2, 3, 4})
	task12([]int{1, 2, 3, 4})
	task13([]int{1, 0, 0, 0, 0, 0, 0, 0, 2, 3, 4})
	task14([]int{3, 1, 13, 546, 12, 321, 1, 2, 4})
	task15([]int{3, 1, 13, 546, 12, 321, 1, 2, 4})
	task16([]int{3, 1, 13, 546, 12, 321, 1, 2, 4})
	task17([]int{3, 1, 13, 546, 12, 321, 1, 2, 4, -1, -2, -3, -4, -5, -6, -7, -8, -9, -19})
	task18([]int{3, 1, 13, 546, 12, 321, 1, 2, 4, -1, -2, -3, -4, -5, -6, -7, -8, -9, -19})
	task19([]int{3, 1, 13, 546, 12, 321, 1, 2, 4})
}

// * 1. Найти сумму элементов массива

func task1(arr []int) {
	fmt.Println("Задача 1")
	sum := 0
	for _, num := range arr {
		sum += num
	}
	fmt.Println("Сумма всех элементов массива:", sum)
	fmt.Println("Конец задачи 1")
}

// * 2. Найти максимальный элемент в массиве

func task2(arr []int) {
	fmt.Println("Задача 2")
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	fmt.Println("Максимальный элемент массива:", max)
	fmt.Println("Конец задачи 2")
}

// * 3. Найти минимальный элемент в массиве

func task3(arr []int) {
	fmt.Println("Задача 3")
	min := arr[0]
	for _, num := range arr {
		if num < min {
			min = num
		}
	}
	fmt.Println("Минимальный элемент массива:", min)
	fmt.Println("Конец задачи 3")
}

// * 4. Найти индекс заданного элемента в массиве (либо -1, если не найден)

func task4(arr []int, indx int) {
	fmt.Println("Задача 4")
	defer fmt.Println("Конец задачи 4")
	for i, num := range arr {
		if num == indx {
			fmt.Println("Индекс заданного элемента:", i)
			return
		}
	}
	fmt.Println("Индекс заданного элемента не найден")
}

// * 5. Отсортировать массив по возрастанию

func task5(arr []int) {
	fmt.Println("Задача 5")
	defer fmt.Println("Конец задачи 5")
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println("Отсортированный массив по возрастанию:", arr)
}

// * 6. Отсортировать массив по убыванию

func task6(arr []int) {
	fmt.Println("Задача 6")
	defer fmt.Println("Конец задачи 6")
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println("Отсортированный массив по убыванию:", arr)
}

// * 7. Найти количество вхождений заданного элемента в массиве

func task7(arr []int, el int) {
	fmt.Println("Задача 7")
	defer fmt.Println("Конец задачи 7")
	count := 0
	for _, num := range arr {
		if num == el {
			count++
		}
	}
	fmt.Println("Количество вхождений заданного элемента в массиве:", count)
}

// * 8. Проверить, есть ли в массиве повторяющиеся элементы

func task8(arr []int) {
	fmt.Println("Задача 8")
	defer fmt.Println("Конец задачи 8")
	count := make(map[int]int)
	for _, num := range arr {
		count[num]++
	}
	hasDuplicates := false
	for _, v := range count {
		if v > 1 {
			hasDuplicates = true
			break
		}
	}
	fmt.Println("Есть ли повторяющиеся элементы в массиве:", hasDuplicates)

}

// * 9. Поменять местами первый и последний элемент массива

func task9(arr []int) {
	fmt.Println("Задача 9")
	defer fmt.Println("Конец задачи 9")
	if len(arr) < 2 {
		fmt.Println("Массив должен содержать хотя бы два элемента")
		return
	}
	arr[0], arr[len(arr)-1] = arr[len(arr)-1], arr[0]
	fmt.Println("Массив после смены первого и последнего элементов:", arr)
}

// * 10. Сдвинуть все элементы массива на одну позицию вправо

func task10(arr []int) {
	fmt.Println("Задача 10")
	defer fmt.Println("Конец задачи 10")
	if len(arr) < 2 {
		fmt.Println("Массив после сдвига всех элементов на одну позицию вправо:", arr)
		return
	}
	last := arr[len(arr)-1]
	for i := len(arr) - 1; i > 0; i-- {
		arr[i] = arr[i-1]
	}
	arr[0] = last
	fmt.Println("Массив после сдвига всех элементов на одну позицию вправо:", arr)
}

// * 11. Сдвинуть все элементы массива на одну позицию влево

func task11(arr []int) {
	fmt.Println("Задача 11")
	defer fmt.Println("Конец задачи 11")
	if len(arr) < 2 {
		fmt.Println("Массив после сдвига всех элементов на одну позицию влево:", arr)
		return
	}
	first := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = first
	fmt.Println("Массив после сдвига всех элементов на одну позицию влево:", arr)
}

// * 12. Инвертировать (развернуть) массив

func task12(arr []int) {
	fmt.Println("Задача 12")
	defer fmt.Println("Конец задачи 12")
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println("Инвертированный массив:", arr)
}

// * 13. Удалить все нулевые элементы из массива

func task13(arr []int) {
	fmt.Println("Задача 13")
	defer fmt.Println("Конец задачи 13")
	var newArr []int
	for _, num := range arr {
		if num != 0 {
			newArr = append(newArr, num)
		}
	}
	fmt.Println("Массив после удаления нулевых элементов:", newArr)
}

// * 14. Подсчитать количество четных и нечетных элементов в массиве

func task14(arr []int) {
	fmt.Println("Задача 14")
	defer fmt.Println("Конец задачи 14")
	countEven := 0
	countOdd := 0
	for _, num := range arr {
		if num%2 == 0 {
			countEven++
		} else {
			countOdd++
		}
	}
	fmt.Println("Количество четных элементов:", countEven)
	fmt.Println("Количество нечетных элементов:", countOdd)
}

// * 15. Суммировать элементы на четных индексах

func task15(arr []int) {
	fmt.Println("Задача 15")
	defer fmt.Println("Конец задачи 15")
	sum := 0
	for i := 0; i < len(arr); i += 2 {
		sum += arr[i]
	}
	fmt.Println("Сумма элементов на четных индексах:", sum)
}

// * 16. Суммировать элементы на нечетных индексах

func task16(arr []int) {
	fmt.Println("Задача 16")
	defer fmt.Println("Конец задачи 16")
	sum := 0
	for i := 1; i < len(arr); i += 2 {
		sum += arr[i]
	}
	fmt.Println("Сумма элементов на нечетных индексах:", sum)
}

// * 17. Создать новый массив, содержащий только положительные элементы исходного массива

func task17(arr []int) {
	fmt.Println("Задача 17")
	defer fmt.Println("Конец задачи 17")
	newArr := make([]int, 0, len(arr))
	for _, num := range arr {
		if num > 0 {
			newArr = append(newArr, num)
		}
	}
	fmt.Println("Массив с положительными элементами:", newArr)
}

// * 18. Создать новый массив, содержащий только отрицательные элементы исходного массива

func task18(arr []int) {
	fmt.Println("Задача 18")
	defer fmt.Println("Конец задачи 18")
	newArr := make([]int, 0, len(arr))
	for _, num := range arr {
		if num < 0 {
			newArr = append(newArr, num)
		}
	}
	fmt.Println("Массив с отрицательными элементами:", newArr)
}

// * 19. Найти второй по величине элемент в массиве

func task19(arr []int) {
	fmt.Println("Задача 19")
	defer fmt.Println("Конец задачи 19")
	if len(arr) < 2 {
		fmt.Println("Массив должен содержать хотя бы два элемента")
		return
	}
	max1, max2 := arr[0], arr[1]
	for _, num := range arr {
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 && num != max1 {
			max2 = num
		}
	}
	fmt.Println("Второй по величине элемент в массиве:", max2)
}
