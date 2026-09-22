// lesson18 ch02 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 計算結果は出力せずに return で返します。
// 「計算する関数」と「表示する場所」を分けておくと、
// 同じ関数をテストや別の計算に使い回せるからです。
func area(w, h int) int {
	return w * h
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fields := strings.Fields(sc.Text())
	w, _ := strconv.Atoi(fields[0])
	h, _ := strconv.Atoi(fields[1])
	fmt.Println(area(w, h))
}
