# Lesson 17 集合の代替

Go に set 型は存在せず、`map[T]bool`（または `map[T]struct{}`）で代用します。和・積・差などの集合演算子もないため、for を使って自分で書きます。

## ch01 Go に set 型は存在しない（map で代用する）

Go に set 型は存在しません。Python の**集合**（重複しない要素の集まり）に相当する処理は、map で代用します。定番は map[T]bool で、要素をキーに、値に true を入れます。メモリを節約したいときは値を空の構造体にする map[T]struct{} という書き方もありますが、まずは読み書きが素直な map[T]bool で十分です。所属判定は set[x] と引くだけで、ないキーはゼロ値 false になるため、そのまま if の条件に使えます。🟡 「set は map で作る」と理屈で押さえましょう。

Python ではこう書いた:

```python
s = {"apple", "banana"}
print("apple" in s)
```

Go ではこう書く:

```go
set := map[string]bool{"apple": true, "banana": true}
fmt.Println(set["apple"])
```

## ch02 存在判定と重複除去

set 代わりの map が最も活躍するのが存在判定と重複除去です。「すでに見たか」を seen という名前の map[T]bool に記録しながらスライスを走査し、初めて見た要素だけを結果に append します。set 型がない Go ではこれが重複除去の定石で、Python の set(lst) と違って**元の順序を保ったまま**取り除けるという利点もあります。スキルチェックでも実務でも頻出のパターンです。🔴 seen マップを使う形を手が覚えるまで書きましょう。

Python ではこう書いた:

```python
unique = list(dict.fromkeys(words))  # 順序を保つ重複除去
```

Go ではこう書く:

```go
seen := map[string]bool{}
unique := []string{}
for _, w := range words {
	if !seen[w] {
		seen[w] = true
		unique = append(unique, w)
	}
}
```

## ch03 和・積・差は自分で書く

Python の集合演算子 `|`（和）`&`（積）`-`（差）は Go にはないため、set 型がない Go では for で自分で書きます。発想は 1 つで、「一方を map にしておき、もう一方を走査して、含まれるかどうかで振り分ける」だけです。積は「A の要素のうち B にあるもの」、差は「A の要素のうち B にないもの」、和は「A と B を両方登録した map のキー」になります。出力の順序が必要なら、前レッスンの定石どおりソートします。🟡 演算子がなくても for で書ける、と理屈で理解しましょう。

Python ではこう書いた:

```python
inter = a & b
diff = a - b
```

Go ではこう書く:

```go
inter := []string{}
for _, w := range a {
	if setB[w] { // B にもあれば積
		inter = append(inter, w)
	}
}
```
