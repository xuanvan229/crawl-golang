// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"github.com/xuanvan229/crawl-golang/pkg/config"
	"github.com/xuanvan229/crawl-golang/pkg/model"
	"github.com/xuanvan229/crawl-golang/pkg/router"
)

func main() {

	config.LoadEnv()
	model.InitDB()

	r := router.InitRouters()
	r.Run(":8080")

}
