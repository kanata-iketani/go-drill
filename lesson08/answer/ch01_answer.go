// lesson08 ch01 解答
package main

import (
	"fmt"
	"strings"
)

func main() {
	parts := []string{"2026", "09", "21"}

	// "-".join(parts) に当たります。Go では「つなぐスライスが第 1 引数、
	// 区切り文字が第 2 引数」という向きになるだけで、働きは同じです。
	fmt.Println(strings.Join(parts, "-"))
}
