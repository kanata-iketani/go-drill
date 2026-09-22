# lesson16 ch06 演習: キーを集めてソートして出力（順序が必要なときの定石）

標準入力の 3 行に「名前 点数」が空白区切りで与えられます（名前の重複はありません）。map に登録したあと、名前の昇順に「名前 点数」を 1 行ずつ出力してください。map の range 順は不定なので、キーを集めてソートしてから出力します。

**入力される値**

```text
名前1 点数1
名前2 点数2
名前3 点数3
```

**入力例1**

```text
suzuki 90
sato 80
tanaka 70
```

**期待出力1**

```text
sato 80
suzuki 90
tanaka 70
```

## テストケース

### ケース 1

入力:

```text
suzuki 90
sato 80
tanaka 70
```

期待出力:

```text
sato 80
suzuki 90
tanaka 70
```

### ケース 2

入力:

```text
c 3
a 1
b 2
```

期待出力:

```text
a 1
b 2
c 3
```

### ケース 3

入力:

```text
yuki 5
ken 9
ami 7
```

期待出力:

```text
ami 7
ken 9
yuki 5
```
