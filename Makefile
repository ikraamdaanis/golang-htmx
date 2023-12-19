install:
	go get -u github.com/cosmtrek/air
	go install github.com/a-h/templ/cmd/templ@latest

run: 
	air

tailwind: 
	npx tailwindcss -i ./styles/input.css -o ./public/styles.css --watch

build:
	go build -o ./tmp/ ./cmd