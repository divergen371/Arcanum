# PasswordCrypt

PasswordCryptは、ファイルや文字列を安全に暗号化・復号化するためのシンプルなコマンドラインツールです。AES-CBCアルゴリズムを使用して、機密情報を保護します。

## 機能

- ファイルの暗号化と復号化
- 文字列の直接暗号化と復号化
- クロスプラットフォーム対応（Windows、Linux、macOS）
- 複数のCPUアーキテクチャに対応（x86_64, ARM64, ARM）

## インストール

### バイナリをダウンロード

最新のリリースから、お使いのプラットフォームに合ったバイナリをダウンロードしてください。

### ソースからビルド

```bash
# リポジトリをクローン
git clone https://github.com/divergen371/passwordcrypt.git
cd passwordcrypt

# 開発用ビルド
make build

# 本番用ビルド（すべてのプラットフォーム）
make build-prod
```

## 使い方

### 暗号化

ファイルまたは文字列を暗号化：

```bash
# ファイルを暗号化
passwordcrypt encrypt /path/to/file --key <16/24/32バイトの鍵> --iv <16バイトのIV>

# 文字列を暗号化
passwordcrypt encrypt "暗号化したい文字列" --key <16/24/32バイトの鍵> --iv <16バイトのIV>
```

### 復号化

ファイルまたは文字列を復号化：

```bash
# ファイルを復号化
passwordcrypt decrypt /path/to/encrypted/file --key <16/24/32バイトの鍵> --iv <16バイトのIV>

# 文字列を復号化
passwordcrypt decrypt "<暗号化された文字列>" --key <16/24/32バイトの鍵> --iv <16バイトのIV>
```

## 開発

### 前提条件

- Go 1.16以降
- Make

### 開発用ビルド

```bash
make build
```

### テスト実行

```bash
make test
```

### ベンチマークテスト

```bash
make benchmark
```

### コードカバレッジ

```bash
make coverage
```

## ビルドオプション

| コマンド | 説明 |
|---------|------|
| `make build` | 開発用ビルド（デバッグ情報付き） |
| `make build-prod` | 本番用ビルド（すべてのプラットフォーム） |
| `make build-windows` | Windows向けビルド（x86_64, ARM64, ARM） |
| `make build-linux` | Linux向けビルド（x86_64, ARM64, ARM） |
| `make build-darwin` | macOS向けビルド（x86_64, ARM64） |
| `make clean` | ビルド成果物の削除 |
| `make all` | テスト実行後に本番用ビルドを実行 |

## セキュリティ

- 16、24、または32バイトの強力な鍵を使用してください
- 一意の16バイトの初期化ベクトル（IV）を使用してください
- 鍵とIVは安全な方法で管理してください

## ライセンス

MIT License 