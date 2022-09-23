# testifyを使う

## 概要

JSONの章では、JSONを使ったCLIアプリケーションのテストを作成しました。
しかしこのテストは少し読みにくい部分もあります。これを [testify](https://github.com/stretchr/testify) を使ってメンテナンスしやすいコードに変更しましょう。また、テストだけでなく全体を見直してコードをわかりやすくしましょう。

## :exclamation: ch05_testify の作成

1. `ch04_json` ディレクトリをコピーして `ch05_testify` ディレクトリを作成します
    - すでに `ch05_testify` ディレクトリが存在する場合はこのREADME.md以外のファイルをコピーして上書きしてください
    - `ch04_json` ディレクトリは [ch04_json/end タグ](https://github.com/akm/second_tour_of_go/releases/tag/ch04_json%2Fend) からソースコードをダウンロード・解凍して取得することもできます
2. `ch05_testify/go.mod` をエディタで編集し、モジュールパス中の `ch04_json` を `ch05_testify` に変更してください
