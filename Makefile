default: debug-unhealthy

debug-unhealthy:
	GOOS=linux go build -o debug-unhealthy cmd/debug-unhealthy/main.go

