run: 
	@go build -o bin/main && ./bin/main
live-dev: # usado para desenvolvimento
	@templ generate --watch -proxybind="localhost" --proxy="http://localhost:8090" --cmd="air"
tailwind-watch: ## compile tailwindcss and watch for changes, também usado para desenvolvimento
	tailwindcss-extra -i ./static/css/custom.css -o ./static/css/style.css --watch
