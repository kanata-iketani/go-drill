// lesson13 ch01 解答
package main

import "fmt"

func main() {
	colors := []string{"red", "green", "blue"}

	// Python の enumerate(colors) に相当します。range がインデックスと
	// 値を同時に返すため、カウンタ変数を自分で管理する必要がありません。
	for i, v := range colors {
		fmt.Println(i, v)
	}
}
