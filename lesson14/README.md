# Lesson 14 内包表記の代替

Go にリスト内包表記は存在せず、「空スライス + for + append」の形で書きます。3 行になりますが上から順に読めるのが Go 流で、繰り返し使う変換は関数に切り出します。

## ch01 Go にリスト内包表記は存在しない（for + append で書く）

Go にリスト内包表記は存在しません。**内包表記**（リストを 1 行で生成する Python の構文）で書いていた処理は、Go では「空スライスを作る → for で回す → append で追加する」という形で書きます。3 行になるが読める、というのが Go の考え方で、処理の流れが上から順にそのまま追えるため、これが標準の書き方です。🟡 1 行で書けない理由を探すより、この 3 行の型を理屈から理解しましょう。

Python ではこう書いた:

```python
squares = [x * x for x in range(5)]
```

Go ではこう書く:

```go
squares := []int{}
for x := range 5 {
	squares = append(squares, x*x)
}
```

## ch02 条件付き内包表記の代替（if でフィルタしながら append）

条件付き内包表記（if で要素を絞り込みながらリストを作る Python の書き方）も Go には存在しません。for の中に if を入れ、条件を満たしたときだけ append します。フィルタと変換を同時に行いたい場合も、if の中で加工してから append するだけです。3 行が 4 行になるが読める、という点は前チャプターと同じで、どの行が何をしているか一目で追えます。🟡 「for + if + append」の組み合わせを理屈で理解しましょう。

Python ではこう書いた:

```python
evens = [x for x in nums if x % 2 == 0]
```

Go ではこう書く:

```go
evens := []int{}
for _, x := range nums {
	if x%2 == 0 {
		evens = append(evens, x)
	}
}
```

## ch03 関数に切り出すのが Go 流

リスト内包表記がない Go で、for + append のかたまりが何度も登場するなら、関数に切り出すのが Go 流です。Python では内包表記を 1 行で使い回しましたが、Go では「スライスを受け取り、新しいスライスを返す関数」を定義して名前を付けます。呼び出す側は 1 行になり、内包表記と同じ読みやすさに戻ります。処理に名前が付くぶん、意図はむしろ伝わりやすくなります。🟡 変換処理を関数にまとめる発想を理屈で理解しましょう。

Python ではこう書いた:

```python
doubled = [x * 2 for x in nums]
```

Go ではこう書く:

```go
func double(nums []int) []int {
	result := []int{}
	for _, x := range nums {
		result = append(result, x*2)
	}
	return result
}

// 呼び出す側は 1 行に戻る
doubled := double(nums)
```
