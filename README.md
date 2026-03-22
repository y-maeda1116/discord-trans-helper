# Discord Translation Helper

Discordのメッセージをボタン一つで翻訳する補助ボット。

TypeScript版とGo版の両方が提供されています。

## 機能

- Discordの各メッセージに「Translate」コンテキストメニューを表示
- メッセージの内容をDeepL APIを使用して翻訳
- 翻訳結果はボタンを押した本人にだけ表示される（Ephemeral Message）
- 翻訳元の言語を自動判定し、日本語なら英語へ、それ以外なら日本語へ翻訳

## インストール方法

### TypeScript版

```bash
cd ts
npm install
```

### Go版

```bash
# 依存関係のインストール
go mod download
```

## 環境変数の設定

プロジェクトルートに `.env` ファイルを作成します：

```env
DISCORD_TOKEN=your_discord_bot_token_here
DEEPL_AUTH_KEY=your_deepl_auth_key_here
```

### Discord開発者ポータルの設定

1. [Discord Developer Portal](https://discord.com/developers/applications) にアクセス
2. 「New Application」をクリックしてアプリケーションを作成
3. 「Bot」タブから「Add Bot」をクリック
4. 「Reset Token」をクリックしてボットトークンを取得
5. **重要**: MESSAGE CONTENT INTENT を有効化してください
6. 「General Information」タブの「APPLICATION ID」を確認
7. DeepL APIキーを取得し、`.env` に設定

## 実行方法

### TypeScript版

```bash
cd ts

# 開発モード（ホットリロード）
npm run dev

# ビルド
npm run build

# 本番実行
npm start
```

### Go版

```bash
# ビルド
make build

# 実行
make run

# クリーンアップ
make clean

# クロスプラットフォームビルド
make build-win  # Windows
make build-mac  # macOS (arm64)
```

## 使い方

1. ボットがサーバーに参加すると、コマンドが自動登録されます
2. 翻訳したいメッセージを右クリック（またはタップ）
3. 「Apps」 > 「Translate」を選択
4. 翻訳結果があなたにだけ表示されます

## プロジェクト構造

```
.
├── ts/                       # TypeScript版
│   ├── src/
│   │   ├── config/
│   │   │   └── index.ts     # 環境変数設定（Zodバリデーション付き）
│   │   ├── lib/
│   │   │   ├── translator.ts # DeepL翻訳機能
│   │   │   └── commands.ts   # Discordコマンド登録
│   │   └── index.ts          # Discordクライアントメイン
│   ├── package.json
│   └── tsconfig.json
├── cmd/                      # Go版
│   └── app/
│       └── main.go          # メインアプリケーション
├── internal/                # Go版内部パッケージ
│   ├── config/
│   │   └── config.go       # 環境変数設定
│   └── translator/
│       └── translator.go   # DeepL翻訳機能
├── go.mod                  # Goモジュール定義
├── go.sum                  # 依存関係ロック
├── Makefile                # クロスプラットフォームビルド
├── .env.example            # 環境変数テンプレート
└── README.md
```

## ライセンス

MIT
