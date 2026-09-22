// lesson01 ch02: Go プログラムの骨格
// package main + import + func main の 3 点セット
package main

// import は使う機能（パッケージ）を取り込む宣言です
import "fmt"

// プログラムは必ず func main から始まります
func main() {
	fmt.Println("1: main 関数の先頭から実行されます")
	fmt.Println("2: 上から順に実行されます")
	fmt.Println("3: main 関数が終わるとプログラムも終わります")
}
