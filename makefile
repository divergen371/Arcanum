# Makefile

# ターゲット名
TARGET = passwordcrypt

# ソースファイル
SRC = ./cmd/passwordcrypt/main.go

# バージョン番号
VERSION = 0.1.0-alpha
GIT_COMMIT = $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# ビルドディレクトリ
BIN_DIR = bin

# 開発用ビルドディレクトリ
DEV_BIN_DIR = dev

# ビルドフラグ - 共通
COMMON_GCFLAGS = -N -l

# ビルドフラグ - 開発用
# -N: 最適化を無効化（デバッグ情報のために）
# -l: インライン化を無効化
DEV_GCFLAGS = $(COMMON_GCFLAGS)
DEV_LDFLAGS = -X 'main.version=$(VERSION)' -X 'main.commit=$(GIT_COMMIT)' -X 'main.buildTime=$(BUILD_TIME)'
DEV_BUILD_FLAGS = -gcflags="$(DEV_GCFLAGS)" -ldflags="$(DEV_LDFLAGS)"

# ビルドフラグ - 本番用
# -m: メモリアロケーションの最適化分析を有効化
PROD_GCFLAGS = -m
# -s: シンボルテーブルを削除
# -w: DWARF情報を削除
# -extldflags '-static': 静的リンク
PROD_LDFLAGS = -s -w -X 'main.version=$(VERSION)' -X 'main.commit=$(GIT_COMMIT)' -X 'main.buildTime=$(BUILD_TIME)' -extldflags '-static'
PROD_TRIMPATH = -trimpath
PROD_BUILD_FLAGS = -gcflags="$(PROD_GCFLAGS)" -ldflags="$(PROD_LDFLAGS)" $(PROD_TRIMPATH)

# デフォルトのビルド (開発用、現在のプラットフォーム用)
build:
	mkdir -p $(BIN_DIR)/$(DEV_BIN_DIR)
	go build $(DEV_BUILD_FLAGS) -o $(BIN_DIR)/$(DEV_BIN_DIR)/$(TARGET) $(SRC)

# 本番用ビルド (最適化済み、すべてのプラットフォーム用)
build-prod: build-windows build-linux build-darwin

# Windows用ビルド
# Windows AMD64 (64bit)
build-windows-amd64:
	mkdir -p $(BIN_DIR)/windows/amd64
	GOOS=windows GOARCH=amd64 go build $(PROD_BUILD_FLAGS) -o $(BIN_DIR)/windows/amd64/$(TARGET).exe $(SRC)

# Windows ARM64 (64bit ARM)
build-windows-arm64:
	mkdir -p $(BIN_DIR)/windows/arm64
	GOOS=windows GOARCH=arm64 go build $(PROD_BUILD_FLAGS) -o $(BIN_DIR)/windows/arm64/$(TARGET).exe $(SRC)

# Windows ARM (32bit ARM)
build-windows-arm:
	mkdir -p $(BIN_DIR)/windows/arm
	GOOS=windows GOARCH=arm go build $(PROD_BUILD_FLAGS) -o $(BIN_DIR)/windows/arm/$(TARGET).exe $(SRC)

# すべてのWindows用ビルド
build-windows: build-windows-amd64 build-windows-arm64 build-windows-arm

# Linux用ビルド
# Linux AMD64 (64bit)
build-linux-amd64:
	mkdir -p $(BIN_DIR)/linux/amd64
	GOOS=linux GOARCH=amd64 go build $(PROD_BUILD_FLAGS) -o $(BIN_DIR)/linux/amd64/$(TARGET) $(SRC)

# Linux ARM64 (64bit ARM)
build-linux-arm64:
	mkdir -p $(BIN_DIR)/linux/arm64
	GOOS=linux GOARCH=arm64 go build $(PROD_BUILD_FLAGS) -o $(BIN_DIR)/linux/arm64/$(TARGET) $(SRC)

# Linux ARM (32bit ARM)
build-linux-arm:
	mkdir -p $(BIN_DIR)/linux/arm
	GOOS=linux GOARCH=arm go build $(PROD_BUILD_FLAGS) -o $(BIN_DIR)/linux/arm/$(TARGET) $(SRC)

# すべてのLinux用ビルド
build-linux: build-linux-amd64 build-linux-arm64 build-linux-arm

# Mac (Darwin)用ビルド
# Darwin AMD64 (Intel Mac)
build-darwin-amd64:
	mkdir -p $(BIN_DIR)/darwin/amd64
	GOOS=darwin GOARCH=amd64 go build $(PROD_BUILD_FLAGS) -o $(BIN_DIR)/darwin/amd64/$(TARGET) $(SRC)

# Darwin ARM64 (Apple Silicon)
build-darwin-arm64:
	mkdir -p $(BIN_DIR)/darwin/arm64
	GOOS=darwin GOARCH=arm64 go build $(PROD_BUILD_FLAGS) -o $(BIN_DIR)/darwin/arm64/$(TARGET) $(SRC)

# すべてのDarwin用ビルド
build-darwin: build-darwin-amd64 build-darwin-arm64

# すべてのプラットフォーム用ビルド
release-build-all: build-prod

# テストコマンド
test:
	go test ./...

# ベンチマークテスト
benchmark:
	go test -bench=. ./...

# カバレッジレポート生成
coverage:
	mkdir -p $(BIN_DIR)/coverage
	go test -coverprofile=$(BIN_DIR)/coverage/coverage.out ./...
	go tool cover -html=$(BIN_DIR)/coverage/coverage.out -o=$(BIN_DIR)/coverage/coverage.html

# クリーンアップ
clean:
	rm -rf $(BIN_DIR)

# 全体のビルドとテスト
all: test build-prod

.PHONY: build build-prod build-windows build-linux build-darwin release-build-all build-windows-amd64 build-windows-arm64 build-windows-arm build-linux-amd64 build-linux-arm64 build-linux-arm build-darwin-amd64 build-darwin-arm64 test benchmark coverage clean all