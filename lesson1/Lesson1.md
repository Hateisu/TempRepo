# Tasks lesson 1

[Link to Live](https://go.dev/play/p/LXNCDchapOz)
Some update
```Golang
package main

import (
	"golang.org/x/exp/errors/fmt"
)

func main() {

	task1(4, 2)
	task2(3, 4)
	fmt.Println("Введите 3 числа : ")
	var num1, num2, num3 int
	// ! FOR LOCAL USAGE
	// fmt.Scanln(&num1)
	// fmt.Scanln(&num2)
	// fmt.Scanln(&num3)
	num1 = 1
	num2 = 2
	num3 = 3
	task3(num1, num2, num3)
	task4(10)
	task5(5, 10)
	task6(5)
	task7("Nigga", "12")
	task8(900, 10)
	task9(1, 2)
	task10(12)
	task11(1, 2, 3)
	task12(4)
	task13()
	task14(7)
	//! For local usage
	//task15local()
	task15([]int{1, 2, 3, 4, 5, 6, 0, 1000})
	//! For local usage
	//task16local()
	task16("Secret")
	task16("SuperSecret")
	task17()
	task18()
	//! For local usage
	//task19local()
	task19(3)
	task19(6)
	task20()
}

// * 1. Напишите программу, которая запрашивает два числа и выводит их сумму, разность, произведение и частное.
func task1(a, b int) {
	fmt.Println("Задание 1")
	fmt.Println(a, " + ", b, " = ", (a + b))
	fmt.Println(a, " - ", b, " = ", (a - b))
	fmt.Println(a, " * ", b, " = ", (a * b))
	fmt.Println(a, " / ", b, " = ", (a / b))
}

// * 2. Создайте программу, которая вычисляет площадь прямоугольника по введённым пользователем длине и ширине.
func task2(width, height int) {
	fmt.Println("Задание 2")
	fmt.Println("S = ", (width * height))
}

// * 3. Реализуйте программу, которая принимает от пользователя три числа и выводит их среднее арифметическо
func task3(a, b, c int) {
	fmt.Println("Задание 3")
	fmt.Println("средн ариф = ", ((a + b + c) / 3))
}

// * 4. Напишите программу, которая запрашивает у пользователя число и увеличивает его на 10, выводя результат.
func task4(a int) {
	fmt.Println("Задание 4")
	fmt.Println("Ответ ", a+10)
}

// * 5. Создайте программу, которая меняет значения двух переменных местами с помощью третьей переменной.
func task5(a, b int) {
	fmt.Println("Задание 5")
	fmt.Println("До")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	temp := a
	a = b
	b = temp
	fmt.Println("После")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

}

// * 6. Напишите программу, которая вычисляет квадрат числа, введённого пользователем, и сохраняет результат в новой переменной.
func task6(a int) {
	fmt.Println("Задание 6")
	var result int = a * a
	fmt.Println("Квадрат числа ", a, " = ", result)
}

// * 7. Напишите программу, которая запрашивает имя и возраст пользователя и выводит приветственное сообщение в формате: "Привет, {имя}, тебе {возраст} лет!"
func task7(name, age string) {
	fmt.Println("Задание 7")
	fmt.Println("Привет, ", name, ", тебе ", age, " лет!")
}

// * 8. Создайте программу, которая запрашивает у пользователя стоимость товара и процент скидки, а затем выводит стоимость товара после применения скидки.
func task8(price, off float32) {
	fmt.Println("Задание 8")
	fmt.Println("Стоимость товара до скидки = ", price)
	var result float32 = price - (price * (off / 100))
	fmt.Println("Стоимость товара после скидки = ", result)
}

// * 9. Напишите программу, которая принимает от пользователя два числа и выводит, какое из них больше, или сообщает, что числа равны.
func task9(a, b int) {
	fmt.Println("Задание 9")
	if a > b {
		fmt.Println(a, " больше ", b)
	} else if b > a {
		fmt.Println(b, " больше ", a)
	} else {
		fmt.Println(a, " и ", b, " равны")
	}
}

// * 10.Напишите программу, которая принимает на вход число и проверяет, положительное оно, отрицательное или равно нулю.
func task10(a int) {
	fmt.Println("Задание 10")
	if a > 0 {
		fmt.Println(a, " положительное число")
	} else if a < 0 {
		fmt.Println(a, " отрицательное число")
	} else {
		fmt.Println(a, " равно нулю")
	}
}

// * 11.Создайте программу, которая принимает три числа и выводит наибольшее из них.
func task11(a, b, c int) {
	fmt.Println("Задание 11")
	if a > b && a > c {
		fmt.Println(a, " больше всех чисел")
	} else if b > a && b > c {
		fmt.Println(b, " больше всех чисел")
	} else {
		fmt.Println(c, " больше всех чисел")
	}
}

// * 12.Реализуйте программу, которая проверяет, является ли введённое число чётным, и если да, то дополнительно проверяет, делится ли оно на 4.
func task12(a int) {
	fmt.Println("Задание 12")
	if a%2 == 0 {
		fmt.Println(a, " чётное число")
		if a%4 == 0 {
			fmt.Println("и делится на 4")
		}
	} else {
		fmt.Println(a, " нечётное число")
	}
}

// * 13.Напишите программу, которая выводит все числа от 1 до 10 с помощью цикла while.
func task13() {
	fmt.Println("Задание 13")
	i := 1
	for i <= 10 {
		if i != 10 {
			fmt.Print(i, ", ")
		} else {
			fmt.Print(i)

		}
		i++
	}
	fmt.Println()
}

// * 14.Создайте программу, которая принимает на вход число и выводит его таблицу умножения от 1 до 10.
func task14(a int) {
	fmt.Println("Задание 14")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", a, i, a*i)
	}
}

// * 15.Реализуйте программу, которая запрашивает у пользователя числа, пока он не введёт число 0, и выводит сумму всех введённых чисел.
func task15local() {
	fmt.Println("Задание 15")
	var sum int
	for {
		var num int
		fmt.Print("Введите число (0 для выхода): ")
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		sum += num
	}
	fmt.Println("Сумма всех введённых чисел:", sum)
}
func task15(slice []int) {
	fmt.Println("Задание 15")
	var sum int
	for _, num := range slice {
		if num == 0 {
			break
		}
		sum += num
	}
	fmt.Println("Сумма всех введённых чисел:", sum)
}

// * 16.Напишите программу, которая запрашивает у пользователя пароль, и если он неверный, программа продолжает запрашивать его, пока не будет введён правильный пароль (предположим, что пароль — "1234").
const secret = "SuperSecret"

func task16local() {
	fmt.Println("Задание 16")
	var password string
	for {
		fmt.Print("Введите пароль: ")
		fmt.Scan(&password)
		if password == secret {
			break
		}
		fmt.Println("Пароль неверный!")
	}
	fmt.Println("Пароль успешно введён!")
}
func task16(password string) {
	fmt.Println("Задание 16")

	if password == secret {
		fmt.Println("Пароль успешно введён!")
	} else {

		fmt.Println("Пароль неверный!")
	}

}

// * 17.Реализуйте программу, которая выводит числа от 1 до 10, используя цикл с постусловием.
func task17() {
	fmt.Println("Задание 17")
	i := 1
	for ok := true; ok; ok = (i <= 10) {
		if i != 10 {
			fmt.Print(i, ", ")
		} else {
			fmt.Print(i)

		}
		i++
	}
	fmt.Println()
}

// * 18.Напишите программу, которая выводит все чётные числа от 1 до 20 с помощью цикла for.
func task18() {
	fmt.Println("Задание 18")
	for i := 1; i <= 20; i += 2 {
		if i != 20-1 {
			fmt.Print(i, ", ")
		} else {
			fmt.Print(i)
		}
	}
	fmt.Println()
}

// * 19.Создайте программу, которая принимает от пользователя число n и выводит все числа от 1 до n.
func task19local() {
	fmt.Println("Задание 19")
	var n int
	fmt.Print("Введите число n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i != n {
			fmt.Print(i, ", ")
		} else {
			fmt.Print(i)

		}
	}
	fmt.Println()
}
func task19(n int) {
	fmt.Println("Задание 19")
	for i := 1; i <= n; i++ {
		if i != n {
			fmt.Print(i, ", ")
		} else {
			fmt.Print(i)

		}
	}
	fmt.Println()
}

// * 20.Напишите программу, которая вычисляет сумму чисел от 1 до 100 и выводит результат.
func task20() {
	fmt.Println("Задание 20")
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Сумма чисел от 1 до 100:", sum)
}

```

