# Discord Translation Helper

Discordのメッセージをボタン一つで翻訳する補助ボット。

## 機能

- Discordの各メッセージに「Translate」コンテキストメニューを表示
- メッセージの内容をDeepL APIを使用して翻訳
- 翻訳結果はボタンを押した本人にだけ表示される（Ephemeral Message）
- 翻訳元の言語を自動判定し、日本語なら英語へ、それ以外なら日本語へ翻訳

## 技術スタック

- TypeScript
- discord.js
- deepl-node
- Node.js >= 20.0.0

## セットアップ

### 1. 依存関係のインストール

```bash
npm install
```

### 2. 環境変数の設定

```bash
cp .env.example .env
```

`.env` ファイルを編集して以下の値を設定します：

```env
DISCORD_TOKEN=your_discord_bot_token_here
DEEPL_AUTH_KEY=your_deepl_auth_key_here
CLIENT_ID=your_discord_client_id_here
```

### 3. Discord開発者ポータルの設定

ボットを正しく動作させるには、以下の手順が必要です：

#### 3.1 ボットの作成とトークンの取得

1. [Discord Developer Portal](https://discord.com/developers/applications) にアクセス
2. 「New Application」をクリックしてアプリケーションを作成
3. 「Bot」タブから「Add Bot」をクリック
4. 「Reset Token」をクリックしてボットトークンを取得し、`.env` の `DISCORD_TOKEN` に設定
5. **重要**: MESSAGE CONTENT INTENT を有効化してください
   - 「Bot」タブの「Privileged Gateway Intents」セクション
   - 「Message Content Intent」をオンにする

#### 3.2 CLIENT_ID の取得

1. 「General Information」タブを開く
2. 「APPLICATION ID」の値を `.env` の `CLIENT_ID` に設定

#### 3.3 DeepL APIキーの取得

1. [DeepL API](https://www.deepl.com/pro-api) に登録
2. APIキーを取得し、`.env` の `DEEPL_AUTH_KEY` に設定

#### 3.4 ボットをサーバーに招待

1. 「OAuth2」 > 「URL Generator」タブを開く
2. 「Scopes」で「bot」を選択
3. 「Bot Permissions」で少なくとも「Send Messages」と「Use Application Commands」を選択
4. 生成されたURLにアクセスしてボットをサーバーに招待

## 実行

### 開発モード（ホットリロード）

```bash
npm run dev
```

### ビルド

```bash
npm run build
```

### 本番実行

```bash
npm start
```

## 使い方

1. ボットがサーバーに参加すると、コマンドが自動登録されます
2. 翻訳したいメッセージを右クリック（またはタップ）
3. 「Apps」 > 「Translate」を選択
4. 翻訳結果があなたにだけ表示されます

## プロジェクト構造

```
.
├── src/
│   ├── config/
│   │   └── index.ts      # 環境変数設定（Zodバリデーション付き）
│   ├── lib/
│   │   ├── translator.ts  # DeepL翻訳機能
│   │   └── commands.ts    # Discordコマンド登録
│   └── index.ts           # Discordクライアントメイン
├── .env.example          # 環境変数テンプレート
├── package.json
└── README.md
```

## ライセンス

MIT
