// lesson19 ch01 解答
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
	sc.Scan()
	fields := strings.Fields(sc.Text())
	nums := make([]int, len(fields))
	for i, f := range fields {
		nums[i], _ = strconv.Atoi(f)
	}

	// for の中で宣言した変数はループを出ると消えるため、
	// ループ後も使う total はブロックの外で宣言しておきます。
	total := 0
	for _, n := range nums {
		total += n // ここは代入なので外の total が更新される
	}
	fmt.Println(total)
}
