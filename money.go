package main

import "fmt"

type money float64

// func (m money) discountprice() {
// 	// m *= (1 - m)
// 	m = m * (1 - 0.2) // 20%
// 	fmt.Println("You get '20%' of discount")
// 	fmt.Printf("Total %.1f$\n", m)
// }

func (m money) string() string {
	return fmt.Sprintf("$%.2f", m)
}
