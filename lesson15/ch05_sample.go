// lesson15 ch05: スキルチェック定型テンプレート総まとめ
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// ここからが丸暗記するテンプレートです
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024) // 長い行に備えて上限を拡大

	// 1 行目: 整数 N
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	// 2 行目: 空白区切りの N 個の整数を Fields + Atoi で読む
	sc.Scan()
	f := strings.Fields(sc.Text())
	sum := 0
	for i := range n {
		v, _ := strconv.Atoi(f[i])
		sum += v
	}
	fmt.Println(sum)
}
