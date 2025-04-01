package ozon

import (
	"context"
	"errors"
	"log"
	product "ozon-product-requester/internal/domain/models"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

type ClientImpl struct {
	ctx context.Context
}

func NewClient(
	disableGpu bool,
	enableAutomation bool,
	disableBlinkFeatures string,
	headless bool,
	windowWidth int,
	windowHeight int,
	userAgent string,
) *ClientImpl {
	ctx, _ := chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.NoDefaultBrowserCheck,
			chromedp.DisableGPU,
			chromedp.Flag("disable-gpu", disableGpu),
			chromedp.Flag("enable-automation", enableAutomation),
			chromedp.Flag("disable-blink-features", disableBlinkFeatures),
			chromedp.Flag("headless", headless),
			chromedp.WindowSize(windowWidth, windowHeight),
			chromedp.UserAgent(userAgent),
		)...,
	)

	ctx, _ = chromedp.NewContext(
		ctx,
		chromedp.WithLogf(log.Printf),
	)

	return &ClientImpl{
		ctx: ctx,
	}
}

func (c *ClientImpl) MakeScreenshot(id string) ([]byte, error) {
	var screenshot []byte
	if err := chromedp.Run(c.ctx, chromedp.Tasks{
		chromedp.Navigate("https://ozon.ru/product/" + id),
		chromedp.Sleep(3 * time.Second),
		chromedp.CaptureScreenshot(&screenshot),
	}); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return screenshot, nil
}

func (c *ClientImpl) GetProduct(id string) (*product.Product, error) {
	var price string
	var name string
	err := chromedp.Run(c.ctx, chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			navCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
			defer cancel()
			return chromedp.Navigate("https://ozon.ru/product/" + id).Do(navCtx)
		}),

		chromedp.ActionFunc(func(ctx context.Context) error {
			waitCtx, cancel := context.WithTimeout(ctx, time.Second)
			defer cancel()
			return chromedp.WaitReady("body").Do(waitCtx)
		}),

		chromedp.ActionFunc(func(ctx context.Context) error {
			priceCtx, cancel := context.WithTimeout(ctx, time.Second)
			defer cancel()
			return chromedp.Text("[data-widget=\"webPrice\"]", &price).Do(priceCtx)
		}),

		chromedp.ActionFunc(func(ctx context.Context) error {
			nameCtx, cancel := context.WithTimeout(ctx, time.Second)
			defer cancel()
			return chromedp.Text("[data-widget=\"webProductHeading\"]", &name).Do(nameCtx)
		}),
	})

	if err != nil {
		return nil, err
	}

	price = strings.ReplaceAll(price, " ", "")

	matches := regexp.MustCompile(`\d+(?:\d+)+?`).FindAllStringSubmatch(price, -1)
	if len(matches) == 0 {
		return nil, errors.New("no price found")
	}

	ozonCardPrice, err := strconv.ParseFloat(matches[0][0], 64)
	if err != nil {
		return nil, err
	}

	usualPrice, err := strconv.ParseFloat(matches[1][0], 64)
	if err != nil {
		return nil, err
	}

	ozonProduct, err := product.NewProduct(
		product.ProductID(id),
		product.Price{
			Value:    ozonCardPrice,
			Currency: product.RUB,
		},
		product.Price{
			Value:    usualPrice,
			Currency: product.RUB,
		},
		product.Name(name),
		product.Available,
	)

	return ozonProduct, err
}
