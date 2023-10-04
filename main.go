package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Калькулятор готов к работе!")

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		input := scanner.Text()
		if input != `` {
			if res := calc(input); res != `` {
				fmt.Println("Порешали: " + res)
			}
		}
	}
}

func calc(input string) string {
	var arabicPattern = `^(?P<First>10|[1-9])\s*(?P<Operand>[+-/*])\s*(?P<Second>10|[1-9])\z`
	var re = regexp.MustCompile(arabicPattern)
	var arabicExpr = re.FindStringSubmatch(input)

	if len(arabicExpr) > 0 {
		return arabicSolver(arabicExpr)
	}
	{
		romanPattern := `I|II|III|IV|V|VI|VII|VIII|IX|X`
		romanNumbers := strings.Split(`|`+romanPattern, `|`)
		romanPattern = `^(?P<First>` + romanPattern + `)\s*(?P<Operand>[+-/*])\s*(?P<Second>` + romanPattern + `)\z`
		var re = regexp.MustCompile(romanPattern)
		var romanExpr = re.FindStringSubmatch(input)
		if len(romanExpr) > 0 {
			return romanSolver(romanExpr, romanNumbers)
		}
		{
			fmt.Println(`ОШИБКА! Неверный ввод.`)
			os.Exit(-1)
		}

	}
	return ""
}
func romanSolver(expr, romanNumbers []string) string {
	var first, second = slices.Index(romanNumbers, expr[1]), slices.Index(romanNumbers, expr[3])
	if first > -1 && second > -1 {
		result := reshalo(first, expr[2], second)
		if result > 0 {
			return convert1_100ToRomans(result, romanNumbers)
		}
		{
			fmt.Println(`ОШИБКА! Римская нотация не имеет нуля или отрицательных чисел.`)
		}
	}
	return ""
}
func arabicSolver(expr []string) string {
	var first, err1 = strconv.Atoi(expr[1])
	var second, err2 = strconv.Atoi(expr[3])
	if err1 == nil && err2 == nil {
		return strconv.Itoa(reshalo(first, expr[2], second))
	}
	return ""
}
func convert1_100ToRomans(inInt int, romanNumerals []string) string {
	var resRom = ``
	switch {
	case inInt == 100:
		resRom = `C`
		inInt -= 100
	case 89 < inInt:
		resRom = `XC`
		inInt -= 90
	case 49 < inInt:
		resRom = `L`
		inInt -= 50
	case 39 < inInt:
		resRom = `XL`
		inInt -= 40
	}

	if inInt >= 10 {
		var decNum int = inInt / 10
		inInt %= 10
		for i := 0; i < decNum; i++ {
			resRom += `X`
		}
	}

	resRom += romanNumerals[inInt]

	return resRom

}
func reshalo(first int, operand string, second int) int {
	var result int
	switch operand {
	case "+":
		{
			result = first + second
		}
	case "-":
		{
			result = first - second
		}
	case "/":
		{
			result = first / second
		}
	case "*":
		{
			result = first * second
		}
	}
	return result
}
