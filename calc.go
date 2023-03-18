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
	switch a {
	case 1:
		return "I", nil
	case 2:
		return "II", nil
	case 3:
		return "III", nil
	case 4:
		return "IV", nil
	case 5:
		return "V", nil
	case 6:
		return "VI", nil
	case 7:
		return "VII", nil
	case 8:
		return "VIII", nil
	case 9:
		return "IX", nil
	case 10:
		return "X", nil
	case 11:
		return "XI", nil
	case 12:
		return "XII", nil
	case 13:
		return "XIII", nil
	case 14:
		return "XIV", nil
	case 15:
		return "XV", nil
	case 16:
		return "XVI", nil
	case 17:
		return "XVII", nil
	case 18:
		return "XVIII", nil
	case 19:
		return "XIX", nil
	case 20:
		return "XX", nil
	default:
		return "", errors.New("Некорректное арабское число")
	}

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
	fmt.Println("tessst")
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
