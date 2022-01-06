package postgres

const (
	sqlBeerTable = "beer"
)

type sqlBeer struct {
	Id      int    `db:"id"`
	Name    string `db:"name"`
	Brewery string `db:"brewery"`
}
type sqlInsertBeer struct {
	Name    string `db:"name"`
	Brewery string `db:"brewery"`
}
