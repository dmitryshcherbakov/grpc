.PHONY: protoc
protoc:
	docker-compose run --rm --entrypoint sh protoc ./protoc.sh
	gofmt -s -w proto
	gofmt -s -w proto
	goimports -w -local "github.com/todo-server" proto
test:
	go test -cover -race ./...