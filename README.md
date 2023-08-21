# GOLANG CRAWLING ECOMMERCE

## WHAT IS THIS?
This is a simple crawler that crawls the e-commerce website and stores the data in the database.

## HOW TO USE?
1. Install golang, docker, docker-compose
2. Pull the repository
3. Update the .env file like the .env.example
4. Run `Make up` to start the database
5. Run `Make category` to crawl the category data
6. Run `Make crawl` to crawl the product data
7. Run `Make crawl-image` to crawl the product image
8. Run `Make run` to run the server

## PROBLEM

1. Lazada has a lot of products, so it takes a lot of time to crawl all the products.
2. Using chromedp to crawl products page can not get the product's image. Because lazada uses lazy loading to load the image.
   
    Solution: After a lot of tries, I just crawl the products without images first. After that. I create a fix to crawl the image of product after that
3. Lazada have a way to prevent the crawler. It will redirect the crawler to the captcha page. I have to use the proxy to prevent this. After I tried to use a proxy to crawl the data, I found that the proxy is not stable. So I just crawl the data that I can crawl.

    Solution: I change the coffee shop everyday to crawl the data :)))

## FrontEnd (ReactJS)

Repository: https://github.com/xuanvan229/crawl-ecommerce-fe

Demo link: http://104.248.126.61/

## HOW TO IMPROVE?
