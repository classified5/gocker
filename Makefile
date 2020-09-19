run: 
	@go run main.go

test:
	@go clean -testcache
	@go test ./... -cover -race

docker-build:
	@docker build . -t classified5/gocker:latest

docker-run:
	@docker run -d -p 9090:9090 --name gocker classified5/gocker:latest

docker-stop:
	@docker stop gocker
	@docker rm gocker	