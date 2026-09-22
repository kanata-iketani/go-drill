// lesson08 ch02 解答
package main

import (
	"fmt"
	"strings"
)

func main() {
	// "apple:banana:cherry".split(":") に当たります。
	// 結果は []string なので、インデックスや len がそのまま使えます。
	fruits := strings.Split("apple:banana:cherry", ":")

	fmt.Println(fruits)
	fmt.Println(fruits[0])
}
