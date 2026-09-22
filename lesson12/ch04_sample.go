// lesson12 ch04: ループのネスト
package main

import "fmt"

func main() {
	// 外側が 1 周する間に、内側は最後まで回ります
	i := 1
	for i <= 3 {
		j := 1
		for j <= 2 {
			fmt.Printf("i=%d j=%d\n", i, j)
			j++
		}
		i++
	}

	// 内側の break は内側のループだけを抜けます。
	// 外側ごと抜けたいときはラベル（outer: と break outer）を使います
	fmt.Println("ネストおわり")
}
