// lesson05 ch04 解答
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 同じ線を 2 回使うので、Repeat の結果を一度変数に入れます。
	// "-" を 15 個手打ちするより、数を変えたいときに 1 か所で済むのが利点です。
	line := strings.Repeat("-", 15)

	fmt.Println(line)
	fmt.Println("ここまで読んだ")
	fmt.Println(line)
}
