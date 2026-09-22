// lesson19 ch02: シャドーイング
package main

import "fmt"

func main() {
	x := 1
	fmt.Println("外:", x)

	if x > 0 {
		// := なので、外の x とは別の新しい x が作られます
		x := 100
		x += 5
		fmt.Println("内:", x) // 105
	}

	// ブロックを出ると、隠れていた外の x が元の値のまま現れます
	fmt.Println("外:", x) // 1
}
