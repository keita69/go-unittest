package main

import (
	"fmt"
)

func main() {
	var book float64 = 3000
	fmt.Printf("消費税【%d％】込み価格は【%d】円です", int(getTax()*100), calc(book))
}

func getTax() float64 {
	tax := 0.10
	return tax
}

func calc(price float64) int {
	return int(price * (1.0 + getTax()))
}
