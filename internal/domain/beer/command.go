package domain

import (
	"context"
	"errors"

	"github.com/david9393/microservices-test/kit/command"
)

const BeerAddCommandType command.Type = "command.creating.beer"
const BeerGetAllCommandType command.Type = "command.getall.beer"
const BeerGetByIdCommandType command.Type = "command.getbyid.beer"

// BeerCommand is the command dispatched to create a new Beer.
type BeerCommand struct {
	Id      int    `json:"id"`
	Name    string `json:"name" binding:"required"`
	Brewery string `json:"brewery" binding:"required"`
}
type BeerByIdCommand struct {
	Id      int    `json:"id"`
	Name    string `json:"name" binding:"required"`
	Brewery string `json:"brewery" binding:"required"`
}
type GetAllBeerCommand struct {
}

// NewBeerCommand creates a new BeerCommand.
func NewBeerCommand(id int, name, brewery string) BeerCommand {
	return BeerCommand{
		Id:      id,
		Name:    name,
		Brewery: brewery,
	}
}

// NewBeerByIdCommand creates a new NewBeerByIdCommand.

func NewBeerByIdCommand(id int) BeerByIdCommand {
	return BeerByIdCommand{
		Id: id,
	}
}

// NewGetAllBeerCommand creates a new GetAllBeerCommand.
func NewGetAllBeerCommand(beer GetAllBeerCommand) GetAllBeerCommand {
	return GetAllBeerCommand{}
}

// NewBeerCommand creates a new BeerCommand.

func (R BeerCommand) Type() command.Type {
	return BeerAddCommandType
}

// NewBeerByIdCommand creates a new BeerByIdCommand.

func (R BeerByIdCommand) Type() command.Type {
	return BeerGetByIdCommandType
}

// NewGetAllBeerCommand creates a new GetAllBeerCommand.

func (R GetAllBeerCommand) Type() command.Type {
	return BeerGetAllCommandType
}

// responsible for creating Beers.
type BeerCommandHandler struct {
	service BeerService
}
type GetAllBeerCommandHandler struct {
	service BeerService
}
type BeerByIdCommandHandler struct {
	service BeerService
}

// NewBeerCommandHandler initializes a new BeerCommandHandler.
func NewBeerCommandHandler(service BeerService) BeerCommandHandler {
	return BeerCommandHandler{
		service: service,
	}
}

// NewBeerByIdCommandHandler initializes a new BeerByIdCommandHandler.
func NewBeerByIdCommandHandler(service BeerService) BeerByIdCommandHandler {
	return BeerByIdCommandHandler{
		service: service,
	}
}

// NewGetAllBeerCommandHandler initializes a new GetAllBeerCommandHandler.
func NewGetAllBeerCommandHandler(service BeerService) GetAllBeerCommandHandler {
	return GetAllBeerCommandHandler{
		service: service,
	}
}

// Handle implements the command.Handler interface.
func (h BeerCommandHandler) Handle(ctx context.Context, cmd command.Command) (interface{}, error) {
	createBeerCmd, ok := cmd.(BeerCommand)
	if !ok {
		return nil, errors.New("unexpected command")
	}

	return nil, h.service.CreateBeer(
		ctx,
		createBeerCmd,
	)

}

// Handle implements the command.Handler interface.
func (h GetAllBeerCommandHandler) Handle(ctx context.Context, cmd command.Command) (interface{}, error) {
	_, ok := cmd.(GetAllBeerCommand)
	if !ok {
		return nil, errors.New("unexpected command")
	}
	return h.service.GetAllBeer(
		ctx,
	)

}

// Handle implements the command.Handler interface.
func (h BeerByIdCommandHandler) Handle(ctx context.Context, cmd command.Command) (interface{}, error) {
	createBeerByIdCmd, ok := cmd.(BeerByIdCommand)
	if !ok {
		return nil, errors.New("unexpected command")
	}
	return h.service.GetByIdBeer(
		ctx,
		createBeerByIdCmd,
	)

}
