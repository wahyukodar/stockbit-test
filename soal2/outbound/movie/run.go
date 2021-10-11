package movieRpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"stockbit_test/soal2/cinit"
	pb "stockbit_test/soal2/outbound/movie/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	SN = "rpc-movie"
)

func Run() {
	cinit.InitOption(SN)

	listener, err := net.Listen("tcp", cinit.Config.GrpcServiceMovie.Port)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterMovieServer(srv, &MovieServer{})
	pb.RegisterMoviesServer(srv, &MoviesServer{})

	reflection.Register(srv)

	// Start a middleware service
	ctx := context.Background()

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handler it
			fmt.Println()
			log.Println("Shutting down gRPC server...")

			srv.GracefulStop()

			<-ctx.Done()
		}
	}()

	if err := srv.Serve(listener); err != nil {
		log.Fatal("Failed to listen: " + err.Error())
	}
}
