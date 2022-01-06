package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/david9393/microservices-test/internal/mooc"
	"github.com/huandu/go-sqlbuilder"
)

// BeerRepository is a MySQL mooc.BeerRepository implementation.
type BeerRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewBeerRepository initializes a MySQL-based implementation of mooc.BeerRepository.
func NewBeerRepository(db *sql.DB, dbTimeout time.Duration) *BeerRepository {
	return &BeerRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// Save implements the mooc.BeerRepository interface.
func (r *BeerRepository) Save(ctx context.Context, beer mooc.Beer) error {
	rolSQLStruct := sqlbuilder.NewStruct(new(sqlInsertBeer))
	_, args := rolSQLStruct.InsertIgnoreInto(sqlBeerTable, sqlInsertBeer{
		Name:    beer.BeerName().String(),
		Brewery: beer.BeerBrewery().String(),
	}).Build()
	const query = `INSERT INTO public.beer (name,brewery) VALUES ($1,$2);`
	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist Beer on database: %v", err)
	}

	return nil
}

// GetAll implements the mooc.BeerRepository interface.
func (r *BeerRepository) GetAll(ctx context.Context) ([]mooc.Beer, error) {
	listBeer := []mooc.Beer{}

	const query = `SELECT * FROM  public.beer`
	rows, err := r.db.Query(query)
	if err != nil {
		return listBeer, err
	}
	for rows.Next() {
		beerSql := &sqlBeer{}
		err = rows.Scan(&beerSql.Id, &beerSql.Name, &beerSql.Brewery)
		if err != nil {
			return listBeer, err
		}
		beer, _ := mooc.NewBeer(beerSql.Id, beerSql.Name, beerSql.Brewery)
		listBeer = append(listBeer, beer)
	}
	println(listBeer)
	return listBeer, nil
}

// GetById implements the mooc.BeerRepository interface.
func (r *BeerRepository) GetById(ctx context.Context, id int) ([]mooc.Beer, error) {
	listBeer := []mooc.Beer{}

	const query = `SELECT * FROM  public.beer WHERE id=$1`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return listBeer, err
	}
	for rows.Next() {
		beerSql := &sqlBeer{}
		err = rows.Scan(&beerSql.Id, &beerSql.Name, &beerSql.Brewery)
		if err != nil {
			return listBeer, err
		}
		beer, _ := mooc.NewBeer(beerSql.Id, beerSql.Name, beerSql.Brewery)
		listBeer = append(listBeer, beer)
	}
	println(listBeer)
	return listBeer, nil
}
