Python の `input()` 相当が Go では `bufio.Scanner` を使った 4 行になります。スキルチェックはこの定型が書ければ戦えるため、理屈と合わせて丸暗記してください。

## スキルチェック用テンプレート

```go
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
	sc.Buffer(make([]byte, 1024*1024), 1024*1024) // 長い行に備えて上限を拡大

	// 1 行目: 整数 N
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	// 2 行目: 空白区切りの N 個の整数
	sc.Scan()
	f := strings.Fields(sc.Text())
	nums := make([]int, n)
	for i := range n {
		nums[i], _ = strconv.Atoi(f[i])
	}

	fmt.Println(nums)
}
```
