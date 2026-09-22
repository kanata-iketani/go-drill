// lesson22 ch01 解答
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	Name string
}

type Employee struct {
	Person  Person
	Company string
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	name := sc.Text()
	sc.Scan()
	company := sc.Text()

	// 継承の代わりに、Employee が Person をフィールドとして持ちます (has-a)。
	// 「Employee は Person である」ではなく「Employee は Person を持つ」と
	// 考えるのが Go の設計スタイルです。
	e := Employee{Person: Person{Name: name}, Company: company}
	fmt.Println(e.Person.Name, "は", e.Company, "の社員です")
}
