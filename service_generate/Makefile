.PHONY: protoc
protoc:
	docker-compose run --rm --entrypoint sh protoc ./protoc.sh
	gofmt -s -w genpassword
	gofmt -s -w genpassword
	goimports -w -local "github.com/todo-server" genpassword