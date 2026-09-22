// lesson09 ch03: 配列 [3]int（固定長・値型）
package main

import "fmt"

func main() {
	// [] の中に長さを書くと配列になります。長さは変えられません
	a := [3]int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(len(a))

	// 値型: 代入すると中身ごとコピーされます
	b := a
	b[0] = 99
	fmt.Println(b)
	fmt.Println(a) // コピー元は変わりません
}
