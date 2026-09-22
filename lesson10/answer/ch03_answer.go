// lesson10 ch03 解答
package main

import "fmt"

func main() {
	word := "golang"

	// 判定にしか使わない変数 l は、条件の前の 1 文で宣言すると
	// この if / else の中だけに閉じ込められ、外のスコープを汚しません。
	if l := len(word); l >= 5 {
		fmt.Println("5文字以上:", l)
	} else {
		fmt.Println("5文字未満")
	}
}
