// lesson06 ch04 解答
package main

import (
	"fmt"
	"strconv"
)

func main() {
	height := 169.5

	// float64 も文字列と直接 + できないため、まず文字列に変換します。
	// 小数桁数を 1 と明示することで「169.5」という表示が保証されます。
	s := strconv.FormatFloat(height, 'f', 1, 64)
	fmt.Println("身長は" + s + "cmです")
}
