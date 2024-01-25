.PHONY: server frontend

server: 
	go run cmd/main.go

frontend:
	cd frontend && npm run dev

install_frontend:
	cd frontend && npm install

install_server:
	go mod tidy