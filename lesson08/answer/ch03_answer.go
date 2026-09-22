// lesson08 ch03 解答
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 引数なしの split() に当たる Fields は、連続する空白を
	// まとめて 1 つの区切りとして扱うため、空要素が混ざりません。
	words := strings.Fields("梅 竹  松")

	fmt.Println(words)
	fmt.Println(len(words))
}
