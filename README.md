
-----

# todo\_code\_test

このプロジェクトは、Go 1.24.4を使用しています。

## 実行方法

プロジェクトの実行はDockerを利用します。以下の手順で簡単にアプリケーションを起動できます。

### 1\. Dockerコンテナの起動

`todo_code_test` ディレクトリに移動し、以下のコマンドを実行してDockerコンテナを起動します。

```bash
make api-run
```

Docker Composeのビルド画面で\*\*`テーブルのマイグレーションが完了しました。`\*\*というメッセージが表示されるまでお待ちください。

### 2\. CLIの実行

別のターミナルを立ち上げ、同様に `todo_code_test` ディレクトリに移動してから、以下のコマンドを実行してください。

```bash
make cli-run
```

このコマンドでCLIクライアントが起動し、アプリケーションの操作が可能になります。

-----

## ディレクトリ構成（サーバーサイド）

### 工夫した点

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
    `sqlboiler`の公式ドキュメントを参考に、`cmd/migrate/up/main.go`にマイグレーション処理を組み込みました。これにより、`sqlboiler`のコマンドを手動で実行する必要がなくなりました。

### ディレクトリツリー

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

## gRPCサービス定義

gRPCサービスは以下のように定義されています。

```protobuf
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


