# Sub-server
hack01-iosのサーバー

## ローカル動作確認
```bash
$ sh ./run-local.sh 
```

## gRPC Code Gen
```bash
$ sh ./grpc-code-gen/code-gen-go.sh ./protobuf ./apiServer/pb
```

## CI/CD関連

### ssh鍵作成

`$ ssh-keygen -t rsa -f my-ssh-key -C [任意のsshユーザーネーム]`

### 設定すべき環境変数

- GOOGLE_PROJECT_ID
- GOOGLE_COMPUTE_REGION

    例: asia-northeast1

- GOOGLE_COMPUTE_ZONE

    例: asia-northeast1-a

- GOOGLE_SERVICE_KEY

    サービスアカンウトをbase64エンコードした文字列
    `base64 -i [.json file path]`

- SSH_USERNAME

    sshユーザーネーム

- SSH_KEY

    ssh秘密鍵をbase64エンコードした文字列
    `base64 -i my-ssh-key`

- SSH_KEY_PUB

    ssh公開鍵をbase64エンコードした文字列
    `base64 -i my-ssh-key.pub`

- SSH_HOST

    ssh接続するGCEサーバーの外部IP
    terraformでインスタンス作成した後判明する外部IP

- SSH_PORT

    ssh接続のポート
    空いていればなんでもよい。基本は22だがセキュリティー上変更した方がいい
