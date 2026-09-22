// lesson04 ch03 解答
package main

import "fmt"

func main() {
	balance := 1000 // 最初の 1 回だけ := で作る
	fmt.Println("残高:", balance)

	// 2 回目以降に := を書くと「新しい変数を作り直す」ことになるため、
	// 同じ変数を更新し続けたいときは必ず = を使います。
	balance = balance + 500
	fmt.Println("残高:", balance)

	balance = balance - 300
	fmt.Println("残高:", balance)
}
