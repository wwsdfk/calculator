package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToInt = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var intToRoman = []string{
	"0", "I", "II", "III", "IV", "V",
	"VI", "VII", "VIII", "IX", "X",
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (например, 3 + 5 или II * III):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result, err := calculate(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}

func calculate(input string) (string, error) {
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		return "", errors.New("некорректный формат ввода")
	}

	num1 := parts[0]
	operator := parts[1]
	num2 := parts[2]

	isRoman := isRoman(num1) && isRoman(num2)
	isArabic := isNumeric(num1) && isNumeric(num2)

	if !isRoman && !isArabic {
		return "", errors.New("числа должны быть либо оба арабскими, либо оба римскими")
	}

	if isRoman {
		a, b := romanToInt[num1], romanToInt[num2]
		result, err := performOperation(a, b, operator)
		if err != nil {
			return "", err
		}
		if result < 1 {
			return "", errors.New("результат меньше единицы недопустим для римских чисел")
		}
		return intToRoman[result], nil
	}

	if isArabic {
		a, _ := strconv.Atoi(num1)
		b, _ := strconv.Atoi(num2)
		result, err := performOperation(a, b, operator)
		if err != nil {
			return "", err
		}
		return strconv.Itoa(result), nil
	}

	return "", errors.New("неизвестная ошибка")
}

func performOperation(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("некорректная операция")
	}
}

func isRoman(input string) bool {
	_, exists := romanToInt[input]
	return exists
}

func isNumeric(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}
