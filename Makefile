
run:
	go run cmd/main.go

up:
	docker-compose up -d

category:
	go run cmd/crawl-category.go

crawl:
	go run cmd/crawl-products.go

crawl-image:
	go run cmd/crawl-product-image.go

build:
	go build cmd/main.go

