package main

import (
	"fmt"
	"math"
)

func main() {
	task1([]int{5, 3, 8, 6, 2})
	task2(5)
	task3(5)
	task4(150)
	fmt.Println("Ответ от функции 5 ", task5())
	task6(10)
	task7(21)
	task8([]int{1, 2, 3, 4, 5, 6, 199, 7, 8, 9, 10})
	task9(21)
	task10(212121)
	task11(11)
	task11withArray([]int{1, 2, 3, 4, 5}, 11)
	task12([]int{1, 2, 3, 4, 5, 6, 199, 7, 8, 9, 10})
	task13("niggin1")
	task14([]int{1, 3, 5, 7, 8, 9})
	task15([]int{1, 2, 3, 4, 5, 6, 199, 7, 8, 9, 10})
	task16(10, 12)
	task17([]int{1, 2, 3, 4, 5, 6, 199, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 198, 7, 8, 9, 10})
	task18(101)
	task19(10, 12)
	task20(10, 12)
	task21([]int{1, 2, 3}, []int{9, 10})
	task22(6, 5)
	task23([]int{1, 2, 3}, []int{9, 10})
	task24(6, 5)
	task25([]int{1, 2, 3}, []int{9, 10})
	task26([]int{1, 2, 3}, []int{9, 10})
	task27("abcd", "qqwesd")
	task28("abcd", "qqwesd")
	task29([]int{1, 2, 3}, []int{7, 9, 10})
	task30([]int{1, 2, 3}, []int{9, 10})

}

// * Задачи O(1):

// * 1. Напишите функцию, которая возвращает первый элемент в списке чисел.

func task1(arr []int) {
	fmt.Println("Задание 1 O(1)")
	defer fmt.Println("Конец Задания 1")
	fmt.Println("Первый элемент списка:", arr[0])
}

// * 2. Функция, которая проверяет, является ли число четным или нечетным.

func task2(a int) {
	fmt.Println("Задание 2 O(1)")
	defer fmt.Println("Конец Задания 2")
	if a%2 == 0 {
		fmt.Println(a, "четное число")
	} else {
		fmt.Println(a, "нечётное число")
	}
}

// * 3. Найдите квадрат числа.

func task3(a int) {
	fmt.Println("Задание 3 O(1)")
	defer fmt.Println("Конец Задания 3")
	fmt.Println("Квадрат числа", a, "равно", a*a)
}

// * 4. Проверьте, больше ли данное число 100.

func task4(a int) {
	fmt.Println("Задание 4 O(1)")
	defer fmt.Println("Конец Задания 4")
	if a > 100 {
		fmt.Println(a, "больше 100")
	} else {
		fmt.Println(a, "меньше или ра��но 100")
	}
}

// * 5. Напишите функцию, которая всегда возвращает число 5.

func task5() int {
	fmt.Println("Задание 5 O(1)")
	defer fmt.Println("Конец Задания 5")
	return 5
}

// * Задачи O(N):

// * 6. Определите сумму всех чисел от 1 до N.

func task6(a int) {
	fmt.Println("Задание 6 O(N)")
	defer fmt.Println("Конец Задания 6")
	sum := 0
	for i := 1; i <= a; i++ {
		sum += i
	}
	fmt.Println("Сумма чисел от 1 до", a, "равна", sum)
}

// * 7. Проверьте, является ли число простым.

func task7(a int) {
	fmt.Println("Задание 7 O(N)")
	defer fmt.Println("Конец Задания 7")
	if a <= 1 {
		fmt.Println(a, "не является простым числом")
		return
	}
	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			fmt.Println(a, "не является простым числом")
			return
		}
	}
	fmt.Println(a, "является простым числом")
}

// * 8. Найдите наибольшее число среди N чисел.

func task8(N []int) {
	fmt.Println("Задание 8 O(N)")
	defer fmt.Println("Конец Задания 8")
	max := N[0]
	for _, v := range N {
		if v > max {
			max = v
		}
	}
	fmt.Println("Наибольшее число:", max)
}

// * 9. Подсчитайте количество четных чисел в последовательности чисел от 1 до N.

func task9(N int) {
	fmt.Println("Задание 9 O(N)")
	defer fmt.Println("Конец Задания 9")
	count := 0
	for i := 1; i <= N; i++ {
		if i%2 == 0 {
			count++
		}
	}
	fmt.Println("Количество четных чисел:", count)
}

// * 10. Найдите количество цифр в числе.

func task10(chislo int) {
	fmt.Println("Задание 10 O(N)")
	defer fmt.Println("Конец Задания 10")
	count := 0
	for chislo > 0 {
		chislo /= 10
		count++
	}
	fmt.Println("Количество цифр в числе:", count)
}

// * Задачи O(N^2):

// * 11. Напишите функцию для проверки, есть ли два числа, сумма которых равназаданному числу, перебирая все пары чисел.

func task11(a int) {
	fmt.Println("Задание 11 (Все числа) O(N^2)")
	defer fmt.Println("Конец Задания 11")
	for i := 0; i <= a; i++ {
		for j := i + 1; j <= a; j++ {
			if i+j == a {
				fmt.Println("Сумма чисел", i, "и", j, "равна", a)
			}
		}
	}
}

func task11withArray(arr []int, a int) {
	fmt.Println("Задание 11 (Передаем массив) O(N^2)")
	defer fmt.Println("Конец Задания 11")
	for _, i := range arr {
		for _, j := range arr {
			if i+j == a {
				fmt.Println("Сумма чисел", i, "и", j, "равна", a)
			}
		}
	}
}

// * 12. Реализуйте алгоритм пузырьковой сортировки для чисел.

func task12(arr []int) {
	fmt.Println("Задание 12 O(N^2)")
	defer fmt.Println("Конец Задания 12")
	fmt.Println("До сортировки ", arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("После сортировки ", arr)
}

// * 13. Проверьте, есть ли в строке палиндром (путем сравнения всех пар символов).

func task13(str string) {
	fmt.Println("Задание 13 O(N^2)")
	defer fmt.Println("Конец Задания 13")
	length := len(str)
	/* ! Решение в O(N)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			fmt.Println("Строка не является палиндромом")
			return
		}
	}*/
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if str[i] == str[j] && str[i] != str[length-j-1] {
				fmt.Println("Строка не является палиндромом")
				return
			}
		}
	}
	// В случае, если дошли до этого места, значит строчка является палиндромом
	// Но мы не знаем, является ли она палиндромом в O(N)
	// В таком случае, мы выводим сообщение о том, что строчка является палиндромом
	fmt.Println("Строка является палиндромом")
}

// * 14. Найдите все уникальные пары чисел, где произведение пары четное.

func task14(arr []int) {
	fmt.Println("Задание 14 O(N^2)")
	defer fmt.Println("Конец Задания 14")
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]*arr[j]%2 == 0 {
				fmt.Println("Пара чисел", arr[i], "и", arr[j], "где произведение пары четное")
			}
		}
	}
}

// * 15. Реализуйте алгоритм сортировки выбором.

func task15(arr []int) {
	fmt.Println("Задание 15 O(N^2)")
	defer fmt.Println("Конец Задания 15")
	fmt.Println("До сортировки ", arr)
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	fmt.Println("После сортировки ", arr)
}

// * Задачи O(N + K):

// * 16. Напишите функцию, которая считает сумму всех чисел от 1 до N и от 1 до K.

func task16(N, K int) {
	fmt.Println("Задание 16 O(N + K)")
	defer fmt.Println("Конец Задания 16")
	sumN := 0
	sumK := 0
	for i := 1; i <= N; i++ {
		sumN += i
	}
	for i := 1; i <= K; i++ {
		sumK += i
	}
	fmt.Println("Сумма чисел от 1 до", N, "и от 1 до", K, "равна", sumN+sumK)
}

// * 17. Определите наибольшее число среди чисел N и чисел K.

func task17(arrN, arrK []int) {
	fmt.Println("Задание 17 O(N + K)")
	defer fmt.Println("Конец Задания 17")
	maxN := 0
	maxK := 0
	for _, v := range arrN {
		if v > maxN {
			maxN = v
		}
	}
	for _, v := range arrK {
		if v > maxK {
			maxK = v
		}
	}
	fmt.Println("Наибольшее число из", arrN, "и", arrK, "равно", maxN, "и", maxK)
}

// * 18. Найдите количество цифр в числе и сумму всех чисел от 1 до K.

func task18(K int) {
	fmt.Println("Задание 18 O(N + K)")
	defer fmt.Println("Конец Задания 18")
	countDigits := 0
	sumK := 0
	// Подсчет суммы чисел от 1 до K
	for i := 1; i <= K; i++ {
		sumK += i
	}

	// Подсчет количества цифр
	Ktmp := K
	for true {
		if Ktmp%10 != 0 {
			countDigits++
		}
		Ktmp /= 10
		if Ktmp == 0 {
			countDigits++
			break
		}
	}
	fmt.Println("Количество цифр в числе от 1 до", K, "равно", countDigits)
	fmt.Println("Сумма чисел от 1 до", K, "равна", sumK)
}

// * 19. Напишите программу, которая выводит все простые числа до N и до K.

func task19(N, K int) {
	fmt.Println("Задание 19 ~O(N + K)~ O(N^2)")
	defer fmt.Println("Конец Задания 19")
	to_this_int := 0
	if N > K {
		to_this_int = N
	} else {
		to_this_int = K
	}
	for i := 0; i < to_this_int; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
func isPrime(a int) bool {
	if a <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

// * 20. Проверьте, содержат ли числа от 1 до N и от 1 до K хотя бы одно общее число.

func task20(N, K int) {
	fmt.Println("Задание 20 O(N + K)")
	defer fmt.Println("Конец Задания 20")
	var mappedN map[int]bool = make(map[int]bool)
	for i := 1; i < N; i++ {
		mappedN[i] = true
	}
	for i := 1; i < K; i++ {
		if mappedN[i] {
			fmt.Println("Общее число ", i)
			break
		}
	}
}

// * Задачи O(NK):

// * 21. Напишите функцию для нахождения всех пар чисел, где одно число из диапазона N, а другое из диапазона K.

func task21(N, K []int) {
	fmt.Println("Задание 21 O(NK)")
	defer fmt.Println("Конец Задания 21")
	for _, vN := range N {
		for _, vK := range K {
			fmt.Println(vN, "и", vK)
		}
	}
}

// * 22. Проверьте, есть ли среди чисел от 1 до N и от 1 до K такие, которые делятся друг на друга без остатка.

func task22(N, K int) {
	fmt.Println("Задание 22 O(NK)")
	defer fmt.Println("Конец Задания 22")
	for i := 1; i <= N; i++ {
		for j := 1; j <= K; j++ {
			if i == j || i%j == 0 || j%i == 0 {
				fmt.Println("Числа ", i, " и ", j, " делятся без остатка")
			}
		}
	}
}

// * 23. Найдите произведение всех пар чисел из диапазона N и K.

func task23(N, K []int) {
	fmt.Println("Задание 23 O(NK)")
	defer fmt.Println("Конец Задания 23")
	for _, i := range K {
		for _, j := range N {
			fmt.Println(i, " * ", j, " = ", (i * j))
		}
	}
}

// * 24. Реализуйте алгоритм поиска наибольшего общего делителя для всех пар чисел от 1 до N и от 1 до K.

func task24(N, K int) {
	fmt.Println("Задание 24 O(NK)")
	defer fmt.Println("Конец Задания 24")
	for i := 1; i <= N; i++ {
		for j := 1; j <= K; j++ {
			fmt.Println("НОД для ", i, " и ", j, " = ", gcd(i, j))
		}
	}
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

// * 25. Напишите функцию для нахождения всех комбинаций сумм чисел из диапазонов N и K.

func task25(N, K []int) {
	fmt.Println("Задание 25 O(NK)")
	defer fmt.Println("Конец Задания 25")
	for _, i := range K {
		for _, j := range N {
			fmt.Println(i, " + ", j, " = ", (i + j))
		}
	}
}

// * 26. Найдите все числа, которые являются кратными как для диапазона N, так и для диапазона K.

func task26(K, N []int) {
	fmt.Println("Задание 26 O(NK)")
	defer fmt.Println("Конец Задания 26")
	for _, k := range K {
		for _, n := range N {
			if (k != 0 && n != 0) && (k%n == 0 || n%k == 0) {
				fmt.Printf("%d кратно как %d, так и %d\n", k, k, n)
			}
		}
	}
}

// * 27. Реализуйте функцию, которая проверяет каждую строку длины N с каждой строкой длины K на совпадение символов.

func task27(N, K string) {
	fmt.Println("Задание 27 O(NK)")
	defer fmt.Println("Конец Задания 27")
	for i := 0; i < len(N); i++ {
		for j := 0; j < len(K); j++ {
			if N[i] == K[j] {
				fmt.Printf("Совпадение символов: %c на позиции %d в N и %d в K\n", N[i], i, j)
			}
		}
	}
}

// * 28. Найдите количество всех возможных перестановок двух строк, одна длиной N, другая — K.

func task28(N, K string) {
	fmt.Println("Задание 28 O(NK)")
	defer fmt.Println("Конец Задания 28")
	permutations := 0
	var allcombines map[string]bool = make(map[string]bool)
	for i := 0; i < len(N); i++ {
		for j := 0; j < len(K); j++ {
			/* // ! Решение в O(1)
			// * Чистые перестановки занимают n! (факториал числа) при совмещении двух перестановок используется умножение
			    permutations += factorial(len(N)) * factorial(len(K))*/
			allcombines["12"] = true
		}
	}
	fmt.Printf("Количество возможных перестановок: %d\n", permutations)
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// * 29. Напишите функцию для проверки того, содержат ли два числа из диапазонов N и K одинаковое количество делителей.

func task29(N, K []int) {
	fmt.Println("Задание 29 O(NK)")
	defer fmt.Println("Конец Задания 29")
	for _, n := range N {
		for _, k := range K {
			//fmt.Println(n, " Имеет ", countDivisors(n))
			//fmt.Println(k, " Имеет ", countDivisors(k))
			if countDivisors(n) == countDivisors(k) {
				fmt.Printf("%d и %d имеют одинаковое количество делителей\n", n, k)
			}
		}
	}
}

func countDivisors(num int) int {
	count := 0
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			count++
		}
	}
	return count
}

// * 30. Найдите все общие простые множители для чисел из диапазона N и K.

func task30(N, K []int) {
	fmt.Println("Задание 30 O(NK)")
	defer fmt.Println("Конец Задания 30")
	for _, n := range N {
		for _, k := range K {
			commonFactors := findCommonPrimeFactors(n, k)
			if len(commonFactors) > 0 {
				fmt.Printf("Общие простые множители для %d и %d: %v\n", n, k, commonFactors)
			}
		}
	}
}

func findCommonPrimeFactors(a, b int) []int {
	factors := []int{}
	for i := 2; i <= a && i <= b; i++ {
		if a%i == 0 && b%i == 0 && isPrime(i) {
			factors = append(factors, i)
		}
	}
	return factors
}
