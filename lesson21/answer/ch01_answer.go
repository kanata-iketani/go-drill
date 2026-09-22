// lesson21 ch01 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 名前と年齢は常にセットで扱うデータなので、構造体にまとめます。
// class と違いメソッドや __init__ は持たず、まずは「型付きの入れ物」です。
type Person struct {
	Name string
	Age  int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fields := strings.Fields(sc.Text())
	name := fields[0]
	age, _ := strconv.Atoi(fields[1])

	// フィールド名: 値 で初期化。順番に頼らないので読みやすく安全です
	p := Person{Name: name, Age: age}
	fmt.Printf("%sは%d歳です\n", p.Name, p.Age)
}
