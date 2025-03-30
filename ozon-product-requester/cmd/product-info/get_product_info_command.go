package main

import (
	"github.com/alexflint/go-arg"
	"log"
	"ozon-product-requester/internal/application/interactors"
)

type GetProductInfoCommand struct {
	requestProductInteractor interactors.RequestProductInteractor
	args                     *GetProductInfoArgs
}

type GetProductInfoArgs struct {
	ID string `arg:"positional,required"`
}

func NewGetProductInfoCommand(
	product interactors.RequestProductInteractor,
	args *GetProductInfoArgs,
) *GetProductInfoCommand {
	return &GetProductInfoCommand{
		requestProductInteractor: product,
		args:                     args,
	}
}

func (command *GetProductInfoCommand) Execute() error {
	err := arg.Parse(command.args)
	if err != nil {
		return err
	}

	product, err := command.requestProductInteractor.Call(command.args.ID)

	if err != nil {
		return err
	}

	log.Println(product)

	return nil

}
