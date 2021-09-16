package main

import (
	"fmt"
)

func isPrimeNumber(number int) bool {
	a := 0
	for i := 0; i >= 2 && i < number; i++ {
		if number % i == 0 {
			a = 1
		}
	}
	if a == 1 {
		return false
	}
	return true
}

func main() {
  fmt.Println(isPrimeNumber(3)) // 예시로 3을 집어넣은거지만, 어떤 수를 집어넣든 작동이 됩니다 :)
}
