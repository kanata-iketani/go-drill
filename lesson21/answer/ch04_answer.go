// lesson21 ch04 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// balance を小文字にしておくと、パッケージの外からは直接読み書きできません。
// Python の _balance と違い、Go ではこれが言語仕様として強制されるため、
// 「必ず Deposit を通る」ことをコンパイラが保証してくれます。
type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(n int) {
	w.balance += n
}

func (w *Wallet) Balance() int {
	return w.balance
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fields := strings.Fields(sc.Text())
	w := Wallet{}
	for _, f := range fields {
		n, _ := strconv.Atoi(f)
		w.Deposit(n)
	}
	fmt.Println(w.Balance())
}
