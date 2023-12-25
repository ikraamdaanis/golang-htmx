install:
	go get -u github.com/cosmtrek/air

run: 
	air

tailwind: 
	npx tailwindcss -i ./styles/input.css -o ./public/styles.css --watch

build:
	go build -o ./tmp/ ./cmd