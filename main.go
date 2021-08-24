package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var num = inputNumber()
	var flag = true
	var sqrt = int(math.Sqrt(float64(num)))
	for i := 3; i <= sqrt && flag; i += 2 {
		if num%i == 0 {
			flag = false
		}
	}
	fmt.Printf("Результат: %t\n", flag)
}

func inputNumber() int {
	fmt.Print("Какое число проверять: ")
	var num = 0
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		os.Exit(1)
	}
	return num
}
