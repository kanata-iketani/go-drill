// lesson09 ch03 解答
package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}

	// 配列は値型なので、b := a の時点で中身ごとコピーされます。
	// スライスと違い、b を書き換えても a には影響しません。
	b := a
	b[0] = 99

	fmt.Println(a)
	fmt.Println(b)
}
