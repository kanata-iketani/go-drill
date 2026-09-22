// lesson07 ch02: s[i] と要素の変更
package main

import "fmt"

func main() {
	colors := []string{"赤", "青", "黄"}

	// 取り出しは Python と同じ 0 始まりです
	fmt.Println(colors[0])

	// 代入で要素を書き換えられます
	colors[1] = "緑"
	fmt.Println(colors)

	// 負のインデックスはないので、末尾は len を使って指します
	fmt.Println(colors[len(colors)-1])
}
