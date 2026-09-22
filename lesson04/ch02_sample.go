// lesson04 ch02: := による短縮宣言
package main

import "fmt"

func main() {
	// 宣言と代入を同時に行い、型は右辺から推論されます
	name := "gopher" // string
	age := 13       // int
	height := 158.3 // float64

	fmt.Println(name, "は", age, "歳です")
	fmt.Println("身長:", height)

	// 複数の変数を同時に作ることもできます
	x, y := 10, 20
	fmt.Println("x + y =", x+y)
}
