package movieRpc

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	pb "stockbit_test/soal2/outbound/movie/proto"
	"strconv"
	"time"
)

type MoviesServer struct {
	pb.UnimplementedMoviesServer
}

type MovieServer struct {
	pb.UnimplementedMovieServer
}

func (s *MoviesServer) GetMovies(ctx context.Context, req *pb.GetMoviesRequest) (*pb.GetMoviesResponse, error) {

	var query = map[string]string{
		"s":    req.SearchWord,
		"page": strconv.Itoa(1),
	}

	client := &http.Client{
		Timeout: time.Duration(30) * time.Second,
	}
	_req, err := http.NewRequest(http.MethodGet, OMDBHost, nil)
	if err != nil {
		return nil, err
	}

	// append query string
	qryString := _req.URL.Query()
	qryString.Add("apikey", OMDBKey)
	for key, val := range query {
		qryString.Add(key, val)
	}
	_req.URL.RawQuery = qryString.Encode()

	resp, err := client.Do(_req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result SearchResult
	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	var pages float64
	data := make([]*pb.GetMoviesResponseData, len(result.Data))
	pages = math.Ceil(float64(result.TotalResults) / OMDBPageSize)

	for i := range result.Data {
		data[i] = &pb.GetMoviesResponseData{
			Title:  result.Data[i].Title,
			Year:   result.Data[i].Year,
			ImdbID: result.Data[i].ImdbID,
			Type:   result.Data[i].Type,
			Poster: result.Data[i].Poster,
		}
	}

	return &pb.GetMoviesResponse{
		Data:  data,
		Total: int32(result.TotalResults),
		Pages: int32(pages),
		Page:  req.Page,
	}, nil
}

func (s *MovieServer) GetMovie(ctx context.Context, req *pb.GetMovieRequest) (*pb.GetMovieResponse, error) {

	var query = map[string]string{
		"i": req.OmdbId,
	}

	client := &http.Client{
		Timeout: time.Duration(30) * time.Second,
	}
	_req, err := http.NewRequest(http.MethodGet, OMDBHost, nil)
	if err != nil {
		return nil, err
	}

	// append query string
	qryString := _req.URL.Query()
	qryString.Add("apikey", OMDBKey)
	for key, val := range query {
		qryString.Add(key, val)
	}
	_req.URL.RawQuery = qryString.Encode()

	resp, err := client.Do(_req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result DetailsResult
	err = json.Unmarshal(body, &result)
	ratings := make([]*pb.Ratings, len(result.Ratings))
	if err != nil {
		return nil, err
	}

	for i := range result.Ratings {
		ratings[i] = &pb.Ratings{
			Source: result.Ratings[i].Source,
			Value:  result.Ratings[i].Value,
		}
	}

	return &pb.GetMovieResponse{
		Title:      result.Title,
		Year:       result.Year,
		Rated:      result.Rated,
		Released:   result.Released,
		Runtime:    result.Runtime,
		Genre:      result.Genre,
		Director:   result.Director,
		Writer:     result.Writer,
		Actors:     result.Actors,
		Plot:       result.Plot,
		Language:   result.Language,
		Country:    result.Country,
		Awards:     result.Awards,
		Poster:     result.Poster,
		Ratings:    ratings,
		MetaScore:  result.Metascore,
		ImdbRating: result.ImdbRating,
		ImdbVotes:  result.ImdbVotes,
		ImdbID:     result.ImdbID,
		Type:       result.Type,
		Dvd:        result.DVD,
		BoxOffice:  result.BoxOffice,
		Production: result.Production,
		Website:    result.Website,
		Response:   result.Response,
		Error:      result.Error,
	}, nil
}
