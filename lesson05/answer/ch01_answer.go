// lesson05 ch01 解答
package main

import "fmt"

func main() {
	family := "山田"
	given := "太郎"

	// Println にカンマで渡しても空白区切りになりますが、
	// この演習は「+ で連結」が課題なので、間の空白も
	// 文字列 " " として自分で連結しています。
	fmt.Println(family + " " + given)
}
