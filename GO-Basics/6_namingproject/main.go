package main

import (
	"bufio"
	"calculate/models"
	"calculate/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// unexported camelCase
func getNumber(prompt string) float64 {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Geçersiz sayı! Tekrar deneyin.")
			continue
		}

		return num
	}
}

// unexported camelCase
func getOperation() string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("İşlem seç (+, -, *, /): ")
		op, _ := reader.ReadString('\n')
		op = strings.TrimSpace(op)

		if op == "+" || op == "-" || op == "*" || op == "/" {
			return op
		}

		fmt.Println("Geçersiz işlem! Tekrar seç.")
	}
}

func main() {

	// Struct kullanımı
	calc := models.Calculation{}

	calc.FirstNumber = getNumber("İlk sayıyı gir: ")
	calc.SecondNumber = getNumber("İkinci sayıyı gir: ")
	calc.Operation = getOperation()

	switch calc.Operation {
	case "+":
		calc.Result = utils.Add(calc.FirstNumber, calc.SecondNumber)
	case "-":
		calc.Result = utils.Subtract(calc.FirstNumber, calc.SecondNumber)
	case "*":
		calc.Result = utils.Multiply(calc.FirstNumber, calc.SecondNumber)
	case "/":
		if calc.SecondNumber == 0 {
			fmt.Println("0'a bölme hatası!")
			return
		}
		calc.Result = calc.FirstNumber / calc.SecondNumber
	}

	fmt.Printf("\nSonuç: %.2f %s %.2f = %.2f\n",
		calc.FirstNumber,
		calc.Operation,
		calc.SecondNumber,
		calc.Result)
}
