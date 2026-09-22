# Lesson 22 埋め込みとインターフェース

Go に継承は存在しないため、`class Child(Parent)` に相当する構文はありません。代わりに構造体の埋め込み（has-a）とインターフェース（振る舞いの契約）で同じ目的を達成します。

## ch01 Go に継承はない（is-a ではなく has-a で考える）

Go に**継承**は存在しません。Python の `class Dog(Animal)` のような親子関係（is-a、「Dog は Animal である」という関係）は書けず、**has-a**（ある型が別の型をフィールドとして持つこと）で考えます。継承ツリーを設計する代わりに、部品となる構造体を組み合わせて型を作り、共通の振る舞いは埋め込みとインターフェースで表現します。🟡 まずは「is-a ではなく has-a で考える」という発想の転換を理解しましょう。

Python ではこう書いた:

```python
class Dog(Animal):  # Dog は Animal である (is-a)
    pass
```

Go ではこう書く:

```go
type Dog struct {
	Animal Animal // Dog は Animal を持つ (has-a)
	Breed  string
}
```

## ch02 構造体の埋め込み（フィールドとメソッドの昇格）

フィールド名を書かずに型名だけを書くと**埋め込み**になります。埋め込んだ内側の型のフィールドとメソッドは外側の型に**昇格**（外側の変数から直接呼べるようになること）し、`e.Person.Name` を `e.Name` と短く書けます。Python の継承でメソッドを引き継いだのと似た見た目になりますが、実体はあくまで has-a であり、親子関係ではありません。🟡 昇格の仕組みを理解しましょう。

Python ではこう書いた:

```python
class Employee(Person):
    pass  # Person のメソッドをそのまま使える

e.greet()
```

Go ではこう書く:

```go
type Employee struct {
	Person  // フィールド名なし = 埋め込み
	Company string
}

e.Greet() // Person のメソッドが昇格して直接呼べる
```

## ch03 インターフェース（振る舞いの契約と暗黙実装）

**インターフェース**は「この型はこのメソッドを持つ」という振る舞いの契約です。Go には `implements` のような宣言がなく、契約どおりのメソッドを持つ型は自動的にそのインターフェースを満たします（**暗黙実装**）。Python のダックタイピング（同名メソッドがあれば動くという流儀）に近い柔らかさを、コンパイル時の型検査つきで実現します。継承の代わりに「同じ振る舞いの型をまとめて扱う」のが役割です。🟡 契約という考え方を理解しましょう。

Python ではこう書いた:

```python
class Dog:
    def speak(self):  # 同名メソッドがあれば動く
        return "ワン"
```

Go ではこう書く:

```go
type Speaker interface {
	Speak() string // この契約を満たす型は自動的に Speaker
}
```
