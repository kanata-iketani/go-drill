// lesson22 ch01: Go に継承はない（has-a で考える）
package main

import "fmt"

// Animal は「名前と鳴き声」を持つ部品です
type Animal struct {
	Name string
	Cry  string
}

// Dog は Animal を継承するのではなく、フィールドとして「持ち」ます (has-a)
type Dog struct {
	Animal Animal
	Breed  string
}

func main() {
	d := Dog{
		Animal: Animal{Name: "ポチ", Cry: "ワン"},
		Breed:  "柴犬",
	}
	// 中の Animal にはフィールド名を経由してアクセスします
	fmt.Println(d.Animal.Name, "は", d.Breed, "です")
	fmt.Println(d.Animal.Name, "は", d.Animal.Cry, "と鳴きます")
}
