# lesson06 ch03 演習: strconv.ParseFloat（文字列→float64）

文字列 `"1.5"` と `"2.25"` をそれぞれ `strconv.ParseFloat` で float64 に変換し、合計を `fmt.Printf` の書式 `%.2f` を使って小数第 2 位まで出力してください。`if err != nil` のチェックも書くこと。`strconv` を import に追加してください。

**期待出力**

```text
3.75
```
