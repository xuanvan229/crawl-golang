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

	//crawl.CrawlLazada()

	r := router.InitRouters()
	r.Run(":8080")

	//
	//// create chrome instance
	//opts := append(chromedp.DefaultExecAllocatorOptions[:],
	//	chromedp.Flag("headless", false),
	//	chromedp.Flag("start-fullscreen", false),
	//	chromedp.Flag("enable-automation", false),
	//	chromedp.Flag("disable-extensions", false),
	//	chromedp.Flag("remote-debugging-port", "9222"),
	//)
	//allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	//
	//ctx, cancel := chromedp.NewContext(
	//	allocCtx,
	//	chromedp.WithLogf(log.Printf),
	//	// chromedp.WithDebugf(log.Printf),
	//)
	//defer cancel()
	//
	//// create a timeout
	//ctx, cancel = context.WithTimeout(ctx, 2000*time.Second)
	//defer cancel()
	//
	//// navigate to a page, wait for an element, click
	//var example string
	//var nodes []*cdp.Node
	//
	//urls := []string{
	//	"https://www.lazada.vn/dien-thoai-di-dong/",
	//	"https://www.lazada.vn/may-tinh-bang/",
	//	"https://www.lazada.vn/laptop/",
	//	"https://www.lazada.vn/may-tinh-de-ban-va-phu-kien/",
	//	"https://www.lazada.vn/am-thanh/",
	//	"https://www.lazada.vn/camera-giam-sat-2/",
	//	"https://www.lazada.vn/may-anh-may-quay-phim/",
	//	"https://www.lazada.vn/man-hinh-vi-tinh/",
	//	"https://www.lazada.vn/man-hinh-may-in/",
	//	"https://www.lazada.vn/dong-ho-thong-minh/",
	//	"https://www.lazada.vn/dieu-khien-choi-game/",
	//}
	//
	//var price, name, image string
	//var products []map[string]string
	//
	//for _, url := range urls {
	//	for i := 1; i < 10; i++ {
	//		pageUrl := url + fmt.Sprintf("?page=%d", i)
	//
	//		task := chromedp.Tasks{
	//			chromedp.Navigate(pageUrl),
	//			chromedp.WaitVisible(`._17mcb`),
	//			chromedp.Nodes(".Bm3ON", &nodes, chromedp.ByQueryAll),
	//		}
	//
	//		err := chromedp.Run(ctx, task)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//
	//		for index, node := range nodes {
	//			chromedp.Run(ctx,
	//				// chromedp.AttributeValue("a", "href", &url, nil, chromedp.ByQuery, chromedp.FromNode(node)),
	//				chromedp.AttributeValue("img", "src", &image, nil, chromedp.ByQuery, chromedp.FromNode(node)),
	//				chromedp.Text(".RfADt", &name, chromedp.ByQuery, chromedp.FromNode(node)),
	//				chromedp.Text(".ooOxS", &price, chromedp.ByQuery, chromedp.FromNode(node)),
	//			)
	//
	//			product := map[string]string{
	//				// "url": url,
	//				"price": price,
	//				"name":  name,
	//				// "image": image,
	//			}
	//			fmt.Println("products", index, ":  =>", product)
	//
	//			products = append(products, product)
	//		}
	//	}
	//}

	//log.Printf("Go's time.After example:\n%s", example)
}
