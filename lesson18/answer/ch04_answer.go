// lesson18 ch04 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(strings.Fields(sc.Text())[0])

	// この場でしか使わない小さな処理なので、名前付き関数にせず
	// 無名関数を変数に代入しました。lambda x: x * x の Go 版ですが、
	// 必要なら中に複数行のロジックも書けます。
	square := func(x int) int {
		return x * x
	}
	fmt.Println(square(n))
	fmt.Println(square(n + 1))
}
