// lesson05 ch03: []rune で日本語を文字単位に
package main

import "fmt"

func main() {
	s := "こんにちは世界"

	// []rune に変換すると 1 要素 = 1 文字になります
	r := []rune(s)

	fmt.Println("文字数:", len(r)) // 7

	// [i] で i 番目の文字（rune）を取り出せます
	fmt.Println(string(r[0])) // こ
	fmt.Println(string(r[5])) // 世

	// r[i:j] で部分文字列も文字単位で切り出せます
	fmt.Println(string(r[5:7])) // 世界
}
