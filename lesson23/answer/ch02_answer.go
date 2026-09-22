// lesson23 ch02 解答
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fields := strings.Fields(sc.Text())
	a, _ := strconv.Atoi(fields[0])
	b, _ := strconv.Atoi(fields[1])

	// math の関数は float64 を受け取るため、int から変換して渡します。
	// 結果も float64 で返るので、整数として表示するために int に戻します。
	// 呼び方は常に「パッケージ名.関数名」です（sqrt 単体では呼べません）。
	fmt.Println(int(math.Pow(float64(a), float64(b))))
	fmt.Println(int(math.Max(float64(a), float64(b))))
}
