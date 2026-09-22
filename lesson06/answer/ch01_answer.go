// lesson06 ch01 解答
package main

import (
	"fmt"
	"strconv"
)

func main() {
	count := 15

	// Go では "残り" + count とは書けません（int と string は連結できない）。
	// Python の str(count) に当たる strconv.Itoa で文字列にしてから + でつなぎます。
	fmt.Println("残り" + strconv.Itoa(count) + "個です")
}
