// lesson24 ch04: Error() string を実装すると独自エラー型になる
package main

import "fmt"

// 追加情報（失敗した値）を持てる独自エラー型です
type AgeError struct {
	Age int
}

// このメソッドを実装した型は error として返せます
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
	if err := checkAge(15); err != nil {
		fmt.Println(err) // 15歳は成人ではありません
	}
	if err := checkAge(30); err == nil {
		fmt.Println("30歳は成人です")
	}
}
