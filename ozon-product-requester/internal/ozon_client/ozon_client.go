package ozon_client

import (
	"context"
	"errors"
	"fmt"
	"log"
	product "ozon-product-requester/internal/domain/product/entity"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

const USER_AGENT string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36"

type Client struct {
	ctx context.Context
}

func NewOzonClient() *Client {
	ctx, _ := chromedp.NewExecAllocator(
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

	ctx, _ = chromedp.NewContext(
		ctx,
		chromedp.WithLogf(log.Printf),
	)

	return &Client{
		ctx: ctx,
	}
}

func (c *Client) MakeScreeshot(id string) ([]byte, error) {
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

func (c *Client) GetProduct(id string) (*product.Product, error) {
	var price string
	var name string
	if err := chromedp.Run(c.ctx, chromedp.Tasks{
		chromedp.Navigate("https://ozon.ru/product/" + id),
		chromedp.WaitReady("body"),
		chromedp.Text("[data-widget=\"webPrice\"]", &price),
		chromedp.Text("[data-widget=\"webProductHeading\"]", &name),
	}); err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println(price)
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
