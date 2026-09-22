// lesson23 ch03 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// トップレベルに置いた関数は、同じ package main なら calc.go など
// 別ファイルへ移しても main からそのまま見えます（import 不要）。
// 「機能ごとにファイルを分ける」が Go の整理の仕方です。
func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func maxOf(nums []int) int {
	m := nums[0]
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(strings.Fields(sc.Text())[0])
	sc.Scan()
	fields := strings.Fields(sc.Text())
	nums := make([]int, n)
	for i := range n {
		nums[i], _ = strconv.Atoi(fields[i])
	}

	fmt.Println(sum(nums))
	fmt.Println(maxOf(nums))
}
