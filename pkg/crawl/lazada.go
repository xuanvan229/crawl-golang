package crawl

import (
	"context"
	"errors"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func GetProductImage(pageUrl string) (string, error) {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("start-fullscreen", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.Flag("disable-extensions", true),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30000*time.Second)
	defer cancel()

	var nodes []*cdp.Node
	var errorPage interface{}
	var image string

	pageUrl = "https:" + pageUrl

	findErrorTask := chromedp.Tasks{
		chromedp.Navigate(pageUrl),
		chromedp.EvaluateAsDevTools(`document.querySelector(".comm-error")`, &errorPage),
	}

	err := chromedp.Run(ctx, findErrorTask)
	if err != nil {
		log.Fatal(err)
	}

	if errorPage != nil {
		return "", errors.New("error")
	}

	task := chromedp.Tasks{
		chromedp.Navigate(pageUrl),
		chromedp.WaitVisible(`.gallery-preview-panel__image`),
		chromedp.Nodes(".gallery-preview-panel__content", &nodes, chromedp.ByQueryAll),
	}

	err = chromedp.Run(ctx, task)

	if err != nil {
		log.Fatal(err)
	}

	for _, node := range nodes {
		err := chromedp.Run(ctx,
			chromedp.AttributeValue(".gallery-preview-panel__image", "src", &image, nil, chromedp.ByQuery, chromedp.FromNode(node)),
		)
		if err != nil {
			return "", err
		}
	}
	chromedp.Stop()
	return image, nil
}
