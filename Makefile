build:
	go build -o build/api cmd/api/main.go

build-docker:
	docker build . -t taxi-cem-api