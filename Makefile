default: debug-unhealthy

debug-unhealthy:
	GOOS=linux go build -o bin/debug-unhealthy cmd/debug-unhealthy/main.go

