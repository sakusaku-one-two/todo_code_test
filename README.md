# todo_code_test
コーディングテスト

go 1.24.4


##

実行方法は以下になります。

①docker立ち上げスクリプト
```
cd todo_code_test
make api-run
```

②docker-composeのビルド画面から下記の文章が出るまで待機をお願いします。
```
api_server-1  | テーブルのマイグレーションが完了しました。
```

③別のターミナルを立ち上げて下記のスクリプトを実行するとCLIが立ち上がります。

```
cd todo_code_test
make cli-run
```
##



## ディレクトリ構成（サーバーサイド）

以下の点を工夫いたしました。

1️⃣： 軽量DDDをイメージして作成しました。具体的にはI/Oに関わるコードは  io_infraにまとめ、アプリケーションのロジックはdomainにまとめました。      
2️⃣： またデプロイのことを考えた際、CLIクライアントをGoの同一モジュールで実装することに違和感を感じたので、別のモジュールとして実装しました。 ただbufのサンプルコードだとprotcol bufferのビルドファイルを
3️⃣： そのままdockerビルドしてデプロイできるように環境変数を通してHostとDBとの接続アドレス・ポートを設定できるようにしました。  
4️⃣: sqlboilerのコマンドを実行しなくていいよう、cmd/migrate/up/main.goにmigrateを実行する処理を記述しました。これはsql boilerの公式ドキュメントで見つけたので取り入れました。



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



## GRPC

```

service TodoService {
    rpc CreateTodo(CreateTodoRequest) returns (CreateTodoResponse){};
    rpc GetAllTodo(GetALLRequest) returns(TodoListResponse){};
    rpc FindTodo(stream SearchRequest) returns(stream TodoListResponse){};
    rpc UpdateTodo(UpdateTodoRequest) returns (UpdateTodoResponse){};
    rpc DeleteTodo(DeleteTodoRequest) returns(DeleteTodoResponse){};
}


```

検索する機能(FindTodo)のみ双方向ストリームで実装しました。これは双方向ストリームに興味があったので個人的な学習目的で行いました。




