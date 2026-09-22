// lesson21 ch02 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rect struct {
	W, H int
}

// 読むだけなので値レシーバ。r にはコピーが入ります
func (r Rect) Area() int {
	return r.W * r.H
}

// フィールドを書き換えて呼び出し元に反映したいので、ポインタレシーバにします。
// 値レシーバのままだとコピーを書き換えるだけで、Scale の結果が消えてしまいます。
func (r *Rect) Scale(n int) {
	r.W *= n
	r.H *= n
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fields := strings.Fields(sc.Text())
	w, _ := strconv.Atoi(fields[0])
	h, _ := strconv.Atoi(fields[1])
	n, _ := strconv.Atoi(fields[2])

	r := Rect{W: w, H: h}
	fmt.Println(r.Area())
	r.Scale(n)
	fmt.Println(r.Area())
}
