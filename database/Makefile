.PHONY: tool/install

# ツールをインストールします。
tool/install:
	go install -v github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: env

# .env が存在しなければ、.env.example から .env を作成します。
env:
	if [ ! -e .env ]; then \
		cp .env.example .env; \
	fi

.PHONY: build run

# 実行形式ファイルを作成します。
build:
	docker build -f ./Dockerfile_migration -t webapi/migration .

# webapi-api を実行します。
run: env build net/create
	docker run -it --network webapi-network --env-file ./.env webapi/migration

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
compose/up: net/create
	docker-compose up -d

# Docker コンテナを終了します。
compose/down:
	docker-compose down

# Docker コンテナを終了、ビルド、開始します。
compose/rebuild: compose/down compose/build compose/up

.PHONY: db/reset db/migration db/reset/test db/migration/test

# ローカル DB を初期化します。
db/reset:
	MYSQL_PWD=P@ssw0rd mysql -h 127.0.0.1 -P 13306 -u root -e 'DROP DATABASE IF EXISTS webapi;'
	MYSQL_PWD=P@ssw0rd mysql -h 127.0.0.1 -P 13306 -u root -e 'CREATE DATABASE webapi DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;'

# ローカル DB に対してマイグレーションを実行します。
db/migration:
	(cd ./migrations; goose mysql 'root:P@ssw0rd@tcp(127.0.0.1:13306)/webapi?charset=utf8mb4&parseTime=True' up)

# 単体テスト用 DB を初期化します。
db/reset/test:
	for db in webapi_test webapi_test_repository webapi_test_usecase; do \
		MYSQL_PWD=P@ssw0rd mysql -h 127.0.0.1 -P 13306 -u root -e "DROP DATABASE IF EXISTS $${db};"; \
		MYSQL_PWD=P@ssw0rd mysql -h 127.0.0.1 -P 13306 -u root -e "CREATE DATABASE $${db} DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;"; \
	done

# 単体テスト用 DB に対してマイグレーションを実行します。
db/migration/test:
	for db in webapi_test webapi_test_repository webapi_test_usecase; do \
		(cd ./migrations; goose mysql "root:P@ssw0rd@tcp(127.0.0.1:13306)/$${db}?charset=utf8mb4&parseTime=True" up); \
	done
