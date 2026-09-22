// lesson24 ch02: errors.New で作り、raise の代わりに return する
package main

import (
	"errors"
	"fmt"
)

// 失敗しうる関数は (結果, error) の 2 値を返します
func divide(a, b int) (int, error) {
	if b == 0 {
		// raise ではなく return でエラーを返します
		return 0, errors.New("0 では割れません")
	}
	return a / b, nil // 成功時はエラーの位置に nil
}

func main() {
	if result, err := divide(10, 2); err == nil {
		fmt.Println(result) // 5
	}
	if _, err := divide(7, 0); err != nil {
		fmt.Println(err) // 0 では割れません
	}
}
