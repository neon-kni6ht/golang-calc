package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertFromRoman(a string) (int, error) {
	switch a {
	case "I":
		return 1, nil
	case "II":
		return 2, nil
	case "III":
		return 3, nil
	case "IV":
		return 4, nil
	case "V":
		return 5, nil
	case "VI":
		return 6, nil
	case "VII":
		return 7, nil
	case "VIII":
		return 8, nil
	case "IX":
		return 9, nil
	case "X":
		return 10, nil
	default:
		return 0, errors.New("Некорректное римское число")
	}
}

func convertToRoman(a int) (string, error) {

	if a > 3999 {
		return "", errors.New("Некорректное арабское число")
	}

	conversions := []struct {
		value     int
		character string
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
		for a >= conversion.value {
			roman.WriteString(conversion.character)
			a -= conversion.value
		}
	}
	return roman.String(), nil
}

func getNumbers(a, b string) (int, int, bool, error) {
	c, e1 := convertFromRoman(a)
	d, e2 := convertFromRoman(b)
	if e1 == nil && e2 == nil {
		return c, d, true, nil
	}
	c, e1 = strconv.Atoi(a)
	d, e2 = strconv.Atoi(b)
	if e1 == nil && e2 == nil && c >= 1 && c <= 10 && d >= 1 && d <= 10 {
		return c, d, false, nil
	}
	return 0, 0, false, errors.New("Некорректный аргумент")
}

func validateSlice(c []string) error {
	if len(c) != 3 {
		return errors.New("Неверное количество аргументов")
	}
	if c[1] != "*" && c[1] != "/" && c[1] != "+" && c[1] != "-" {
		return errors.New("Некорректная операция")
	}
	return nil
}

func getAnswer(a, b int, o string, isRoman bool) (string, error) {
	r := 0
	switch o {
	case "+":
		r = a + b
	case "-":
		r = a - b
	case "*":
		r = a * b
	case "/":
		r = a / b
	}
	if !isRoman {
		return strconv.Itoa(r), nil
	}
	if s, err := convertToRoman(r); err == nil {
		return s, nil
	} else {
		return s, err
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')
	s = strings.Trim(s, "\n")
	c := strings.Split(s, " ")

	if err := validateSlice(c); err != nil {
		fmt.Println(err)
	} else {
		if a, b, isRoman, err := getNumbers(c[0], c[2]); err == nil {
			if s, err1 := getAnswer(a, b, c[1], isRoman); err1 == nil {
				fmt.Println(s)
			} else {
				fmt.Println(err1)
			}
		} else {
			fmt.Println(err)
		}
	}
}
