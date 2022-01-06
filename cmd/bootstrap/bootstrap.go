package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	domain "github.com/david9393/microservices-test/internal/domain/beer"
	"github.com/david9393/microservices-test/internal/platform/bus"
	"github.com/david9393/microservices-test/internal/platform/server"
	"github.com/david9393/microservices-test/internal/platform/storage/postgres"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

func Run() error {
	var cfg config
	err := envconfig.Process("config", &cfg)
	if err != nil {
		return err
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=		%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	var (
		commandBus = bus.NewCommandBus()
		eventBus   = bus.NewEventBus()
	)

	beerRepository := postgres.NewBeerRepository(db, cfg.DbTimeout)
	creatingBeerService := domain.NewBeerService(beerRepository, eventBus)
	//Create Command Handler
	createBeerCommandHandler := domain.NewBeerCommandHandler(creatingBeerService)
	getAllBeerCommandHandler := domain.NewGetAllBeerCommandHandler(creatingBeerService)
	getByIdBeerCommandHandler := domain.NewBeerByIdCommandHandler(creatingBeerService)

	//Register Command
	commandBus.Register(domain.BeerAddCommandType, createBeerCommandHandler)
	commandBus.Register(domain.BeerGetAllCommandType, getAllBeerCommandHandler)
	commandBus.Register(domain.BeerGetByIdCommandType, getByIdBeerCommandHandler)
	//Se Levanta el servidor
	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, commandBus)
	return srv.Run(ctx)
}

type config struct {
	// Server configuration
	Host            string        `default:"localhost"`
	Port            uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"10s"`
	// Database configuration
	DbUser     string        `default:"postgres"`
	DbPass     string        `default:"admin"`
	DbHost     string        `default:"localhost"`
	DbPort     uint          `default:"5432"`
	DbName     string        `default:"Beer"`
	DbPassword string        `default:"1234"`
	DbTimeout  time.Duration `default:"10s"`
}
