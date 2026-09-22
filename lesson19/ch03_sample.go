// lesson19 ch03: パッケージ変数（global 文はない）
package main

import "fmt"

// 関数の外で宣言した変数は、どの関数からも読み書きできます
var balance = 1000

// Python と違い global 宣言は不要（そもそも存在しません）
func deposit(n int) {
	balance += n
}

func withdraw(n int) {
	balance -= n
}

func main() {
	fmt.Println("最初:", balance)
	deposit(500)
	withdraw(200)
	fmt.Println("最後:", balance)
}
