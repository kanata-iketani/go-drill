// lesson12 ch04 解答
package main

import "fmt"

func main() {
	// 外側の i が 1 周する間に内側の j が 1〜3 まで回るので、
	// 出力は 1*1 から 3*3 まで 9 行になります。
	// j はネストの内側で毎回 1 に戻して宣言し直すのがポイントです。
	i := 1
	for i <= 3 {
		j := 1
		for j <= 3 {
			fmt.Printf("%d*%d=%d\n", i, j, i*j)
			j++
		}
		i++
	}
}
