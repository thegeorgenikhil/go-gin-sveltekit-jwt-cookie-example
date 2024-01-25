.PHONY: server frontend

server: 
	go run cmd/main.go

frontend:
	cd frontend && npm run dev