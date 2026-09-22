// lesson03 ch01 解答
package main

import "fmt"

func main() {
	var price int = 120
	var rate float64 = 1.5

	// int と float64 はそのまま掛けられないため、
	// float64(price) で型を揃えてから計算します。
	// Println は 180.0 ではなく 180 と表示します（小数部が 0 なら省略される）。
	fmt.Println(float64(price) * rate)
}
