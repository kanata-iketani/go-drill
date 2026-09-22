// lesson20 ch02: ポインタ渡し（*int と &）
package main

import "fmt"

// *int は「int のポインタ」。呼び出し元の変数の場所を受け取ります
func double(p *int) {
	*p *= 2 // * を付けると指した先の値を読み書きできる
}

func main() {
	x := 5
	fmt.Println(x) // 5

	double(&x)     // & で x の場所（ポインタ）を渡す
	fmt.Println(x) // 10: 呼び出し元の x が書き換わった

	double(&x)
	fmt.Println(x) // 20
}
