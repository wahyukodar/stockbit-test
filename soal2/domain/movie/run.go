package movie

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"stockbit_test/soal2/cinit"
	pb "stockbit_test/soal2/outbound/movie/proto"
)

const (
	SN = "api-movie"
)

var (
	ClientMovie  pb.MovieClient
	ClientMovies pb.MoviesClient
)

func Run() {
	cinit.InitOption(SN, cinit.MySQL)

	conn, err := grpc.Dial(cinit.Config.GrpcServiceMovie.Address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to service" + err.Error())
	}

	ClientMovie = pb.NewMovieClient(conn)
	ClientMovies = pb.NewMoviesClient(conn)

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH},
	}))

	service := InitService()

	g := e.Group("/v1/movie/search")
	g.GET("/details/:omdbId", service.SearchMovie)
	g.GET("/:searchWord/:page", service.SearchMovies)

	check := e.Group("/check")
	check.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	if err := e.Start(cinit.Config.RestServiceMovie.Port); err != nil {
		log.Fatal("Failed to start movie service" + err.Error())
	}
	defer conn.Close()
}
