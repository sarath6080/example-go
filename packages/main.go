package main

import (
	"fmt"
	"golang-consumer/consumer"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := consumer.Average(xs)
	fmt.Println(avg)
}
