package mooc

import (
	"context"
	"errors"

	"github.com/david9393/microservices-test/kit/event"
)

var ErrInvalidIdBeer = errors.New("invalid Beer Id")

// BeerId represents the beer id.
type BeerId struct {
	value int
}

// NewBeerId instantiate the VO for BeerId
func NewBeerId(value int) (BeerId, error) {
	if value < 1 {
		return BeerId{}, ErrInvalidIdBeer
	}

	return BeerId{
		value: value,
	}, nil
}

// int type converts the BeerId into int.
func (id BeerId) Int() int {
	return id.value
}

var ErrEmptyBeerName = errors.New("invalid Name")

// BeerName represents the Beer name.
type BeerName struct {
	value string
}

// NewBeerName instantiate VO for BeerName
func NewBeerName(value string) (BeerName, error) {
	if value == "" {
		return BeerName{}, ErrEmptyBeerName
	}

	return BeerName{
		value: value,
	}, nil
}

// String type converts the BeerName into string.
func (name BeerName) String() string {
	return name.value
}

var ErrEmptyBeerBrewery = errors.New("invalid  Beer Brewery ")

// BeerBrewery represents the beer  Brewery.
type BeerBrewery struct {
	value string
}

// NewBeerName instantiate VO for BeerName
func NewBeerBrewery(value string) (BeerBrewery, error) {
	if value == "" {
		return BeerBrewery{}, ErrEmptyBeerBrewery
	}

	return BeerBrewery{
		value: value,
	}, nil
}

// String type converts the  BreweryName into string.
func (name BeerBrewery) String() string {
	return name.value
}

// 	User is the data structure that represents a user.
type Beer struct {
	id      BeerId
	name    BeerName
	brewery BeerBrewery
	events  []event.Event
}

// BeerRepository defines the expected behaviour from a user storage.
type BeerRepository interface {
	Save(ctx context.Context, beer Beer) error
	GetAll(ctx context.Context) ([]Beer, error)
	GetById(ctx context.Context, id int) ([]Beer, error)
}

// NewUser creates a new user.
func NewBeer(id int, name string, brewery string) (Beer, error) {
	idVO, err := NewBeerId(id)
	nameVO, err := NewBeerName(name)
	breweryVO, err := NewBeerBrewery(brewery)
	if err != nil {
		return Beer{}, err
	}
	beer := Beer{
		id:      idVO,
		name:    nameVO,
		brewery: breweryVO,
	}
	return beer, nil
}

// ID returns the Beer unique identifier.
func (r Beer) BeerId() BeerId {
	return r.id
}

// Name returns the user name.
func (r Beer) BeerName() BeerName {
	return r.name
}

// Name returns the user name.
func (r Beer) BeerBrewery() BeerBrewery {
	return r.brewery
}

// Record records a new domain event.
func (r *Beer) Record(evt event.Event) {
	r.events = append(r.events, evt)
}

// PullEvents returns all the recorded domain events.
func (r Beer) PullEvents() []event.Event {
	evt := r.events
	r.events = []event.Event{}

	return evt
}
