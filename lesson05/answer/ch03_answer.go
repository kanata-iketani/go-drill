// lesson05 ch03 解答
package main

import "fmt"

func main() {
	// 日本語を文字単位で数えたり取り出したりするときは、
	// まず []rune に変換するのが Go の定石です。
	// r[0] は rune（数値）なので、表示用に string() で文字列へ戻します。
	r := []rune("駅前留学")

	fmt.Println("文字数:", len(r))
	fmt.Println("最初の文字:", string(r[0]))
}
