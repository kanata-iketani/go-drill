// lesson13 ch03: _ で値を捨てる
package main

import "fmt"

func main() {
	scores := []int{80, 65, 92}

	// 値だけ使うときは、1 つ目のインデックスを _ で捨てます
	for _, v := range scores {
		fmt.Println(v)
	}

	// インデックスだけなら 2 つ目を丸ごと省略できます
	for i := range scores {
		fmt.Println("index:", i)
	}

	// i, v を受け取って使わないと unused のコンパイルエラーになります
	total := 0
	for _, v := range scores {
		total += v
	}
	fmt.Println("合計:", total)
}
