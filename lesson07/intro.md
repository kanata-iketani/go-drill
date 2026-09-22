Go でリストに当たるのは**スライス**です。`pop` `remove` `del` に当たる関数はなく、末尾削除 `s[:len(s)-1]` や `slices.Delete` のようにスライス操作で書きます。`sorted` の代わりは `slices.Clone` でコピーしてから `slices.Sort` です。
