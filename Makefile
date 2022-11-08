folder: ## create folder structure
	@mkdir -p cmd
	@mkdir -p cmd/server
	@mkdir -p config
	@mkdir -p internal
	@mkdir -p internal/adapter
	@mkdir -p internal/adapter/model
	@mkdir -p internal/api
	@mkdir -p internal/repository
	@mkdir -p internal/repository/model
	@mkdir -p internal/dto
	@mkdir -p internal/helper
	@mkdir -p internal/helper/db
	@mkdir -p internal/helper/log
	@mkdir -p internal/helper/memcache
	@mkdir -p internal/registry
	@mkdir -p internal/usecase
	@mkdir -p internal/util

depends:
	@go get github.com/rakyll/statik
	@go get github.com/mattn/go-oci8

install:
	@sudo chmod +x scripts/swagger.sh
	sh scripts/swagger.sh
	@go install github.com/rakyll/statik@latest

gen-swagger:
	@swagger generate spec -w ./cmd/server -o ./docs/swagger.json --scan-models

run:
	@statik -f -src=docs -dest=cmd/server && cd cmd/server && go run main.go

docker-image:
	@docker build --tag project-base .

docker-run:
	@docker run --publish 10001:10001 project-base
