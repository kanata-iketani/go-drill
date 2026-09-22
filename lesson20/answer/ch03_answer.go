// lesson20 ch03 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ...int で受けると、呼び出し側は個数を気にせず渡せます。
// 中では nums はただの []int なので、range で普通に走査できます。
func largest(nums ...int) int {
	best := nums[0] // 入力は 1 個以上という前提
	for _, n := range nums[1:] {
		if n > best {
			best = n
		}
	}
	return best
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fields := strings.Fields(sc.Text())
	nums := make([]int, len(fields))
	for i, f := range fields {
		nums[i], _ = strconv.Atoi(f)
	}
	// スライスは nums... とばらして可変長引数に渡します（Python の *nums）
	fmt.Println(largest(nums...))
}
