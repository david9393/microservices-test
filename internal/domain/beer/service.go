package domain

import (
	"context"

	"github.com/david9393/microservices-test/internal/mooc"
	"github.com/david9393/microservices-test/kit/event"
)

// BeerService is the default BeerService interface
// implementation returned by creating.NewBeerService.
type BeerService struct {
	beerRepository mooc.BeerRepository
	eventBus       event.Bus
}

// NewBeerService returns the default Service interface implementation.
func NewBeerService(beerRepository mooc.BeerRepository, eventBus event.Bus) BeerService {
	return BeerService{
		beerRepository: beerRepository,
		eventBus:       eventBus,
	}
}

// CreateBeer implements the creating.BeerService interface.
func (s BeerService) CreateBeer(ctx context.Context, r BeerCommand) error {

	beer, err := mooc.NewBeer(r.Id, r.Name, r.Brewery)
	if err != nil {
		return err
	}

	if err := s.beerRepository.Save(ctx, beer); err != nil {
		return err
	}
	return nil
}

// GetAllBeer implements the getAllBeer.BeerService interface.
func (s BeerService) GetAllBeer(ctx context.Context) ([]BeerCommand, error) {
	listBeers := []BeerCommand{}

	beers, err := s.beerRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range beers {
		beer := BeerCommand{}
		beer.Id = v.BeerId().Int()
		beer.Name = v.BeerName().String()
		beer.Brewery = v.BeerBrewery().String()
		listBeers = append(listBeers, beer)
	}

	return listBeers, nil
}

// GetByIdBeer implements the getByIdBeer.BeerService interface.
func (s BeerService) GetByIdBeer(ctx context.Context, r BeerByIdCommand) ([]BeerCommand, error) {
	listBeers := []BeerCommand{}

	beers, err := s.beerRepository.GetById(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	for _, v := range beers {
		beer := BeerCommand{}
		beer.Id = v.BeerId().Int()
		beer.Name = v.BeerName().String()
		beer.Brewery = v.BeerBrewery().String()
		listBeers = append(listBeers, beer)
	}

	return listBeers, nil
}
