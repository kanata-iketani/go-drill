// lesson20 ch01: 値渡し（引数はコピーされる）
package main

import "fmt"

// n には呼び出し元の値の「コピー」が入ります
func addTen(n int) int {
	n += 10 // コピーを書き換えているだけ
	return n
}

func main() {
	x := 5
	result := addTen(x)

	fmt.Println(result) // 15
	fmt.Println(x)      // 5: 呼び出し元の x は変わらない

	// 変えたいなら戻り値を代入し直します
	x = addTen(x)
	fmt.Println(x) // 15
}
