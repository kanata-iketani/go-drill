// lesson24 ch04 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type AgeError struct {
	Age int
}

// Error() string を実装すると、この構造体は error として返せます。
// errors.New と違い、失敗した値（Age）をフィールドとして持ち運べるので、
// メッセージの組み立てや呼び出し側での詳細な処理に使えます。
func (e *AgeError) Error() string {
	return fmt.Sprintf("%d歳は成人ではありません", e.Age)
}

func checkAge(age int) error {
	if age < 20 {
		return &AgeError{Age: age}
	}
	return nil
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	age, _ := strconv.Atoi(sc.Text())

	if err := checkAge(age); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("成人です")
}
