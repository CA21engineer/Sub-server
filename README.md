# Sub-server

hack01-ios のサーバー

## エンドポイント
`35.221.100.76:18080`

[モニタリングダッシュボード](https://github.com/CA21engineer/Subs-server/wiki/%E3%83%A2%E3%83%8B%E3%82%BF%E3%83%AA%E3%83%B3%E3%82%B0%E3%83%80%E3%83%83%E3%82%B7%E3%83%A5%E3%83%9C%E3%83%BC%E3%83%89)

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
  
- FIREBASE_SERVICE_KEY
    `base64 -i firebase.json`

## er 図

![スクリーンショット 2020-06-08 18 27 52](https://user-images.githubusercontent.com/37885842/84014807-dd962200-a9b5-11ea-9067-0b402ad969e9.png)

## ディレクトリ構成について

apiServer 配下に go の api ロジックが全て入っている

```
apiServer
|
├── adopter
├── models
├── pb
├── responses
└── service
```

- adopter  
  db から引っ張ってきたデータを gRPC の struct に変換する

- models  
  db からデータを引っ張ってきたり作成したりする

- pb  
  gRPC で自動作成されるファイルをおく

- response  
  error レスポンスなど固定のレスポンスを入れておく

- service  
  Router みたいな役割を持つファイルをおく

## Tips：立てたサーバに対してリクエストを送りたい！

[grpcurl](https://github.com/fullstorydev/grpcurl)を使う（ちゃんと見たほうがいいかも w）

### インストール方法

```shell
# grpcurl のインストール
$ brew install grpcurl

# インストール確認
$ grpcurl -help
```

### 実際にリクエストを送ってみる

```shell

# grpcurl -plaintext localhost:18080 subscription.SubscriptionService/任意のメソッド名

# 以下例
$ grpcurl -plaintext localhost:18080 subscription.SubscriptionService/GetSubscriptions
{
  "iconImage": [
    {
      "iconId": "1",
      "iconUri": "https://images-fe.ssl-images-amazon.com/images/I/411j1k1u9yL.png"
    },
    {
      "iconId": "2",
      "iconUri": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAOEAAADhCAMAAAAJbSJIAAAAeFBMVEX/AAD/////u7v/6+v/0ND/aWn/mZn/zc3/9fX/h4f/5ub/+Pj/tbX/t7f/oKD/ODj/Kir/3d3/wMD/MzP/UVH/1dX/yMj/qan/lJT/TU3/gID/Fxf/PT3/eHj/Dg7/jY3/ZWX/Q0P/XV3/IiL/cXH/q6v/X1//eXnZjuaVAAAE20lEQVR4nO2d63aiQAyAOyKgXAVF0Lbrtbvv/4YLZ9didUDQQEjI9689RyffwcCQub0p7rxhB9A5YkgfMaSPGNJHDOkjhvQRQ/qIIX3EkD5iSB8xpI8Y0kcM6SOG9BFD+oghfcSQPmJIHzGkjxi+jHEhml4Tff+/6wBeMXTdWU6SrD0/DQIry3Z709yEYbh4e45F/tmNae53WWYFQep76yQpmnDd3gxdw4mi2LPm27O5OZzenzRpy/vpsDHP27nlxVHkGO18Gxo6dmx9HjfLvpzqeF9ujp9WbDtQhnZqPvur65qFmdovGjrBCtviIaug/mLWGdp77Ogbsq+7ktWG0fAvX8kqam+4xQ66JduWhk6IHXFrwop01BuuscN9inVzQw871ifxmhr62JE+jd/MMMGO8wWSJoYudpQvcd9nvTc8Ygf5EsfHhlTvMhfu7ja3hu4JO8QXOd3+Tm8NA+wIXyZ4YIgdHwD1htSzsMCrNTxghwfAoc7QwI4OBKPGcIIdHAiTGsMzdnAgnGsMP7CDA+Gj2tDBjg0Ip9LQxg4NCLvS0MIODQir0vATOzQgPisNedxobm4114Yz7MjAmFUY8ujRFBgVhjRriDrWFYYpdmBgpBWGXB4WPx8X14Y8eqUF5wrDJXZgYCwrDLHjAmSshnwehz8eiFeGlMcrbkm0hjF0M4hTOGKtIfig2iRGm37jaw3Ba6UTpTLo72yIpzWcQzdT1LwMnG7EXGtoQjfzr6qXbKC/twGm1hB84PBSt4x/QX/zQ45aQ/A3/LIy23uf/kNrCH5zv6o9Gz2XgBa9GyoV9ZqOWsMZeLb8HD9QcY/Dy79mGIZ9pqPWMAJv5s5QGb3NB4yQDPub06kznIK3ojPM07GXDvkU0VCpoIcegM4QfuCpylDNwLvAd5TDTyiGSjldp6POEL7iXWOYN9dtZa+sepeGX+Ct1BoqlXaZjl8aQ/ia/gPDTtOxrOtjGubp2NlUT53hDryVx4Z5OnY0KrvTGML/ZJoY5vnfSTqWZQx8w27ScVCGeTr+Bm9bZwjf7W9smKcj9KTIcpFQaQheamtjqJQHm45lsW0whsBpMkhDZezh2h6mYf4GB5aOQzVUygcazBmuoVIwd5zhGnK/htzzkP29lPvzsIc+Da9+6R+NIf93C3k/bM/Q3vH512n419r410v517z5j1vwH3viP37IfwyY/zg+/7kY/OfT8J8TxX9eG/+5iSOYX8p/jjD/ed785+rzX2/Bf80M/3VP/Neu8V9/yH8NKf91wCMw5L8en/+eCvz3xeC/twn//Wn4PBCr9hjiv08U/72+RrBfG5fHRfWee/z3TeS/9yWTW03N/qVMeqZ1e9Dy30eYR6+mbi9o/vt5s9iT3a815FDJuDW6+Zv/2Qj8z7cgn4kPzygZwTkzIzgriPT4RaPznij33XSzsHSGZO82jc9do1pXbHF23gjOP1T8z7BUIziHVPE/S7aA+3nA/68k6zOdy4vJ+Vzuaxifra73dYvD65Nk7flpEFhZttub5ioMw2d/1Yv8syvT3O+yzAqC1PfsJCmacNs5wRk2wrgQTa+Jvv/fdQCdG6IjhvQRQ/qIIX3EkD5iSB8xpI8Y0kcM6SOG9BFD+oghfcSQPmJIHzGkjxjSRwzpI4b0EUP6iCF9/gKCNIdvAhrz/wAAAABJRU5ErkJggg=="
    },
    {
      "iconId": "3",
      "iconUri": "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcRuX7izxLGFnXQ7k79lGWEew3njHyI2NCmkq3-y_RN3An1lS7cj\u0026usqp=CAU"
    }
  ]
}

```
