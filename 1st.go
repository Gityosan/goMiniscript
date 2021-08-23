package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Good morning")
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Int31())
	fmt.Println(rand.Int31n(10))
	fmt.Println()
	fmt.Println(time.Now())
	fmt.Println(time.Now().UnixNano())
	t := time.Now()
	fmt.Println(t)
	fmt.Println(time.January)
	fmt.Println(time.Microsecond.Round(10))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(time.Now().UTC())
	time, _ := time.ParseDuration("1h30m")
	fmt.Println(time.Minutes())
}
