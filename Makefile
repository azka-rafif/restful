dev: docs  generate 
	go run .

generate:
	go generate ./...

docs:
	go run github.com/swaggo/swag/cmd/swag init 