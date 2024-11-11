# Lesson 6 
[Tasks](https://cloud-f.edupage.org/cloud?z%3AmgcytW6YkLEm7BQxSfipNOnSvZ2eO8i5uTwatxFFeP2xOzLTu3T93x1KCPOxrbnz)

[Live](https://go.dev/play/p/LbLYCacRNJt)

# Code

```Golang
package main

import "fmt"

func main() {
	task1()
	task2()
	task3()
	task4()
	task5()
	task6()
	task7()
	task8()
	task9()
	task10()
}

// * 1. Факториал числа

func task1() {
	fmt.Println("Задание 1")
	defer fmt.Println("Конец Задания 1")
	fmt.Println(factorial(5))
}

func factorial(a int) int {
	if a <= 1 {
		return 1
	}
	return a * factorial(a-1)
}

// * 2. Числа Фибоначчи

func task2() {
	fmt.Println("Задание 2")
	defer fmt.Println("Конец Задания 2")
	fmt.Println(fibonacci(6))
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// * 3. Сумма чисел от 1 до n

func task3() {
	fmt.Println("Задание 3")
	defer fmt.Println("Конец Задания 3")
	fmt.Println(summ(10))
}
func summ(n int) int {
	if n <= 1 {
		return n
	}
	return n + summ(n-1)
}

// * 4. Возведение числа в степень

func task4() {
	fmt.Println("Задание 4")
	defer fmt.Println("Конец Задания 4")
	fmt.Println(pow(2, 3))
}

func pow(a int, n int) int {
	if n <= 1 {
		return a
	}
	return a * pow(a, n-1)
}

// * 5. Палиндром

func task5() {
	fmt.Println("Задание 5")
	defer fmt.Println("Конец Задания 5")
	fmt.Println(palindrome("level"))
}

func palindrome(current string) bool {
	if len(current) <= 1 {
		return true
	}
	if current[0] != current[len(current)-1] {
		return false
	}
	return palindrome(current[1 : len(current)-1])
}

// * 6. Нахождение наибольшего элемента в массиве

func task6() {
	fmt.Println("Задание 6")
	defer fmt.Println("Конец Задания 6")
	arr := []int{1, 4, 2, 9, 3}
	fmt.Println("Наибольший элемент массива:", max(arr))
}

func max(pointer []int) (max_el int) {
	max_el = pointer[0]
	for _, el := range pointer {
		if el > max_el {
			max_el = el
		}
	}
	return max_el
}

// * 7. Нахождение суммы элементов массива

func task7() {
	fmt.Println("Задание 7")
	defer fmt.Println("Конец Задания 7")
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Сумма всех элементов массива:", sum(arr))
}

func sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + sum(arr[1:])
}

// * 8. Реверс строки

func task8() {
	fmt.Println("Задание 8")
	defer fmt.Println("Конец Задания 8")
	fmt.Println(reverse("hello"))
}

func reverse(last string) (new string) {
	if len(last) == 0 {
		return ""
	}
	return string(last[len(last)-1]) + reverse(last[:len(last)-1])
}

// * 9. Поиск элемента в массиве

func task9() {
	fmt.Println("Задание 9")
	defer fmt.Println("Конец Задания 9")
	arr := []int{1, 3, 5, 7}
	target := 5
	fmt.Println("Элемент", target, "есть в массиве:", find(arr, target))
}

func find(array []int, target int) bool {
	if len(array) == 0 {
		return false
	}
	if array[len(array)-1] == target {
		return true
	}
	return find(array[:len(array)-1], target)
}

// * 10. Число комбинаций (биномиальные коэффициенты)

func task10() {
	fmt.Println("Задание 10")
	defer fmt.Println("Конец Задания 10")
	n := 5
	k := 2
	fmt.Println("C(", n, ",", k, ") =", combinations(n, k))
}

func combinations(n, k int) (c int) {
	if k == 0 || k == n {
		return 1
	}
	return combinations(n-1, k-1) + combinations(n-1, k)
}

```
