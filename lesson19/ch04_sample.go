// lesson19 ch04: := が作る意図しない新変数
package main

import (
	"fmt"
	"strconv"
)

func main() {
	total := 0

	// if の頭で := した v と err は、この if の中でだけ有効です
	if v, err := strconv.Atoi("10"); err == nil {
		total += v // 中で使い切るならこの形が便利
	}
	fmt.Println(total) // 10

	// 外でも使いたいときは、先に宣言して = で代入します
	var v int
	var err error
	v, err = strconv.Atoi("20")
	if err == nil {
		total += v
	}
	fmt.Println(total) // 30
}
