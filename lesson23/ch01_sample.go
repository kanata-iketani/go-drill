// lesson23 ch01: package main と func main が入口
// Python の if __name__ == "__main__": に相当するガードはありません。
package main

import "fmt"

// main 以外の関数は、呼ばれるまで実行されません
func greet(name string) string {
	return "こんにちは、" + name + " さん"
}

func main() {
	// 実行は必ずここから始まります（プログラム唯一の入口）
	fmt.Println(greet("Gopher"))
	fmt.Println("main が終わるとプログラムも終わります")
}
