// lesson25 ch03: ポインタで呼び出し元の変数を変える
package main

import "fmt"

// 値渡し: コピーが増えるだけで呼び出し元は変わりません
func addTenValue(n int) {
	n += 10
}

// ポインタ渡し: *p で指し先を書き換えると呼び出し元が変わります
func addTen(p *int) {
	*p += 10
}

func main() {
	x := 5
	addTenValue(x)
	fmt.Println(x) // 5（変わらない）

	addTen(&x)     // &x で x のポインタ（場所）を渡す
	fmt.Println(x) // 15

	p := &x
	fmt.Println(*p) // 15（*p で指し先を読む）
}
