// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"github.com/xuanvan229/crawl-golang/pkg/config"
	"github.com/xuanvan229/crawl-golang/pkg/crawl"
	"github.com/xuanvan229/crawl-golang/pkg/model"
)

func main() {

	config.LoadEnv()
	model.InitDB()
	crawl.CrawlLazada()

	//r := router.InitRouters()
	//r.Run(":8080")
}
