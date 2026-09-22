// lesson10 ch01 解答
package main

import "fmt"

func main() {
	score := 85

	// 条件に () は付けず、本体が 1 行でも {} を省略できないのが Go の規則です。
	// 「合格」は 80 以上のときだけ、「判定終了」は if の外なので必ず出力されます。
	if score >= 80 {
		fmt.Println("合格")
	}
	fmt.Println("判定終了")
}
