// lesson16 ch01 解答
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	prices := map[string]int{"apple": 120, "banana": 90, "orange": 60}

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fruit := sc.Text()

	// map の取得は Python の辞書と同じ m[key] です。
	// 入力された文字列をそのままキーにして引けるのが map の強みです。
	fmt.Println(prices[fruit])
}
