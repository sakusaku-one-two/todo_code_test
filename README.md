# todo_code_test
コーディングテスト

```shtodo01/├── cmd/ # アプリケーションエントリーポイント│ ├── client/ # CLI client│ └── server/ # gRPC server├── internal/│ ├── db/ # データベース接続とリポジトリ│ │ └── generated/ # SQLBoilerによる自動生成コード│ ├── handler/ # gRPCハンドラー│ └── models/ # 内部データモデル├── proto/ # Protocol Buffers定義│ └── todo/│ └── v1/ # バージョン管理されたAPI定義├── migrations/ # データベースマイグレーション├── scripts/ # セットアップスクリプト├── Dockerfile # Docker設定├── docker-compose.yml # Docker Compose設定├── Makefile # ビルドタスク└── README.md # ドキュメント```