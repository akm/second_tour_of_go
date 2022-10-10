# HTTPサーバー

## 予習してきてください

### URL

[MDN ウェブ開発を学ぶ - URL とは？](https://developer.mozilla.org/ja/docs/Learn/Common_questions/What_is_a_URL) と
[MDN Guides - What is a URL?](https://developer.mozilla.org/ja/docs/Learn/Common_questions/What_is_a_URL) を読んで以下の質問に答えてください。

1. URLを構成する６つの要素を英語で言うと何と言いますか？
2. 以下のURLをそれぞれ６つの要素に分類してください
    1. http://localhost:8080/hello
    2. http://localhost/
    3. https://developer.mozilla.org/en-US/search?q=URL&page=2
    4. https://github.com/akm/second_tour_of_go/pull/2/files?diff=split&w=1#diff-78728faaaba530c64071edff8afd0d30a8d1103bfd75e27c8c34a690526e831d

### HTTP

[MDN 開発者向けのウェブ技術 - HTTPガイド](https://developer.mozilla.org/ja/docs/Web/HTTP/Basics_of_HTTP) と [とほほのWWW入門 - HTTP入門](https://www.tohoho-web.com/ex/http.htm) から調べて以下の質問に答えてください。特にHTTPの用語は英語をベースに使うので[MDN 開発者向けのウェブ技術 - HTTP メッセージ](https://developer.mozilla.org/ja/docs/Web/HTTP/Messages) と [MDN References - HTTP Messages](https://developer.mozilla.org/en-US/docs/Web/HTTP/Messages) を見比べてください。

1. 「HTTPは○○である」。○○に該当する単語は？
2. HTTPリクエストを構成する３つの要素を英語では何と言いますか？
3. HTTPレスポンスを構成する３つの要素を英語では何と言いますか？
4. 以下のそれぞれの場合に、一般的にどのHTTPメソッドを使いますか？
    1. サーバーから情報を取得するとき
    2. サーバーにデータを追加するとき
    3. サーバーのデータを削除するとき
5. 以下のそれぞれの場合に使用するHTTPステータスコードは何ですか？
    1. 正常にレスポンスを返した場合
    2. クライアントが送信したリクエストが不正な場合
    3. 対象のデータが存在しない場合
    4. 認証できなかった場合
    5. 権限がなかった場合
    6. サーバーで何らかのエラーが発生した場合
6. curl コマンドがインストールされ使用できることを確認すしてください
    - `curl -v` でヘルプメッセージが表示されればOKです
    - MacやWindows 10以降ならば標準でインストールされているはずです
    - 万が一使えない場合は、使えるようにしておいてください

## :exclamation: Hello, world!

1. `ch06_http_server` ディレクトリがなければ作成し `go mod init github.com/akm/second_tour_of_go/ch06_http_server` を実行してください
2. [http.ListenAndServe の Example](https://pkg.go.dev/net/http#example-ListenAndServe) を写経（あるいはコピー）して`main.go` を作成してください。
3. `go run .` を実行し、別のターミナルでcurlコマンドを以下のように色々試してみてください
    - URL
        - http://localhost:8080/
        - http://localhost:8080/hello
        - http://localhost:8080/hoge
        - http://localhost:8080/hellohello
    - `-v` or `-i` オプション
    - `-X GET` , `-X POST` など

## :question: Echoサーバー

1. main.goをリクエストの内容を以下のようにレスポンスのボディとして返すEchoサーバーに作り変えてください。エラーが発生した場合はステータスコード `500 Internal Server Error` のレスポンスを返してください
   ```
   --start line--
   method: (メソッド)
   scheme: (URLのscheme)
   domain: (URLのdomain)
   port: (URLのport)
   path: (URLのpath)
   parameters: (URLのparameters のキー毎に以下を出力）
     (パラメータのキー): (パラメータの値をカンマで区切った文字列)
   anchor: (URLのanchor)

   --headers--
   (ヘッダーのキー): (ヘッダーの値をカンマで区切ったもの)
   ....

   --body--
   (ボディの内容)
   ```
    - ヒント
        - `fmt.Fprintf(w, "--start line--\n")`
        - `start line` のデータの中には取得できないものがいくつかあります。
2. curlを使って色々なリクエストを送ってみてください。Hello, world! のときに指定したオプションに加え、以下のオプションも指定してみてください。
    - URLのパラメータ `?baz=300`
    - URLのパラメータ `?key=foo&key=bar`
    - URLのアンカー `abc`
    - `-H "X-Foo: 100"`
    - `-H "X-Foo: 100" -H "X-Foo: 200"`
    - `--data '{"foo":100,"bar":200}'`

## :question: 四則演算API

以下の仕様を満たすようにmain.goを実装してください

演算 | メソッド | パス | 入力 | レスポンスボディ
-----|--------|-------|-------------|----------------
加算  |  GET   | /add  | パラメータ `a` , `b` | 演算結果 を表す文字列
加算  |  GET   | /subtract/(a)/(b)  | パスの中の `a` , `b` | 演算結果 を表す文字列
減算  |  POST  | /multiply | ヘッダー `X-VALUE-A` , `X-VALUE-B` | `{"result": (演算結果)}`
減算  |  POST  | /divide | `{"a":(数値): "b":(数値)}` というJSONのリクエストボディ | `{"result": (演算結果)}`

レスポンスのステータスコードは `200` `400` `404` `405` `500` のいずれかを使ってください。

(余裕があれば)パラメータに `debug` が指定された際に echoサーバーのレスポンスの内容を標準出力に出すようにしてみてください。

後でリファクタリングするので、あまりリファクタリングはしないようにお願いします

## :question: 四則演算APIのテスト

main.go をテスト可能な形に変更した上で、テストを追加してください。
またテストを作成、実行して見つかった問題点も修正してください。

- ヒント
    - `testify`
    - `net/http/httptest` パッケージ
