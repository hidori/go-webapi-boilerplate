.PHONY: tool/install

# ツールをインストールします。
tool/install:
	go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install -v github.com/masakurapa/gover-html@latest
	go install -v github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: mod/download mod/tidy

# 依存関係パッケージをダウンロードします。
mod/download:
	go mod download

# 依存関係パッケージの参照を整理します。
mod/tidy:
	go mod tidy

.PHONY: env env/test env/compose

# .env が存在しなければ、.env.example から .env を作成します。
env:
	if [ ! -e .env ]; then \
		cp .env.example .env; \
	fi

# .env.test が存在しなければ、.env.test.example から .env.test を作成します。
env/test:
	if [ ! -e .env.test ]; then \
		cp .env.test.example .env.test; \
	fi

# .env.compose が存在しなければ、.env.compose.example から .env.compose を作成します。
env/compose:
	if [ ! -e .env.compose ]; then \
		cp .env.compose.example .env.compose; \
	fi

.PHONY: lint test cover build

# ソースコードの書式を検査します。
lint:
	golangci-lint -v run `find ./cmd/ ./internal/ ./pkg/ -type f -and -name '*.go' | xargs dirname | grep '^\.$$' | sort | uniq | grep -v '^\.$$'`

# 単体テストを実行します。
test: env/test db/reset/test db/migration/test
	-rm -r ./reports
	mkdir -p ./reports
	env `grep -v -e '^\s*#' -e '^\s*$$' .env.test` go test -v -shuffle=on -coverprofile=./reports/coverage.out -cover ./cmd/... ./internal/... ./pkg/...

# カバレッジレポートを作成します。
cover: test
	gover-html -i ./reports/coverage.out -o ./reports/coverage.html

# 実行形式ファイルを作成します。
build:
	go build -ldflags "-w -s" -o ./bin/webapi-api ./cmd/api/main.go

.PHONY: net/create net/rm

# Docker ネットワーク webapi-network を作成します。
net/create:
	-docker network create webapi-network

# Docker ネットワーク webapi-network を削除します。
net/rm:
	docker network rm webapi-network

.PHONY: compose/build compose/up compose/down compose/rebuild

# Docker コンテナをビルドします。
compose/build:
	docker-compose build

# Docker コンテナを開始します。
compose/up: env/compose net/create
	docker-compose --env-file ./.env.compose up -d

# Docker コンテナを終了します。
compose/down: env/compose
	docker-compose down

# Docker コンテナを終了、ビルド、開始します。
compose/rebuild: compose/down compose/build compose/up

.PHONY: db/reset db/migration db/reset/test db/migration/test

# ローカル DB を初期化します。
db/reset:
	(cd ../database; make db/reset)

# ローカル DB に対してマイグレーションを実行します。
db/migration:
	(cd ../database; make db/migration)

# 単体テスト用 DB を初期化します。
db/reset/test:
	(cd ../database; make db/reset/test)

# 単体テスト用 DB に対してマイグレーションを実行します。
db/migration/test:
	(cd ../database; make db/migration/test)
