// lesson12 ch01: for 条件 {}（while の代わり）
package main

import "fmt"

func main() {
	// while i <= 3: に相当します。条件に () は付けません
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++ // 更新を忘れると無限ループになります
	}

	// 条件が最初から false なら 1 回も実行されません
	n := 10
	for n < 10 {
		fmt.Println("ここは実行されません")
	}
	fmt.Println("おわり")
}
