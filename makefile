css-watch:
	npx tailwindcss -i ./src/input.css -o ./assets/output.css --watch
css:
	npx tailwindcss -i ./src/input.css -o ./assets/output.css
css-minify:
	npx tailwindcss -i ./src/input.css -o ./assets/output.css --minify
build-app:
	CGO_ENABLED=0 go build -o /app ./main.go