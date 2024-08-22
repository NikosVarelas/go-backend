run:
	@go run cmd/app/main.go

db:
	@docker-compose up -d db

db-down:
	@docker-compose down

templ:
	@templ generate -watch -proxy=http://localhost:3000

tailwind:
	@npx tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch  

tailwind-build:
	@npx ./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify
