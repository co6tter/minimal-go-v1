# minimal-go-v1

## Overview

Go言語の基礎から並行処理まで段階的に学習するためのサンプルプロジェクトです。7つのアプリケーションを通じて、Go言語の主要な機能を体系的に学習できます。

## Tech Stack

- Go 1.23
- Dev Container対応（VSCode）

## Setup

```bash
# リポジトリのクローン
git clone <repository-url>
cd minimal-go-v1

# Dev Container環境を使用する場合
# VSCodeでリポジトリを開き、「Reopen in Container」を選択
```

## Usage

各アプリケーションは独立して実行可能です：

```bash
# app01の実行例
go run cmd/app01/main.go

# 他のアプリケーションも同様に実行
go run cmd/app02/main.go
go run cmd/app03/main.go
# ...
```

## Directory Structure

```
.
├── cmd/
│   ├── app01/    # 基本文法：変数、型変換、switch、forループ
│   ├── app02/    # 制御構造：if文、forループ、真偽値
│   ├── app03/    # 関数：複数戻り値、高階関数、ジェネリクス
│   ├── app04/    # データ構造：配列、スライス、マップ、構造体
│   ├── app05/    # ポインタ：参照渡し、構造体のポインタ、nil
│   ├── app06/    # オブジェクト指向：カスタム型、メソッド、インタフェース、型アサーション
│   └── app07/    # 並行処理：ゴルーチン、チャネル、WaitGroup、Mutex
└── .devcontainer/  # Dev Container設定
```

### 各アプリケーションの学習内容

- **app01**: 変数宣言、型変換、定数、switch文、forループ
- **app02**: 条件分岐（if/else）、forループ、真偽値の扱い
- **app03**: 関数定義、複数戻り値、関数型、高階関数、ジェネリクス
- **app04**: 配列とスライス、マップ、構造体、range構文
- **app05**: ポインタの基本、参照渡し、nil処理
- **app06**: カスタム型、メソッド、埋め込みフィールド、インタフェース、型アサーション
- **app07**: ゴルーチン、チャネル、sync.WaitGroup、sync.Mutex、select文

## License

This repository is for personal/private use only.
