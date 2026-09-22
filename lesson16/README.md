# Lesson 16 マップ

Go では `map[キーの型]値の型` を使います。最大の差は **range で取り出す順序が保証されない** ことで、順序が必要な出力はキーをソートしてから行います。辞書内包表記も存在しません。

## ch01 map リテラルと取得

Python の辞書に相当するのが **map**（キーと値の組を保持するデータ構造）です。map[string]int のようにキーと値の型を指定し、リテラルは map[string]int{"apple": 120} の形で書きます。取得は Python と同じ m[key] です。ただし Python と 2 つ大きな差があります。要素の順序が保証されないこと（後のチャプターで扱います）と、辞書内包表記が存在しないことです。🔴 リテラルの書き方と取得を手が覚えるまで書きましょう。

Python ではこう書いた:

```python
prices = {"apple": 120, "banana": 90}
print(prices["apple"])
```

Go ではこう書く:

```go
prices := map[string]int{"apple": 120, "banana": 90}
fmt.Println(prices["apple"])
```

## ch02 make と追加・更新

空の map は make で作ります。var m map[string]int と宣言しただけの map は **nil map**（実体を持たない map）で、書き込むと実行時エラーになるため、あとから追加していく map は必ず make で作るのが鉄則です。追加と更新はどちらも m[key] = value で、キーがなければ追加、あれば上書きになります。この書き味は Python の d[key] = value と同じです。🔴 「make で作ってから書き込む」流れを手が覚えるまで書きましょう。

Python ではこう書いた:

```python
d = {}
d["apple"] = 100
d["apple"] = 120  # 上書き
```

Go ではこう書く:

```go
m := make(map[string]int)
m["apple"] = 100
m["apple"] = 120 // 上書き
```

## ch03 カンマ ok イディオム v, ok := m[k]

Go の map は、存在しないキーを読んでもエラーにならず、値の型の**ゼロ値**（int なら 0、string なら空文字列）を返します。そのため「値が 0」と「キーがない」を区別できません。区別するには v, ok := m[k] という**カンマ ok イディオム**を使います。ok はキーが存在すれば true になる bool です。Python の in 演算子や get() で行っていた存在チェックは、Go ではすべてこの形で書きます。🔴 手が覚えるまで書きましょう。

Python ではこう書いた:

```python
if "apple" in prices:
    print(prices["apple"])
```

Go ではこう書く:

```go
if v, ok := prices["apple"]; ok {
	fmt.Println(v)
}
```

## ch04 delete と len

map の要素の削除は delete(m, key)、要素数は len(m) で調べます。Python の del d[key] は存在しないキーを指定すると KeyError になりましたが、Go の delete は存在しないキーを渡しても何も起きず、エラーになりません。この仕様のおかげで、削除前の存在チェックは不要です。len はスライスと同じ感覚で使えます。🟡 「delete は失敗しない」という Python との差を理屈で押さえましょう。

Python ではこう書いた:

```python
del prices["apple"]  # キーがないと KeyError
print(len(prices))
```

Go ではこう書く:

```go
delete(prices, "apple") // キーがなくても何も起きない
fmt.Println(len(prices))
```

## ch05 range で全要素（順序は保証されない）

map の全要素は for k, v := range m で取り出せます。ここで最重要の注意点があります。**Go の map は range で取り出す順序が保証されません**。同じプログラムを 2 回動かすと違う順序で出ることさえあります。挿入順を保つ Python 3.7 以降の辞書との明確な差です。そのため range は合計や件数など順序に依存しない集計にだけ使い、順序が問われる出力には使ってはいけません（対処は次チャプター）。🟡 この制約を理屈で理解しましょう。

Python ではこう書いた:

```python
for k, v in d.items():
    print(k, v)  # 挿入順に出る
```

Go ではこう書く:

```go
for k, v := range m {
	fmt.Println(k, v) // 順序は毎回変わりうる
}
```

## ch06 キーを集めてソートして出力（順序が必要なときの定石）

順序を揃えて map を出力したいときの定石は「キーをスライスに集めてソートし、その順に map を引く」です。make([]string, 0, len(m)) で空スライスを用意し、range でキーだけを append、slices.Sort で並べ替えてから for で出力します。順序が保証されない Go の map では、Python の sorted(d) に相当するこの処理を自分で書きます。スキルチェックで map の中身を出力する問題は、ほぼ必ずこの形になります。🔴 手が覚えるまで書きましょう。

Python ではこう書いた:

```python
for k in sorted(d):
    print(k, d[k])
```

Go ではこう書く:

```go
keys := make([]string, 0, len(m))
for k := range m {
	keys = append(keys, k)
}
slices.Sort(keys)
for _, k := range keys {
	fmt.Println(k, m[k])
}
```
