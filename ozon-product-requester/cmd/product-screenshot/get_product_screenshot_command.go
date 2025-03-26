package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"log"
	"os"
	"ozon-product-requester/internal/application/interactors"
	"sync"
)

type GetProductScreenshotArgs struct {
	Path       string   `arg:"positional,required"`
	IDs        []string `arg:"positional,required"`
	WindowPool int      `arg:"-w,--window-pool" default:"1"`
}

type GetProductScreenshotCommand struct {
	args              *GetProductScreenshotArgs
	interactorFactory func() interactors.MakeProductScreenshot
	wg                sync.WaitGroup
	errorChan         chan error
}

func NewGetProductScreenshotCommand(
	interactorFactory func() interactors.MakeProductScreenshot,
	args *GetProductScreenshotArgs,
) *GetProductScreenshotCommand {
	return &GetProductScreenshotCommand{
		args:              args,
		interactorFactory: interactorFactory,
		errorChan:         make(chan error, 1),
	}
}

func (c *GetProductScreenshotCommand) Execute() error {
	if err := arg.Parse(c.args); err != nil {
		return err
	}

	if _, err := os.Stat(c.args.Path); os.IsNotExist(err) {
		return fmt.Errorf("path does not exist: %s", c.args.Path)
	}

	idQueue := make(chan string, len(c.args.IDs))
	defer close(idQueue)
	for _, id := range c.args.IDs {
		idQueue <- id
	}

	for i := 0; i < c.args.WindowPool; i++ {
		c.wg.Add(1)
		go c.worker(i, idQueue)
	}

	go func() {
		c.wg.Wait()
		close(c.errorChan)
	}()

	if err := <-c.errorChan; err != nil {
		return err
	}

	return nil
}

func (c *GetProductScreenshotCommand) worker(workerID int, idQueue chan string) {
	defer c.wg.Done()

	interactor := c.interactorFactory()

	for {
		select {
		case id, ok := <-idQueue:
			if !ok {
				return
			}
			productImage, err := interactor.Call(id)
			if err != nil {
				c.errorChan <- fmt.Errorf("worker %d failed on ID %s: %w", workerID, id, err)
				return
			}

			err = os.WriteFile(c.args.Path+"/"+id+".png", productImage, os.ModePerm)

			if err != nil {
				c.errorChan <- fmt.Errorf("worker %d failed on ID %s: %w", workerID, id, err)
			}

			log.Printf("Worker %d processed ID %s", workerID, id)
		case err := <-c.errorChan:
			log.Printf("Worker %d stopping due to error: %v", workerID, err)
			return
		default:
			return
		}

	}

}
