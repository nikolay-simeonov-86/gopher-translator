start-dev:
	redis-server --daemonize yes
ifdef PORT
	@echo Started server on port $(PORT)
	go run cmd/server/main.go -port $(PORT)
else
	@echo Started server on port 8080
	go run cmd/server/main.go
endif