// lesson11 ch01: bool 型と比較演算子
package main

import "fmt"

func main() {
	// bool 型の値は小文字の true / false です
	isSunny := true
	fmt.Println(isSunny)

	age := 18
	// 比較演算子の結果は bool 型になります
	fmt.Println(age >= 20)
	fmt.Println(age == 18)

	// 結果は変数に入れて使い回せます
	canVote := age >= 18
	fmt.Println("投票できる:", canVote)
}
