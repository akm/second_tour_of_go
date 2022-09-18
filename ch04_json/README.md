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


## :question: summary サブコマンド

1. 指定されたJSONファイル( 例えば [people.json](./people.json) ) を読み込んで、人数と平均年齢を出力するサブコマンド summary を追加してください。
    - 出力例
        ```
        5 people, average age: 30
        ```
2. 以下の場合にどのように振る舞うのかをしらべてください
    1. フィールド名がマッチしないJSONファイルを指定した場合
    2. JSON形式じゃないファイルが指定した場合

### ヒント

- [json.Unmarshal](https://pkg.go.dev/encoding/json#Unmarshal)

## :question: スライス型の拡張

以下のような`People` 型を定義して、そのメソッド `AverageAge` で平均年齢を求めるようにサブコマンド summary を変更してください。
Peopleの要素数が0の場合は0を返すものとします。このテストも作成してください。順番は初めてのテストで紹介したやり方を思い出してください。

```golang
type Person []*Person
```

## :question: estimate サブコマンド

指定された商品JSONファイルと見積もりリクエストJSONファイルを読み込んで、見積もり結果JSONを標準出力に出力するサブコマンド `estimate` を作成してください。軽減税率の対象の商品の消費税は `8%` 、対象外の商品の消費税は `10%` とします。
消費税は各商品毎に求め、端数は切り捨てください（問題を簡単にするため）。
商品JSONに含まれない商品名が指定された場合は商品が見つからないという旨のエラーにしてください。

### 商品JSONファイル

```json
{
  "Apple": {"unit_price": 200, "reduced_rate": false},
  "Orange": {"unit_price": 120, "reduced_rate": true},
  "Banana": {"unit_price": 250, "reduced_rate": true},
  "Kiwi Fruit": {"unit_price": 100, "reduced_rate": true},
  "Lemon": {"unit_price": 150, "reduced_rate": false}
}
```

商品名のキーに対して、unit_price(価格)とreduced_rate(軽減税率対象)のフィールドを持つオブジェクトを値とするマップです。

### 見積もりリクエストJSONファイル

```json
{
  "client_name": "John Doe",
  "items": [
    {
      "product_name": "Apple",
      "quantity": 3
    },
    {
      "product_name": "Orange",
      "quantity": 4
    },
    {
      "product_name": "Banana",
      "quantity": 2
    }
  ]
}
```

client_name(顧客名)とitems(明細)を持つオブジェクトです。
itemsの要素は、product_name(商品名)とquantity(数量)のフィールドを持つオブジェクトです。

### 見積もり結果JSON

```json
{
  "client_name": "John Doe",
  "estimated_at": "2022-09-18T16:27:20.467167+09:00",
  "subtotal": 1580,
  "tax": 138,
  "total": 1718,
  "items": [
    {
      "product_name": "Apple",
      "quantity": 3,
      "subtotal": 600,
      "tax_rate": 10,
      "tax": 60
    },
    {
      "product_name": "Orange",
      "quantity": 4,
      "subtotal": 480,
      "tax_rate": 8,
      "tax": 38
    },
    {
      "product_name": "Banana",
      "quantity": 2,
      "subtotal": 500,
      "tax_rate": 8,
      "tax": 40
    }
  ]
}
```

以下のフィールドを持つオブジェクトです。

- client_name(顧客名)
- estimated_at(見積もり日時)
- subtotal(小計、税抜)
- tax(消費税)
- total(合計金額)
- items(明細の配列)

明細は以下のフィールドを持つオブジェクトです。

- product_name(商品名)
- quantity(数量)
- subtotal(小計、税抜)
- tax_rate(税率、%)
- tax(税額)

### ヒント

- 商品のデータについては、商品名をキー、(unit_priceとreduced_rateに相当するフィールドを持つ構造体)を値とする `map` を使うと作りやすいと思います
- 型の名前の候補
    - Product 製品
    - Estimate 見積もり
    - Request リクエスト
    - Response 結果
    - Item 要素
    - Map マップ


### チャレンジ

1. どのような型を作るのかリストアップする
2. 何がどのメソッドを呼ぶのかを考える
3. 構造体やスライスの型を定義
4. メソッドを仮実装
5. テストを作成
6. サブコマンド estimate を実装
