# todo_code_test
コーディングテスト

```shtodo01/├── cmd/ # アプリケーションエントリーポイント│ ├── client/ # CLI client│ └── server/ # gRPC server├── internal/│ ├── db/ # データベース接続とリポジトリ│ │ └── generated/ # SQLBoilerによる自動生成コード│ ├── handler/ # gRPCハンドラー│ └── models/ # 内部データモデル├── proto/ # Protocol Buffers定義│ └── todo/│ └── v1/ # バージョン管理されたAPI定義├── migrations/ # データベースマイグレーション├── scripts/ # セットアップスクリプト├── Dockerfile # Docker設定├── docker-compose.yml # Docker Compose設定├── Makefile # ビルドタスク└── README.md # ドキュメント```


```

MYSQL_ROOT_PASSWORD=root_password
MYSQL_DATABASE=todo_database
MYSQL_USER=mysql_user
MYSQL_PASSWORD=todo_database_password
MYSQL_HOSTNAME=db


migrate -path ./migration -database "mysql://mysql_user:todo_database_password@tcp(localhost:3306)/todo_database" up

 grpcurl \
    -protoset <(buf build -o -) -plaintext \
    -d '{"request": "Jane"}' \
    localhost:8080 todo.v1.TodoService/GetAll

     grpcurl     -protoset <(buf build -o -) -plaintext     -d '{"query": "self create todo"}'     localhost:8080 proto.todo.v1.TodoService/FindTodo

```