# lesson17 ch03 演習: 和・積・差は自分で書く

標準入力の 1 行目に集合 A の単語、2 行目に集合 B の単語が空白区切りで与えられます（各行の中に重複はありません）。1 行目に積（A にも B にもある単語）、2 行目に差（A にだけある単語）を、それぞれ辞書順にソートして空白区切りで出力してください。

**入力される値**

```text
A の単語1 A の単語2 ...
B の単語1 B の単語2 ...
```

**入力例1**

```text
apple banana orange
banana grape apple
```

**期待出力1**

```text
apple banana
orange
```

## テストケース

### ケース 1

入力:

```text
apple banana orange
banana grape apple
```

期待出力:

```text
apple banana
orange
```

### ケース 2

入力:

```text
a b c d
b d e
```

期待出力:

```text
b d
a c
```

### ケース 3

入力:

```text
x y z
y
```

期待出力:

```text
y
x z
```
