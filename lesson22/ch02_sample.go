// lesson22 ch02: 構造体の埋め込み（フィールドとメソッドの昇格）
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Person のメソッドです
func (p Person) Greet() string {
	return fmt.Sprintf("こんにちは、%s です", p.Name)
}

// フィールド名を書かずに型名だけを書くと「埋め込み」になります
type Employee struct {
	Person
	Company string
}

func main() {
	e := Employee{
		Person:  Person{Name: "田中", Age: 30},
		Company: "ペイザ",
	}
	// 内側のフィールドとメソッドが「昇格」して直接使えます
	fmt.Println(e.Name)    // e.Person.Name と書かなくてよい
	fmt.Println(e.Greet()) // Person のメソッドもそのまま呼べる
	fmt.Println(e.Company)
}
