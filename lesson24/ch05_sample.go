// lesson24 ch05: panic と recover（使う場面は少ない）
package main

import "fmt"

func mustPositive(n int) int {
	if n < 0 {
		// 続行できない状況では panic で停止へ向かいます
		// （通常の失敗は error で返すのが原則です）
		panic("負の数は扱えません")
	}
	return n * 10
}

func main() {
	// recover は defer した関数の中でだけ panic を捕まえられます
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover:", r)
		}
	}()

	fmt.Println(mustPositive(3)) // 30
	fmt.Println(mustPositive(-1))
	fmt.Println("ここは実行されません")
}
