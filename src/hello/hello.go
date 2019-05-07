package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("hello, world!\n")
	fmt.Printf("The time is %s", time.Now())
	fmt.Print("My favorite number is:", rand.Intn(10))

}
