// lesson19 ch01: ブロックスコープ {}
package main

import "fmt"

func main() {
	x := 5

	// 外で使いたい変数は、ブロックの外で先に宣言します
	msg := ""
	if x > 0 {
		msg = "plus" // 代入なら外の msg に入る
		inner := "if の中だけの変数"
		fmt.Println(inner)
	}
	// fmt.Println(inner) // ← ブロックの外なのでコンパイルエラー
	fmt.Println(msg)

	total := 0
	for i := range 3 {
		total += i // ループ変数 i もループの中だけで有効
	}
	fmt.Println(total)
}
