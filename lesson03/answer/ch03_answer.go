// lesson03 ch03 解答
package main

import (
	"fmt"
	"math"
)

func main() {
	// Go には ** がないので累乗は math.Pow を使います。
	// math.Pow の戻り値は float64（81 でも表示は 81 になるが型は小数）なので、
	// 整数として扱いたいことを明示するため int() で変換しています。
	fmt.Println(int(math.Pow(3, 4)))
}
