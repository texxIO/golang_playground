package main

import "fmt"

func startConditionals() {
	fmt.Println("****Starting Conditionals****")
}

func forLoop() (sum int) {

	sum = 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	return
}

func forContinued() int {
	sum := 1

	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
	}
	return sum
}

func ifStatement(x, y int) bool {

	if x > y {
		return true
	}

	return false
}

func ifStatement2() int {
	var test = true

	if test {
		return 1
	} else {
		return 0
	}
}
