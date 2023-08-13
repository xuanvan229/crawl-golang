package crawl

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/google/uuid"
	"github.com/xuanvan229/crawl-golang/pkg/model"
	"log"
	"time"
)

var lazada []model.Lazada = []model.Lazada{
	//{
	//	Url:      "https://www.lazada.vn/dien-thoai-di-dong/",
	//	Category: "Điện thoại di động",
	//},
	//{
	//	Url:      "https://www.lazada.vn/may-tinh-bang/",
	//	Category: "Máy tính bảng",
	//},
	//{
	//	Url:      "https://www.lazada.vn/laptop/",
	//	Category: "Laptop",
	//},
	//{
	//	Url:      "https://www.lazada.vn/may-tinh-de-ban-va-phu-kien/",
	//	Category: "Máy tính để bàn và phụ kiện",
	//},
	//{
	//	Url:      "https://www.lazada.vn/am-thanh/",
	//	Category: "Âm thanh",
	//},
	//{
	//	Url:      "https://www.lazada.vn/camera-giam-sat-2/",
	//	Category: "Camera giám sát",
	//},
	{
		Url:      "https://www.lazada.vn/may-anh-may-quay-phim/",
		Category: "Máy ảnh, máy quay phim",
	},
	{
		Url:      "https://www.lazada.vn/man-hinh-vi-tinh/",
		Category: "Màn hình vi tính",
	},
	{
		Url:      "https://www.lazada.vn/man-hinh-may-in/",
		Category: "Màn hình máy in",
	},
	{
		Url:      "https://www.lazada.vn/dong-ho-thong-minh/",
		Category: "Đồng hồ thông minh",
	},
	{
		Url:      "https://www.lazada.vn/dieu-khien-choi-game/",
		Category: "Điều khiển chơi game",
	},
}

func CrawlLazada() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("start-fullscreen", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("remote-debugging-port", "9222"),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 2000*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var nodes []*cdp.Node

	var url, category string
	for _, item := range lazada {

		url = item.Url
		category = item.Category

		for i := 1; i < 10; i++ {
			pageUrl := url + fmt.Sprintf("?page=%d", i)
			var price, name, href, image, sold string
			var products []model.Product

			fmt.Println("pageUrl", pageUrl)
			task := chromedp.Tasks{
				chromedp.Navigate(pageUrl),
				chromedp.WaitVisible(`._17mcb`),
				chromedp.Nodes(".Bm3ON", &nodes, chromedp.ByQueryAll),
			}

			err := chromedp.Run(ctx, task)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("nodes", len(nodes))

			for index, node := range nodes {
				fmt.Println("index", index)
				chromedp.Run(ctx,
					chromedp.AttributeValue("a", "href", &href, nil, chromedp.ByQuery, chromedp.FromNode(node)),
					chromedp.AttributeValue("img", "src", &image, nil, chromedp.ByQuery, chromedp.FromNode(node)),
					chromedp.Text(".RfADt", &name, chromedp.ByQuery, chromedp.FromNode(node)),
					chromedp.Text(".ooOxS", &price, chromedp.ByQuery, chromedp.FromNode(node)),
					//chromedp.Text("._1cEkb", &sold, chromedp.ByQuery, chromedp.FromNode(node)),
				)

				product := model.Product{
					ID:        uuid.New(),
					Url:       href,
					Price:     price,
					Name:      name,
					Image:     image,
					Category:  category,
					TotalSold: sold,
					ShopName:  "lazada",
				}
				products = append(products, product)
			}

			fmt.Println("products", ":  =>", products)

			result, err := model.InsertProduct(products)
			fmt.Println("result", result)
			if err != nil {
				panic(err)
				return
			}
		}

	}

}