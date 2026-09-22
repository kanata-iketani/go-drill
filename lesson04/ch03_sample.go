// lesson04 ch03: 再代入は = / 未使用変数はエラー
package main

import "fmt"

func main() {
	score := 80 // 新しく作るときは :=
	fmt.Println("最初の score:", score)

	score = 95 // 入れ直すときは =
	fmt.Println("再代入後の score:", score)

	// score = "high"  // コンパイルエラー: int に文字列は入らない
	// unused := 1     // コンパイルエラー: declared and not used

	score = score + 5 // 自分自身を使った更新もできます
	fmt.Println("加点後の score:", score)
}
