package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your input: ")
	expression, _ := reader.ReadString('\n')
	result, err := calculate(expression)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println("Result:", result)
}

func calculate(expression string) (string, error) {
	expression = strings.TrimSpace(expression)
	parts := strings.Split(expression, " ")
	if len(parts) != 3 {
		return "", fmt.Errorf("Wrong expression")
	}
	identical, system := numberTypeIdentical(parts[0], parts[2])
	if !identical {
		return "", fmt.Errorf("Must be both arabic or roman numbers")
	}
	x, err := parseNumber(parts[0])
	if err != nil {
		return "", fmt.Errorf("Error while parsing a number: %s", err.Error())
	}
	y, err := parseNumber(parts[2])
	if err != nil {
		return "", fmt.Errorf("Error while parsing a number: %s", err.Error())
	}
	operator := parts[1]
	var result int

	switch operator {
	case "+":
		result = x + y
		if result <= 0 && system == "roman" {
			return "", fmt.Errorf("Roman number can not be negative or be equal to zero.")
		}
	case "-":
		result = x - y
		if result <= 0 && system == "roman" {
			return "", fmt.Errorf("Roman number can not be negative or be equal to zero.")
		}
	case "*":
		result = x * y
		if result <= 0 && system == "roman" {
			return "", fmt.Errorf("Roman number can not be negative or be equal to zero.")
		}
	case "/":
		if y == 0 {
			return "", fmt.Errorf("Division by zero")
		}
		result = x / y
		if result <= 0 && system == "roman" {
			return "", fmt.Errorf("Roman number can not be negative or be equal to zero.")
		}
	default:
		return "", fmt.Errorf("Wrong operand")
	}
	if system == "roman" {
		return integerToRoman(result), nil
	} else {
		return strconv.Itoa(result), nil
	}
}

func numberTypeIdentical(numberOne string, numberTwo string) (bool, string) {
	_, err1 := strconv.Atoi(numberOne)
	_, err2 := strconv.Atoi(numberTwo)

	if err1 == nil && err2 == nil {
		return true, "arabic"
	} else if err1 != nil && err2 != nil {
		return true, "roman"
	} else {
		return false, ""
	}

}

func integerToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

func parseNumber(numberStr string) (int, error) {
	x, err := strconv.Atoi(numberStr)
	if err == nil {
		if x < 1 || x > 10 {
			return 0, fmt.Errorf("The number should be from 1 to 10.")
		}
		return x, nil
	}
	// If the number could not be parsed as Arabic, then we are try to parse it as Roman numeral.
	romanNumbers := map[string]int{
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
	x, ok := romanNumbers[numberStr]
	if !ok {
		return 0, fmt.Errorf("Wrong number: %s", numberStr)
	}
	return x, nil
}
