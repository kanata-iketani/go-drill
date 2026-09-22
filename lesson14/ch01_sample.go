// lesson14 ch01: Go にリスト内包表記はない
package main

import "fmt"

func main() {
	// Python: squares = [x * x for x in range(5)]
	// Go は「空スライス → for → append」の 3 行で書きます
	squares := []int{}
	for x := range 5 {
		squares = append(squares, x*x)
	}
	fmt.Println(squares) // [0 1 4 9 16]

	// 既存のスライスから作る場合も形は同じです
	prices := []int{100, 250, 80}
	taxed := []int{}
	for _, p := range prices {
		taxed = append(taxed, p*11/10)
	}
	fmt.Println(taxed) // [110 275 88]
}
