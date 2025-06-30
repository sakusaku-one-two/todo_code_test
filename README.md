
-----

# todo\_code\_test

このプロジェクトは、Go 1.24.4を使用しています。

使用技術  
・ connect-go (gRPCの実装で使用)  
・ sqlboiler/go-migrate (migrateとモデルの作成で使用)  
・ buf        (protocol bufferの作成で使用)    
・ docker-compose (mysqlとサーバーサイドのコンテナのために使用)  

## セットアップ手順

プロジェクトの実行はサーバーサイドでDockerを利用します。以下の手順で簡単にアプリケーションを起動できます。
```
mkdir test_todo
cd test_todo
git init
git clone https://github.com/sakusaku-one-two/todo_code_test.git
cd todo_code_test
```

### 1\. Dockerコンテナの起動

`todo_code_test` ディレクトリに移動し、以下のコマンドを実行してDockerコンテナを起動します。

```bash
make api-run
```

Docker Composeのビルド画面で`テーブルのマイグレーションが完了しました。`というメッセージが表示されるまでお待ちください。

### 2\. CLIの実行

別のターミナルを立ち上げ、同様に `todo_code_test` ディレクトリに移動してから、以下のコマンドを実行してください。  
※こちらはDockerに含めることができませんでした。ローカルにイントールされたGolangを利用しています。

```bash
make cli-run
```

このコマンドでCLIクライアントが起動し、アプリケーションの操作が可能になります。

-----
## 操作方法


### 作成
作成と入力してください。
```
'削除'・'作成'・'更新'・'検索'・'取得'・'終了'のどれかを入力してください。 ctl+cで強制終了
 作成

```

各種項目を入力してください。
```
タイトルを入力してください。
新規作成タイトルを入力 
説明を入力してください。
新規作成説明
期限の月(1~12)を半角数値で指定してください。
7
期限の日を半角数値で入力してください。
1
```
作成に成功するとTodoが表示されます。
```
1 ___________________________________________________
タイトル 新規作成タイトルを入力
説明　　 新規作成説明
期限　　 2025年7月1日
状態　　 未完了
作成日　 2025年6月30日
________________________________________________________
```

### 取得（全てのTodo）
```
'削除'・'作成'・'更新'・'検索'・'取得'・'終了'のどれかを入力してください。 ctl+cで強制終了
取得
1 ___________________________________________________
タイトル 新規作成タイトルを入力
説明　　 新規作成説明
期限　　 2025年7月1日
状態　　 未完了
作成日　 2025年6月30日
________________________________________________________
```
　
### 更新
```
'削除'・'作成'・'更新'・'検索'・'取得'・'終了'のどれかを入力してください。 ctl+cで強制終了
更新
```
削除されてないTodoが一覧として表示されます。左上にある番号で選択します。
```
1 ___________________________________________________
タイトル 新規作成タイトルを入力
説明　　 新規作成説明
期限　　 2025年7月1日
状態　　 未完了
作成日　 2025年6月30日
________________________________________________________
左上のindex番号を指定してください。
1
更新する項目を数値で選んでください。タイトル:1 説明:2 未完了/完了:3
3
完了は1未完了は2と数値で入力してください。
1
```
更新に問題がなければ更新後のTodoが表示されます。
```
1 ___________________________________________________
タイトル 新規作成タイトルを入力
説明　　 新規作成説明
期限　　 2025年7月1日
状態　　 完了
作成日　 2025年6月30日
________________________________________________________
```


## 検索
```
'削除'・'作成'・'更新'・'検索'・'取得'・'終了'のどれかを入力してください。 ctl+cで強制終了
検索
```
検索するタイトルを入力してください。
```
検索するタイトルを入力してください。（部分検索可能）
新
```
検索結果が表示されます。
```
1 ___________________________________________________
タイトル 新規作成タイトルを入力
説明　　 新規作成説明
期限　　 2025年7月1日
状態　　 完了
作成日　 2025年6月30日
________________________________________________________
```
検索を終了するか問われます。
```
検索終了しますか？ 'yes'で終了
yes
```

## 削除
```
'削除'・'作成'・'更新'・'検索'・'取得'・'終了'のどれかを入力してください。 ctl+cで強制終了
削除
```
削除対象となるTodo一覧が表示されます。左上のIndex番号から指定してください。
```
1 ___________________________________________________
タイトル 新規作成タイトルを入力
説明　　 新規作成説明
期限　　 2025年7月1日
状態　　 完了
作成日　 2025年6月30日
________________________________________________________
削除する対象のTodo左上のindex番号を指定してください。
1
```
（論理削除を行います）
```
削除に成功しました

```

# 工夫した点

このプロジェクトは、以下の点を念頭に置いて設計・実装されています。

1.  **軽量DDD（ドメイン駆動設計）**
    I/O処理（インフラ層）とアプリケーションのロジックを明確に分離しました。
      * **`io_infra`**: データベース接続やgRPCサービスといったI/Oに関連するコードをまとめています。
      * **`domain`**: アプリケーションのコアロジック（エンティティ、リポジトリ、ユースケース）を配置しています。
2.  **モジュール分割**
    デプロイの観点から、サーバーとCLIクライアントを別のGoモジュールとして開発しました。これにより、それぞれの独立性を高めています。
3.  **環境変数による設定**
    Dockerビルド後のデプロイを考慮し、データベースのホストやポートといった接続情報は環境変数から設定できるようにしています。
4.  **自動マイグレーション**
    `go-migrate`の公式ドキュメントを参考に、`cmd/migrate/up/main.go`にマイグレーション処理を組み込みました。これにより、`sqlboiler`のコマンドを手動で実行する必要がなくなりました。

### ディレクトリツリー(サーバー)
```
api_server
  ├── Dockerfile
  ├── cmd
  │   ├── api
  │   │   └── main.go
  │   └── migrate
  │       ├── down
  │       │   └── main.go
  │       └── up
  │           └── main.go
  ├── go.mod
  ├── go.sum
  ├── internal
  │   ├── domain
  │   │   ├── entitys
  │   │   │   └── todo_entity
  │   │   │       └── todo.go
  │   │   ├── repository
  │   │   │   ├── Irepository.go
  │   │   │   └── todo_repository
  │   │   │       ├── repo_test.go
  │   │   │       └── todo_repository.go
  │   │   ├── use_cases
  │   │   │   └── todo_usecase
  │   │   │       ├── todo_usecase.go
  │   │   │       ├── todo_usecase_helper.go
  │   │   │       └── todo_usecase_test.go
  │   │   └── values
  │   │       ├── IValue.go
  │   │       └── todo_values
  │   │           ├── description.go
  │   │           ├── limit.go
  │   │           ├── status.go
  │   │           ├── task_id.go
  │   │           ├── title.go
  │   │           └── todo_values_test.go
  │   ├── grpc_gen
  │   │   └── todo
  │   │       └── v1
  │   │           ├── todo.pb.go
  │   │           └── todov1connect
  │   │               └── todo.connect.go
  │   ├── handler
  │   │   └── todo_handler.go
  │   └── io_infra
  │       ├── config
  │       │   ├── my_sql_config
  │       │   │   ├── config_test.go
  │       │   │   └── database_config.go
  │       │   └── server_config
  │       │       └── server_config.go
  │       ├── database
  │       │   ├── driver
  │       │   │   ├── db_driver_test.go
  │       │   │   └── driver.go
  │       │   ├── migration
  │       │   │   ├── 000001_create_todo.down.sql
  │       │   │   └── 000001_create_todo.up.sql
  │       │   ├── models
  │       │   │   └── .... sqlboilerによる自動生成テンプレート
  │       │   └── sqlboiler.toml
  │       ├── grpc_services
  │       │   ├── todo_service.go
  │       │   └── todo_service_test.go
  │       └── server
  │           └── server.go
  └── util
      ├── env.go
      ├── env_test.go
      ├── nil_checker.go
      └── nil_checker_test.go
```

-----

## ディレクトリツリー（クライアント）

```
client
  ├── Dockerfile
  ├── cmd
  │   └── cli
  │       └── main.go
  ├── go.mod
  ├── go.sum
  └── internal
      └── grpc_gen
          └── todo
              └── v1
                  ├── todo.pb.go
                  └── todov1connect
                      └── todo.connect.go
```


## gRPCサービス定義

gRPCサービスは以下のように定義されています。

```protobuf

// proto/todo/v1/todo.proto

service TodoService {
    // 新しいTODOを作成
    rpc CreateTodo(CreateTodoRequest) returns (CreateTodoResponse){};

    // 全てのTODOを取得
    rpc GetAllTodo(GetALLRequest) returns(TodoListResponse){};

    // TODOを検索 (双方向ストリーム)
    rpc FindTodo(stream SearchRequest) returns(stream TodoListResponse){};

    // 既存のTODOを更新
    rpc UpdateTodo(UpdateTodoRequest) returns (UpdateTodoResponse){};

    // TODOを削除
    rpc DeleteTodo(DeleteTodoRequest) returns(DeleteTodoResponse){};
}
```

`FindTodo`機能は**双方向ストリーム**で実装しています。これは、個人的な学習目的で双方向ストリームに興味があったため、実験的に取り入れました。


