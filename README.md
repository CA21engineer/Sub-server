# Sub-server

hack01-ios のサーバー

## エンドポイント

`35.221.100.76:18080`

## ローカル動作確認

```bash
$ sh ./run-local.sh
```

## gRPC Code Gen

```bash
$ sh ./grpc-code-gen/code-gen-go.sh ./protobuf ./apiServer/pb
```

## CI/CD 関連

### ssh 鍵作成

`$ ssh-keygen -t rsa -f my-ssh-key -C [任意のsshユーザーネーム]`

### 設定すべき環境変数

- GOOGLE_PROJECT_ID
- GOOGLE_COMPUTE_REGION

  例: asia-northeast1

- GOOGLE_COMPUTE_ZONE

  例: asia-northeast1-a

- GOOGLE_SERVICE_KEY

  サービスアカンウトを base64 エンコードした文字列
  `base64 -i [.json file path]`

- SSH_USERNAME

  ssh ユーザーネーム

- SSH_KEY

  ssh 秘密鍵を base64 エンコードした文字列
  `base64 -i my-ssh-key`

- SSH_KEY_PUB

  ssh 公開鍵を base64 エンコードした文字列
  `base64 -i my-ssh-key.pub`

- SSH_HOST

  ssh 接続する GCE サーバーの外部 IP
  terraform でインスタンス作成した後判明する外部 IP

- SSH_PORT

  ssh 接続のポート
  空いていればなんでもよい。基本は 22 だがセキュリティー上変更した方がいい

## er 図

![スクリーンショット 2020-06-08 18 27 52](https://user-images.githubusercontent.com/37885842/84014807-dd962200-a9b5-11ea-9067-0b402ad969e9.png)

## ディレクトリ構成について

apiServer 配下に go の api ロジックが全て入っている

```
apiServer
|
├── adopter
├── form
├── models
├── pb
├── responses
└── service
```

- adopter  
  db から引っ張ってきたデータを gRPC の struct に変換する

- form  
  バリデーションとかを扱う

- models  
  db からデータを引っ張ってきたり作成したりする

- pb  
  gRPC で自動作成されるファイルをおく

- response  
  error レスポンスなど固定のレスポンスを入れておく

- service  
  Router みたいな役割を持つファイルをおく
