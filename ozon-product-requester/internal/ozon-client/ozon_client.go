package ozon_client

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

const USER_AGENT string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36"

var ctx context.Context

func init() {
	ctx, _ = chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.NoDefaultBrowserCheck,
			chromedp.DisableGPU,
			chromedp.Flag("disable-gpu", true),
			chromedp.Flag("enable-automation", true),
			chromedp.Flag("disable-blink-features", "AutomationControlled"),
			chromedp.Flag("headless", true),
			chromedp.WindowSize(1920, 1080),
			chromedp.UserAgent(USER_AGENT),
		)...,
	)

	ctx, _ = context.WithTimeout(ctx, 5*time.Second)
	ctx, _ = chromedp.NewContext(
		ctx,
		chromedp.WithLogf(log.Printf),
	)
}

func MakeScreeshot(id string) (*[]byte, error) {
	var screenshot []byte
	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate("https://ozon.ru/product/" + id),
		chromedp.WaitReady(".webSale"),
		chromedp.CaptureScreenshot(&screenshot),
	}); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &screenshot, nil
}
