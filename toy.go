package main

import "fmt"

type toy struct {
	title string
	price money
}

func (t *toy) print() { // Method (game struct)
	fmt.Printf("%-15s:  %s\n", t.title, t.price.string())
}

func (t *toy) discount(ratio float64) {
	t.price *= money(1 - ratio)
}
