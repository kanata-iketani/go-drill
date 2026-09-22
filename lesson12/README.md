# Lesson 12 for（while相当）

Go に `while` はなく、`for 条件 {}` と無限ループの `for {}` が while の役割を兼ねます。`break` と `continue` は Python と同じ単語です。

## ch01 for 条件 {}（while の代わり）

Go に `while` はありません。繰り返しはすべて `for` が担当し、`for 条件 {}` と書くと Python の `while 条件:` と同じ動きになります。条件に括弧は不要で、ブロックは `{}` で囲みます。条件が false になるまで本体を繰り返すため、本体でカウンタを更新し忘れると止まらなくなる点も while と同じです。なお Python の `i += 1` は Go では `i++` とも書けます。🔴 while の置き換えとして最頻出の形なので、手が覚えるまで書きましょう。

Python ではこう書いた:

```python
i = 1
while i <= 3:
    print(i)
    i += 1
```

Go ではこう書く:

```go
i := 1
for i <= 3 {
	fmt.Println(i)
	i++
}
```

## ch02 for {} 無限ループと break

条件を書かない `for {}` は無限ループになり、Python の `while True:` に相当します。抜けるには `break`（いま実行中のループを即座に終了する文）を使います。単語は Python と同じです。「無限ループの中で処理し、条件を満たしたら break で抜ける」という形は、後で学ぶ標準入力の処理でも使う定番の書き方です。🟡 break が終わらせるのはループ 1 つ分だけ、という範囲を理解しましょう。

Python ではこう書いた:

```python
while True:
    total += 3
    if total >= 12:
        break
```

Go ではこう書く:

```go
for {
	total += 3
	if total >= 12 {
		break
	}
}
```

## ch03 continue（次の繰り返しへ進む）

`continue` は、ループ本体の残りを飛ばして次の繰り返しへ進む文です。意味は Python と同じですが、`for 条件 {}` の形で使うときは注意が必要です。カウンタの更新より前に `continue` すると更新まで飛ばされ、無限ループになります。更新を本体の先頭に書いてから continue するのが安全な形です。🟡 continue がどこへ飛ぶのかを、カウンタ更新の位置とあわせて理解しましょう。

Python ではこう書いた:

```python
i = 0
while i < 5:
    i += 1
    if i == 3:
        continue
    print(i)
```

Go ではこう書く:

```go
i := 0
for i < 5 {
	i++
	if i == 3 {
		continue
	}
	fmt.Println(i)
}
```

## ch04 ループのネスト

ループの中にループを書く**ネスト**（入れ子構造のこと）は、Python と同じ考え方で `for` を重ねるだけです。外側が 1 周する間に、内側は最後まで回ります。内側で `break` や `continue` を使うと、内側のループだけに効きます。Python で二重ループを一気に抜けるにはフラグ変数が必要でしたが、Go にはラベルという目印があり、`break outer` のように外側ごと抜けられます。🟡 いま内側と外側のどちらが回っているのかを追えるようになりましょう。

Python ではこう書いた:

```python
for i in range(1, 3):
    for j in range(1, 4):
        print(i, j)
```

Go ではこう書く:

```go
i := 1
for i <= 2 {
	j := 1
	for j <= 3 {
		fmt.Println(i, j)
		j++
	}
	i++
}
```
