// lesson06 ch01: strconv.Itoa（数値→文字列）
package main

import (
	"fmt"
	"strconv"
)

func main() {
	age := 28

	// int と文字列は + で直接つなげられないため、先に文字列へ変換します
	s := strconv.Itoa(age)
	fmt.Println("年齢は" + s + "歳です")

	// 変換結果はふつうの文字列なので、len なども使えます
	price := 1200
	ps := strconv.Itoa(price)
	fmt.Println(ps + " は " + strconv.Itoa(len(ps)) + " 文字です")
}
