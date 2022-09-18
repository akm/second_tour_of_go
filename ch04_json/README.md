# 初めての JSON

## 予習してきてください

1. 以下を調べてください（文章にすることを強く推奨）
    - JSON とは何か
    - JSON で使える型の種類
    - CSVとの違い
1. https://pkg.go.dev/encoding/json を読んでおいてください

## :exclamation: セットアップ

1. ターミナルあるいはコマンドプロンプトで `ch04_json` に移動
2. `go mod init github.com/akm/second_tour_of_go/ch04_json` を実行
   - `github.com/akm` の部分は他の文字列に変更しても OK です。
     - Github 等で管理するのであれば、自分の環境に合わせて変更してください

以下、実行形式のファイル名は `ch04_json` とします。
該当しないサブコマンドが指定された場合にはヘルプを表示してください。
何かエラーが起きた場合とヘルプを表示する際には、原因となるエラーを標準エラー出力に出力し、終了コードを1としてください。

## :question: example サブコマンド

1. 以下のJSON形式の文字列を標準出力に出力するサブコマンド `example` を `Person` というstruct型を定義して作成してください
    ```json
    {"first_name":"Blake","last_name":"Serilda","birthday":"1989-07-10","age":33}
    ```
2. 上のJSON形式の文字列を読みやすく整形した文字列を標準出力に出力するようにサブコマンド `example` を変更してください
    ```json
    {
      "first_name": "Blake",
      "last_name": "Serilda",
      "birthday": "1989-07-10",
      "age": 33
    }
    ```
3. `Person` 型のスライスを使って以下のように出力するようにサブコマンド `example` を変更してください
    ```json
    [
      {
        "first_name": "Blake",
        "last_name": "Serilda",
        "birthday": "1989-07-10",
        "age": 33
      },
      {
        "first_name": "Libbie",
        "last_name": "Drisko",
        "birthday": "1998-06-15",
        "age": 24
      },
      {
        "first_name": "Anestassia",
        "last_name": "Truc",
        "birthday": "1973-04-02",
        "age": 48
      }
    ]
    ```

### ヒント

- [json.Marshal](https://pkg.go.dev/encoding/json#Marshal)
- [json.MarshalIndent](https://pkg.go.dev/encoding/json#MarshalIndent)
