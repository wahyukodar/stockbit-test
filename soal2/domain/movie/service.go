package movie

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"stockbit_test/soal2/internal/handler"
	"stockbit_test/soal2/internal/utils"
	pb "stockbit_test/soal2/outbound/movie/proto"
	"strconv"
)

type LogMovieService struct {
	logMovie AddLogRequestMovieInterface
}

func InitService() *LogMovieService {
	return &LogMovieService{
		logMovie: InitLogRequestModel(),
	}
}

func (l *LogMovieService) SearchMovies(c echo.Context) error {
	ctx := c.Request().Context()

	searchWord := c.Param("searchWord")
	page, err := utils.S2I32(c.Param("page"))

	if err != nil {
		log.Error(err.Error(), ctx)
		return handler.HandleError(c, handler.BusParamConvertError, err.Error())
	}

	l.logMovie.Set(LogMovie{
		CreatedAt:  utils.MicroTime(),
		LogRequest: "{SearchWord: " + searchWord + ", Page: " + strconv.Itoa(int(page)) + "}",
		Action:     "searchAllBySearchWord",
	})

	err = l.logMovie.AddLogRequestMovie(ctx)
	if err != nil {
		log.Error(err.Error(), ctx)
		return handler.HandleError(c, handler.DatabaseError, err.Error())
	}

	req := &pb.GetMoviesRequest{SearchWord: searchWord, Page: page}

	rsp, err := ClientMovies.GetMovies(ctx, req)
	if err != nil {
		log.Error(err.Error(), ctx)
		return handler.RPCErr(c, err)
	}
	return handler.HandleSuccess(c, rsp.Data, rsp.Pages, rsp.Page, rsp.Total)
}

func (l *LogMovieService) SearchMovie(c echo.Context) error {
	ctx := c.Request().Context()

	omdbId := c.Param("omdbId")

	if omdbId == "" {
		return handler.HandleError(c, handler.BusParamConvertError, "omdbId is required")
	}

	l.logMovie.Set(LogMovie{
		CreatedAt:  utils.MicroTime(),
		LogRequest: "{omdbIds: " + omdbId + "}",
		Action:     "searchOneByOmdbId",
	})

	err := l.logMovie.AddLogRequestMovie(ctx)
	if err != nil {
		log.Error(err.Error(), ctx)
		return handler.HandleError(c, handler.DatabaseError, err.Error())
	}

	req := &pb.GetMovieRequest{OmdbId: omdbId}

	rsp, err := ClientMovie.GetMovie(ctx, req)
	if err != nil {
		log.Error(err.Error(), ctx)
		return handler.RPCErr(c, err)
	}
	return handler.HandleSuccess(c, rsp)
}
