// lesson03 ch04: 演算子の優先順位
package main

import "fmt"

func main() {
	// * / % は + - より先に計算されます（Python と同じ）
	fmt.Println(100 + 200*3) // 700
	fmt.Println((100+200)*3 == 900)

	// 括弧で順番を変えられます
	fmt.Println((100 + 200) * 3) // 900

	// 迷ったら括弧を付けて意図を明示するのが読みやすさのコツです
	total := (500 + 300) / 2
	fmt.Println("平均:", total)
}
