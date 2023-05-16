package main

import "fmt"

func main() {
	var consume, discount float64
	var dataDiscount string
	fmt.Print("Consumo: ")
	fmt.Scanln(&consume)

	if consume >= 0 {
		dataDiscount = "10%"
		fmt.Println("Discount % :", dataDiscount)
		discount = consume * 0.10
		totalDiscount := consume - discount
		igv := totalDiscount * 0.19
		totalPay := totalDiscount + igv
		fmt.Println("Discount ", discount)
		fmt.Println("totalDiscount ", totalDiscount)
		fmt.Println("igv ", igv)
		fmt.Println("totalPay ", totalPay)

	} else {
		fmt.Println("Consume must have positive number")
	}
}
