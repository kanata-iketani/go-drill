// lesson21 ch04: 大文字小文字で公開制御
package main

import "fmt"

// Wallet は公開（大文字）、balance は非公開（小文字）です。
// 別パッケージからは w.balance と直接触れず、メソッド経由になります
type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(n int) {
	if n > 0 { // メソッドを通すので、こうした検証を挟めます
		w.balance += n
	}
}

func (w *Wallet) Balance() int {
	return w.balance
}

func main() {
	w := Wallet{}
	w.Deposit(100)
	w.Deposit(-50)           // 不正な入金は Deposit がはじく
	fmt.Println(w.Balance()) // 100
}
