.PHONY: protoc
protoc:
	docker-compose run --rm --entrypoint sh protoc ./protoc.sh
	gofmt -s -w chat
	gofmt -s -w chat
	goimports -w -local "github.com/todo-server" chat