# A Tour of Go の復習問題

## Q1. [Goの基礎](https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/02.2.html) の `make, new操作`

図にで示されている `Point` `Circle` `Rect` を struct を使って定義してください。
また、以下のデータを new して生成し、 `fmt.Printf("%#v\n", ...)` を使って出力してください。

1. Px: 100, Py: 50 を持つPoint のデータのポインタ p1
2. Point: p1, radius: 30 を持つCircle のデータのポインタ c1
3. Point: p1, width: 20, length: 10 を持つRectのデータのポインタ r1


## Q2. [フローと関数](https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/02.3.html) の `値渡しと参照渡し` を参考に以下の関数を定義してください

1. `func NewCircle(p *Point, r int) *Circle`
    引数で渡された p を Point, r を radius として持つ 新たな `Circle` のポインタを返す関数
1. `func ExpandCircle(c *Circle, dr int)`
    引数 c で渡された *Circle の参照する Circle の radius に dr を足す関数
1. 上の２つの関数で引数や戻り値にポインタを使わなかったらどうなるのか説明してください。
1. これらの関数のテストあるいは動作を確認できるようなmain関数を作成してください。
1. defer をどのような場合に使うのか例を示して説明してください(ネットで調べても可)。

## Q3. [オブジェクト指向](https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/02.5.html)

1. この章をすべて読んでください
1. Q2で作成した  `ExpandCircle` を `*Circle` のメソッド `Expand` として作り直してください
1. `*Rect` に面積を求めるメソッド `Area() int` を追加してください
1. BoxList を参考に 最も大きな面積の `Rect` (のポインタ) を返すメソッド `Biggest` を持つ `RectList` を `[]*Rect` から作成してください。
   (ヒント Biggest のレシーバーは RectListのポインタを使いません)

## Q4. [interface](https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/02.6.html)

1. この章をすべて読んでください
1. 面積を求めるメソッド `Area` を持つinterface `Shape` を定義してください
1. `Rect` と `Circle` を `Shape` として使えるように変更してください
    円の面積は `円周率 * 半径 * 半径` で求められます。円周率は [math.Pi](https://pkg.go.dev/math#pkg-constants) で求められます。
1. 最も面積が大きな `Shape` を返すメソッド `Biggest` を持つ `Shapes` を `[]Shape` から作成してください。
1. 引数dxとdy で指定された差分だけ移動する `MoveBy` メソッドを `Shape` に追加し `Rect` と `Circle` に実装してください。
1. `Shapes` に その要素全てに対して `MoveBy` 呼び出すメソッド `MoveBy` を実装してください
