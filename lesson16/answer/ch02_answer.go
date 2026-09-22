// lesson16 ch02 解答
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

	// 書き込んでいく map は make で作るのが鉄則です（nil map には書き込めない）。
	// m[key] = value は、キーがなければ追加、あれば上書きになるため、
	// 「再登場したら上書き」は特別な処理なしで実現できます。
	m := make(map[string]int)
	for range 3 {
		sc.Scan()
		f := strings.Fields(sc.Text())
		price, _ := strconv.Atoi(f[1])
		m[f[0]] = price
	}

	sc.Scan()
	fmt.Println(m[sc.Text()])
}
