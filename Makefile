
run:
	go run cmd/main.go

up:
	docker-compose up -d

crawl:
	go run cmd/crawl.go

fix:
	go run cmd/fixCrawl.go

build:
	go build cmd/main.go