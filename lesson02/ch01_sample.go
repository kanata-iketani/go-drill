// lesson02 ch01: fmt.Println
package main

import "fmt"

func main() {
	// 文字列も数値もそのまま出力できます
	fmt.Println("Hello, world!")
	fmt.Println(42)

	// カンマで並べると空白区切りになり、最後に改行が付きます
	fmt.Println("score", 80)
	fmt.Println("x", 1, "y", 2)
}
