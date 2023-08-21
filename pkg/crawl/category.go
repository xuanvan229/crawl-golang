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

var lazadaUrl = "https://www.lazada.vn/"

func GetOpts(disableImage bool, headless bool) []chromedp.ExecAllocatorOption {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("start-fullscreen", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", true),
	)

	if headless {
		opts = append(opts, chromedp.Flag("headless", true))
	} else {
		opts = append(opts, chromedp.Flag("headless", false))
	}

	if disableImage {
		opts = append(opts, chromedp.Flag("blink-settings", "imagesEnabled=false"))
	}

	return opts
}

func GetAllCategory() {
	opts := GetOpts(true, true)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	var nodes []*cdp.Node

	task := chromedp.Tasks{
		chromedp.Navigate(lazadaUrl),
		chromedp.WaitVisible(`.lzd-home-page-category`),
		chromedp.Nodes(`.lzd-home-page-category .lzd-site-menu-root-item`, &nodes),
	}

	err := chromedp.Run(ctx, task)
	if err != nil {
		log.Fatal("err GetAllCategory", err)
	}
	var urls []string

	for index, _ := range nodes {
		urls = append(urls, GetCategoriesOfMasterCategory(index+1)...)
	}

	var categories []model.CategoryLazada

	for i := 0; i < len(urls); i++ {
		category := model.CategoryLazada{
			ID:      uuid.New(),
			Url:     urls[i],
			Crawled: false,
		}
		categories = append(categories, category)

	}

	_, err = model.InsertCategory(categories)

	if err != nil {
		panic(err)
	}

}

func GetProductOfCategoryNotCrawled() {
	categories, err := model.GetCategoryNotCrawled()
	if err != nil {
		panic(err)
	}

	for _, category := range *categories {
		products := GetProductsOfCategory(category.Url)
		category.Crawled = true
		_, err := model.UpdateCategory(category)
		if err != nil {
			panic(err)
		}

		_, err = model.InsertProduct(products)
		if err != nil {
			panic(err)
			return
		}
	}
}

func GetCategoriesOfMasterCategory(index int) []string {

	var urls []string
	opts := GetOpts(true, true)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30000*time.Second)
	defer cancel()

	className := fmt.Sprintf(".lzd-home-page-category .Level_1_Category_No%d", index)
	var nodes []*cdp.Node

	task := chromedp.Tasks{
		chromedp.Navigate(lazadaUrl),
		chromedp.WaitVisible(`.lzd-home-page-category`),
		chromedp.Nodes(className, &nodes),
	}

	err := chromedp.Run(ctx, task)
	if err != nil {
		log.Fatal("err GetCategoriesOfMasterCategory", err)
	}

	for _, node := range nodes {
		var lis []*cdp.Node

		chromedp.Run(ctx,
			chromedp.Nodes("li > a", &lis, chromedp.ByQueryAll, chromedp.FromNode(node)),
		)

		for _, li := range lis {
			urls = append(urls, li.AttributeValue("href"))
		}
	}

	return urls
}

func GetProductsOfCategory(url string) []model.Product {

	var result []model.Product

	opts := GetOpts(true, false)

	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30000*time.Second)
	defer cancel()

	var nodes []*cdp.Node

	var category = ""

	for i := 1; i <= 2; i++ {
		pageUrl := "https:" + url + fmt.Sprintf("?page=%d", i)
		var price, name, href, image, sold string

		task := chromedp.Tasks{

			chromedp.Navigate(pageUrl),

			chromedp.WaitVisible(`._17mcb`),
			chromedp.Nodes(".Bm3ON", &nodes, chromedp.ByQueryAll),
			chromedp.Text("._8akZL", &category, chromedp.ByQuery),
		}

		err := chromedp.Run(ctx, task)
		if err != nil {
			log.Fatal(err)
		}

		for _, node := range nodes {
			chromedp.Run(ctx,
				chromedp.AttributeValue("a", "href", &href, nil, chromedp.ByQuery, chromedp.FromNode(node)),
				chromedp.AttributeValue("img", "src", &image, nil, chromedp.ByQuery, chromedp.FromNode(node)),
				chromedp.Text(".RfADt", &name, chromedp.ByQuery, chromedp.FromNode(node)),
				chromedp.Text(".ooOxS", &price, chromedp.ByQuery, chromedp.FromNode(node)),
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

			result = append(result, product)
		}

	}

	return result
}
