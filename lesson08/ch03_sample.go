// lesson08 ch03: strings.Fields（空白区切りの分割）
package main

import (
	"fmt"
	"strings"
)

func main() {
	line := "go  is   fun"

	// 引数なしの split() に当たります。連続空白もまとめて区切ります
	words := strings.Fields(line)
	fmt.Println(words)
	fmt.Println(len(words))

	// Split(" ") だと空文字列の要素が混ざり、数が変わってしまいます
	fmt.Println(len(strings.Split(line, " ")))
}
