package beer

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/david9393/microservices-test/internal/mooc"

	domain "github.com/david9393/microservices-test/internal/domain/beer"
	"github.com/david9393/microservices-test/kit/command"
	"github.com/gin-gonic/gin"
)

// CreateHandler returns an HTTP handler for beer creation.
func CreateHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domain.BeerCommand
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		_, err := commandBus.Dispatch(ctx, domain.NewBeerCommand(
			req.Id,
			req.Name,
			req.Brewery,
		))

		if err != nil {
			switch {
			case errors.Is(err, mooc.ErrEmptyBeerName),
				errors.Is(err, mooc.ErrEmptyBeerBrewery):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}

// CreateHandler returns an HTTP handler for get beers.
func GetAllBeerHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domain.GetAllBeerCommand
		beers, err := commandBus.Dispatch(ctx, domain.NewGetAllBeerCommand(req))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return

		}

		ctx.JSON(http.StatusOK, beers)
	}
}

// CreateHandler returns an HTTP handler for get beers by id.
func GetByIdBeerHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Query("beerID"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		beers, err := commandBus.Dispatch(ctx, domain.NewBeerByIdCommand(id))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return

		}

		ctx.JSON(http.StatusOK, beers)
	}
}
