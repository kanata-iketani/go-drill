// lesson12 ch01 解答
package main

import "fmt"

func main() {
	// while i <= 5: の置き換えです。カウンタは for の前で宣言し、
	// 本体の最後の i++ を忘れると条件が false にならず無限ループになります。
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
}
