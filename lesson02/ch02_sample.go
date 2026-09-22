// lesson02 ch02: fmt.Print（改行しない出力）
package main

import "fmt"

func main() {
	// Print は改行を付けないので、続けて出力されます
	fmt.Print("Hello")
	fmt.Print(", ")
	fmt.Println("world!") // 最後だけ Println で改行

	// 改行だけしたいときは空の Println を呼びます
	fmt.Print("1行目のつもり")
	fmt.Println()
	fmt.Println("2行目")
}
