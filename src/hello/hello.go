package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Let's start!!!")
	fmt.Printf("The time is %s \n", time.Now())
	fmt.Println("My favorite number is:", rand.Intn(10))
	startConditionals()
	fmt.Println(forLoop())
	fmt.Println(forContinued())
	fmt.Println(ifStatement(8, 6))
	fmt.Println(ifStatement2())

}
