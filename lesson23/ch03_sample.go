// lesson23 ch03: go mod init と複数ファイル
// go mod init example.com/hello で go.mod が作られ、モジュールになります。
// 下の関数たちは、同じ package main の別ファイル（例: calc.go）に
// 移しても import なしでそのまま呼べます。実行は go run . です。
package main

import "fmt"

// 本来は calc.go など別ファイルに置ける関数です
func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func average(nums []int) int {
	return sum(nums) / len(nums)
}

func main() {
	scores := []int{80, 90, 70, 100}
	fmt.Println("合計:", sum(scores))
	fmt.Println("平均:", average(scores))
}
