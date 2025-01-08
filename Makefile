run: 
	@templ generate
	@go run main.go
live-dev:
	@templ generate --watch -proxybind="localhost" --proxy="http://localhost:8090" --cmd="air"
tailwind-watch: ## compile tailwindcss and watch for changes
	tailwindcss-extra -i ./static/css/custom.css -o ./static/css/style.css --watch
