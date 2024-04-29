package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"os"
)

func isRoman(x string) bool {
	if strings.ContainsAny(x, "IVXLCDM") { // простая проверка на то, являются ли символы римскими цифрами
		return true
	}
	return false
}

func fromRomanToInt(romanString string) int {

	romanSimbol := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	if _, ok := romanSimbol[romanString]; ok { // сопоставление римских цифр с арабскими
		return romanSimbol[romanString]
	} else {
		panic("формат математической операции не удовлетворяет заданию — " +
			"операнды от 1 до 10 включительно.")
	}
}

func fromIntToRoman(num int) string {

	// Определение соответствий между римскими цифрами и их значением
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	numerals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result string
	for i := 0; i < len(values); i++ {
		// Пока число больше или равно значению в списке values
		for num >= values[i] {
			result += numerals[i] // Добавить соответствующую римскую цифру в результат
			num -= values[i]      // Уменьшить число на значение римской цифры
		}
	}
	return result
}
func plus(str string) string {

	if len(strings.Split(str, "+")) > 2 { // проверка на количество операндов
		panic("формат математической операции не удовлетворяет заданию — " +
			"два операнда и один оператор (+, -, /, *).")
	}
	var str1 = strings.Split(str, "+")[0] //находим операнды в строках
	var str2 = strings.Split(str, "+")[1] //находим операнды в строках
	str1 = strings.TrimSpace(str1)        //очищаем от лишних пустот
	str2 = strings.TrimSpace(str2)        //очищаем от лишних пустот
	var num1, num2 int
	switch {
	case isRoman(str1) && isRoman(str2):
		num1 = fromRomanToInt(str1) //конвертируем строки в числа
		num2 = fromRomanToInt(str2) //конвертируем строки в числа
	case isRoman(str1) != isRoman(str2):
		panic("используются одновременно разные системы счисления")
	default:
		num1, _ = strconv.Atoi(str1) //конвертируем строки в числа
		num2, _ = strconv.Atoi(str2) //конвертируем строки в числа
	}

	if num1 <= 10 && num1 >= 1 && num2 <= 10 && num2 >= 1 {
		result := num1 + num2
		if isRoman(str1) {
			return fromIntToRoman(result)
		} else {
			return strconv.Itoa(result)
		}
	} else {
		panic("формат математической операции не удовлетворяет заданию — " +
			"операнды от 1 до 10 включительно.")
	}

}

func deduction(str string) string {
	if len(strings.Split(str, "-")) > 2 { // проверка на количество операндов
		panic("формат математической операции не удовлетворяет заданию — " +
			"два операнда и один оператор (+, -, /, *).")
	}
	var str1 = strings.Split(str, "-")[0] //находим операнды в строках
	var str2 = strings.Split(str, "-")[1] //находим операнды в строках
	str1 = strings.TrimSpace(str1)        //очищаем от лишних пустот
	str2 = strings.TrimSpace(str2)        //очищаем от лишних пустот

	var num1, num2 int
	switch {
	case isRoman(str1) && isRoman(str2):
		num1 = fromRomanToInt(str1) //конвертируем строки в числа
		num2 = fromRomanToInt(str2) //конвертируем строки в числа
		if num1 < num2 {
			panic("в римской системе нет отрицательных чисел")
		}
	case isRoman(str1) != isRoman(str2):
		panic("используются одновременно разные системы счисления")
	default:
		num1, _ = strconv.Atoi(str1) //конвертируем строки в числа
		num2, _ = strconv.Atoi(str2) //конвертируем строки в числа
	}

	if num1 <= 10 && num1 >= 1 && num2 <= 10 && num2 >= 1 { // проверка на соответствие условиям
		result := num1 - num2
		switch {
		case isRoman(str1):
			return fromIntToRoman(result)
		}
		return strconv.Itoa(result)
	} else {
		panic("формат математической операции не удовлетворяет заданию — " +
			"операнды от 1 до 10 включительно.")
	}

}

func multiplication(str string) string {
	if len(strings.Split(str, "*")) > 2 { // проверка на количество операндов
		panic("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	var str1 = strings.Split(str, "*")[0] //находим операнды в строках
	var str2 = strings.Split(str, "*")[1] //находим операнды в строках
	str1 = strings.TrimSpace(str1)        //очищаем от лишних пустот
	str2 = strings.TrimSpace(str2)        //очищаем от лишних пустот

	var num1, num2 int
	switch {
	case isRoman(str1) && isRoman(str2):
		num1 = fromRomanToInt(str1) //конвертируем строки в числа
		num2 = fromRomanToInt(str2) //конвертируем строки в числа
	case isRoman(str1) != isRoman(str2):
		panic("используются одновременно разные системы счисления")
	default:
		num1, _ = strconv.Atoi(str1) //конвертируем строки в числа
		num2, _ = strconv.Atoi(str2) //конвертируем строки в числа
	}

	if num1 <= 10 && num1 >= 1 && num2 <= 10 && num2 >= 1 { // проверка на соответствие условиям
		result := num1 * num2
		switch {
		case isRoman(str1):
			return fromIntToRoman(result)
		}
		return strconv.Itoa(result)
	} else {
		panic("формат математической операции не удовлетворяет заданию — " +
			"операнды от 1 до 10 включительно.")
	}

}

func division(str string) string {
	if len(strings.Split(str, "/")) > 2 { // проверка на количество операндов
		panic("формат математической операции не удовлетворяет заданию — " +
			"два операнда и один оператор (+, -, /, *).")
	}
	var str1 = strings.Split(str, "/")[0] //находим операнды в строках
	var str2 = strings.Split(str, "/")[1] //находим операнды в строках
	str1 = strings.TrimSpace(str1)        //очищаем от лишних пустот
	str2 = strings.TrimSpace(str2)        //очищаем от лишних пустот

	var num1, num2 int
	switch {
	case isRoman(str1) && isRoman(str2):
		num1 = fromRomanToInt(str1) //конвертируем строки в числа
		num2 = fromRomanToInt(str2) //конвертируем строки в числа
	case isRoman(str1) != isRoman(str2):
		panic("используются одновременно разные системы счисления")
	default:
		num1, _ = strconv.Atoi(str1) //конвертируем строки в числа
		num2, _ = strconv.Atoi(str2) //конвертируем строки в числа
	}

	if num1 <= 10 && num1 >= 1 && num2 <= 10 && num2 >= 1 { // проверка на соответствие условиям
		result := num1 % num2
		switch {
		case isRoman(str1):
			return fromIntToRoman(result)
		}
		return strconv.Itoa(result)
	} else {
		panic("формат математической операции не удовлетворяет заданию — " +
			"операнды от 1 до 10 включительно.")
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите значение ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		switch {
		case strings.Contains(text, "+"):
			fmt.Println(plus(text))

		case strings.Contains(text, "-"):
			fmt.Println(deduction(text))

		case strings.Contains(text, "*"):
			fmt.Println(multiplication(text))

		case strings.Contains(text, "/"):
			fmt.Println(division(text))
		default:
			panic("формат математической операции не удовлетворяет заданию —" +
				" два операнда и один оператор (+, -, /, *).")

		}

	}
}
