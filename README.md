# vacca-note-server
![example workflow](https://github.com/Okaki030/vacca-note-api/actions/workflows/deploy.yml/badge.svg)

「ワクチン接種体験共有サービス Vacca note」 のAPI

## プロジェクトの概要
[Vacca note ーコロナワクチン接種体験共有サービスー](https://indecisive-berry-33f.notion.site/Vacca-note-e390c4ad207d44209535d5a94b18d2cd)

## 使用技術

- Go 1.17
    - フレームワーク: github.com/gorilla/mux
    - ロガー: go.uber.org/zap
- MySQL 5.7.22
- Docker 20.10.12
- docker-compose 1.29.2

## ローカル環境構築手順

### 準備

```
# リポジトリのClone
$ git clone https://github.com/Okaki030/vacca-note-server.git

# dbの作成
$ docker-compose build
$ docker network create vacca-note
$ docker-compose up -d

# DB作成
$ docker-compose exec mysql mysql -uroot -proot -e'create database vacca_note_db_local;'

# DBマイグレーション
$ docker-compose exec server sql-migrate up
```

### 実行

```
$ docker-compose up
```

ホットリロードのライブラリとして [Air](https://github.com/cosmtrek/air) を利用しているため、ソースコードの修正を検知して実行アプリケーションが自動更新される。

## デプロイ
mainブランチにソースコードがpushされると自動で本番環境にリリースされる。

ここでは手動デプロイの手順を記載する。

```
# ECRにコンテナイメージをpush
$ aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin ***********.dkr.ecr.ap-northeast-1.amazonaws.com
$ docker build -t vacca-note-app .
$ docker tag vacca-note-app:latest ***********.dkr.ecr.ap-northeast-1.amazonaws.com/vacca-note-app:latest
$ docker push ***********.dkr.ecr.ap-northeast-1.amazonaws.com/vacca-note-app:latest

# Fargateに反映
$ SYSTEM_NAME=vacca-note IMAGE_TAG=latest AWS_REGION=ap-northeast-1 ecspresso deploy --config=ecspresso/config.yaml
```

## Tips
### 本番環境のDBに接続する

```
# ECSタスクへの接続をポートフォワーディングする
$ ecspresso exec --port-forward --local-port 3306 --port 3306 --config=ecspresso/config.yaml

# ローカルから本番環境のDBに接続する
$ mysql -h 127.0.0.1 -u vaccanote -p
```

### Fargateコンテナに入る

```
$ ecspresso exec --config=ecspresso/config.yaml
```

### Fargateコンテナを一時的に実行する

```
$ SYSTEM_NAME=vacca-note IMAGE_TAG=latest AWS_REGION=ap-northeast-1 ecspresso run --config=ecspresso/config.yaml --overrides '{"containerOverrides": [{"name":"app", "command": ["go","version"]}]}'
```

## ER図
[vacca-note | DrawSQL](https://drawsql.app/vacca-note/diagrams/vacca-note#)
