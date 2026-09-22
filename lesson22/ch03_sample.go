// lesson22 ch03: インターフェースは「振る舞いの契約」
package main

import "fmt"

// Speak() string を持つ型はすべて Speaker になります（暗黙実装）
type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + ": ワン"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return c.Name + ": ニャー"
}

func main() {
	// 契約を満たす型なら、同じスライスにまとめて扱えます
	speakers := []Speaker{Dog{Name: "ポチ"}, Cat{Name: "タマ"}}
	for _, s := range speakers {
		fmt.Println(s.Speak())
	}
}
