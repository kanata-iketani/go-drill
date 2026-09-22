// lesson16 ch05: range で全要素（順序は保証されない）
package main

import "fmt"

func main() {
	scores := map[string]int{"sato": 80, "suzuki": 90, "tanaka": 70}

	// range の順序は保証されません。実行のたびに変わることもあるため、
	// このサンプルでは順序に依存しない「集計」だけを行います
	sum := 0
	count := 0
	for _, v := range scores {
		sum += v
		count++
	}

	// 合計や件数は足す順番に関係なく同じ結果になるので安全です
	fmt.Println("人数:", count)
	fmt.Println("合計:", sum)
}
