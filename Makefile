run:
	@go run cmd/app/main.go

db:
	@docker-compose up -d

db-down:
	@docker-compose down

templ:
	@templ generate -watch -proxy=http://localhost:3000

