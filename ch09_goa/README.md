# goa

## 概要

goaを使うことでOpenAPIを使ったWebアプリケーションのAPIサーバーを堅牢にメンテナンスしやすい形で作ることができます。

## OpenAPIと比較対象

### API定義のない開発

[gin](https://gin-gonic.com/ja/) や [echo](https://echo.labstack.com/) のようなWebアプリケーションフレームワークを使う場合、クライアントが使うことができるAPIを試すためのUIや文書を用意するのはサーバーの開発者の責務になります。また、APIの変更と内部の実装の変更を区別することが難しくなりやすいです。

### GraphQL

Facebookが開発し、GitHubなどでも採用されているAPI定義のフォーマットです。定義を元にクライアントのソースコードなどを生成することが可能です。データ構造を定義することで、型のチェックが厳しく、クライアント側が指定した情報を取得するクエリを発行することが可能です。また、QLはQuery Languageを意味しますが、クエリだけでなくデータを操作するAPIも定義可能です。

### OpenAPI

OpenAPIは RESTful なAPIを定義するためのフォーマットです。データ構造にはJSONあるいはYAMLを用います。具体的なHTTP(S)のプロトコルにマッピングしやすい構造になっているのでとてもわかり易いです。また、SwaggerUIなどを使ってAPIの確認や実行することも可能です。

### GraphQL vs OpenAPI

- わかりやすさ・ハードルの低さ ---> OpenAPI
    - HTTP層の言葉だけで扱えるのは非常にわかりやすい
- データ構造の定義 ---> GraphQL
    - 実際のスキーマとの整合性を保証できる（ただしgoaはなら保証可能）
- クライアント側TypeScriptのコードとの親和性 ---> GraphQL
    - [元記事](https://zenn.dev/adwd/scraps/ae6e1d177bd721)
- 開発の安定度・将来性 ---> GraphQL?
    - [OpenAPI Generatorはちょっと大変だった時期もあるみたい](https://ackintosh.github.io/blog/2018/05/12/openapi-generator/)
- ログのわかりやすさ ---> OpenAPI
    - Cloud Loggingなどに `POST /graphql` が並んでいるとログで特定のリクエストを探すことが非常に辛いはず [元記事](https://tech.jxpress.net/entry/graphql-vs-rest)

単純に比較するとGraphQLの方が将来性があって堅牢そうです。しかしデータ構造の定義と実装との整合性のチェックは後述のgoaが保証できるので、ハードルの低さを重視するならばOpenAPIが僅差で勝ちます。

## goaを選ぶ理由

[openapi-generator](https://github.com/OpenAPITools/openapi-generator)でも [goaと同じようなことができます](https://tech.pepabo.com/2022/09/09/go-middleware/)。しかしAPI定義をYAMLあるいはJSONで書く必要があるopenapi-generatorの場合、APIの量やその内部の構造が複雑になった場合、非常に大きな定義ファイルを扱う必要が生じます。

goaの場合は、goaが提供するgo言語でのDSLを使ってAPI定義を書くことができるので、似たような表現を関数としてまとめることも可能ですし、ファイルを分割していくことも可能です。

またgoaはHTTPだけでなくgRPCもサポートしますが、そのビジネスロジックはHTTPやgRPCなどのトランスポートに依存しません。つまりHTTPリクエストやレスポンスをgoaによってマッピングされた抽象的な入出力データとして扱うので、ビジネスロジックとHTTPでの通信を分けて考えることが可能です（関心事の分離）。
