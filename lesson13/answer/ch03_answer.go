// lesson13 ch03 解答
package main

import "fmt"

func main() {
	prices := []int{100, 250, 380}

	// 合計にインデックスは不要なので、1 つ目を _ で捨てて値だけ受け取ります。
	// i と名前を付けて使わないと「declared and not used」のエラーになります。
	total := 0
	for _, p := range prices {
		total += p
	}
	fmt.Println(total)
}
