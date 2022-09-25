# testifyを使う

## 概要

JSONの章では、JSONを使ったCLIアプリケーションのテストを作成しました。
しかしこのテストは少し読みにくい部分もあります。これを [testify](https://github.com/stretchr/testify) を使ってメンテナンスしやすいコードに変更しましょう。また、テストだけでなく全体を見直してコードをわかりやすくしましょう。

## :exclamation: ch05_testify の作成

1. `ch04_json` ディレクトリをコピーして `ch05_testify` ディレクトリを作成します
    - すでに `ch05_testify` ディレクトリが存在する場合はこのREADME.md以外のファイルをコピーして上書きしてください
    - `ch04_json` ディレクトリは [ch04_json/end タグ](https://github.com/akm/second_tour_of_go/releases/tag/ch04_json%2Fend) からソースコードをダウンロード・解凍して取得することもできます
2. `ch05_testify/go.mod` をエディタで編集し、モジュールパス中の `ch04_json` を `ch05_testify` に変更してください


## :exclamation: testifyの導入

Goの標準ライブラリ [testing](https://pkg.go.dev/testing) は最低限の機能しか持っていないので、ややこしいテストを行おうとするとテストのコードも複雑になります。
それを解消するべく簡単にテストを作成できるようにするためのライブラリが [testify](https://github.com/stretchr/testify) です。以下のコマンドを実行して
インストールしてください。

```
cd path/to/ch05_testify
go get github.com/stretchr/testify
```

## :question: assertとrequireを使うように変更

[testify](https://github.com/stretchr/testify) のREADME.md と、以下のパッケージのドキュメントを読んでください。

- [github.com/stretchr/testify/assert](https://pkg.go.dev/github.com/stretchr/testify/assert)
- [github.com/stretchr/testify/require](https://pkg.go.dev/github.com/stretchr/testify/require)

その上で [estimate_test.go](./estimate_test.go) を `assert` パッケージ や `require` パッケージを使うように変更してください。


## :question: まるごと比較するように変更

構造体の各フィールドを比較する関数 `assertResponseItem` を削除し、代わりに構造体のポインタやそのポインタのスライスの中身を比較するように変更します。

1. `assert.Len` と `assertResponseItem` を3回呼び出している部分を 一つの `assert.Equal` を使うように変更してください
2. `TestNewResponseItem` で `assertResponseItem` を呼び出している部分を `assert.Equal` を使うように変更してください
3. `assertResponseItem` を削除してください

以上を行った上で、これまでと同じ動作をするように変更してください。
