// lesson15 ch02: 読んだ文字列を数値へ（strconv.Atoi）
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	// Text() は文字列なので、計算する前に Atoi で int へ変換します。
	// スキルチェックでは err を _ で捨てるのが定石です
	n, _ := strconv.Atoi(sc.Text())

	fmt.Println("2倍:", n*2)
	fmt.Println("10足す:", n+10)
}
