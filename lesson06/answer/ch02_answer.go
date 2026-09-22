// lesson06 ch02 解答
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Atoi は (int, error) を返すため、2 変数で受けて err を必ず確認します。
	// このチェックを毎回書く習慣が、この先の Go コードの基本形になります。
	a, err := strconv.Atoi("120")
	if err != nil {
		fmt.Println("変換に失敗しました")
		return
	}
	b, err := strconv.Atoi("80")
	if err != nil {
		fmt.Println("変換に失敗しました")
		return
	}
	fmt.Println(a + b)
}
