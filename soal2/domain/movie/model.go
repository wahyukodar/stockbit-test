package movie

import (
	"context"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/gommon/log"
	"stockbit_test/soal2/cinit"
	"stockbit_test/soal2/internal/utils"
	"time"
)

type LogMovie struct {
	ID         int64     `json:"id" db:"id" valid:"int"`
	LogRequest string    `json:"log_request" db:"log_request" valid:"required"`
	Action     string    `json:"action" db:"action" valid:"required"`
	CreatedAt  time.Time `json:"created_at" db:"created_at" valid:"required"`
}

type AddLogRequestMovieInterface interface {
	AddLogRequestMovie(context.Context) error
	Set(lM LogMovie)
}

func (l *LogMovie) Set(lM LogMovie) {
	l.CreatedAt = lM.CreatedAt
	l.Action = lM.Action
	l.LogRequest = lM.LogRequest
}

func InitLogRequestModel() AddLogRequestMovieInterface {
	return &LogMovie{}
}

func (l *LogMovie) validate() error {
	_, err := govalidator.ValidateStruct(l)
	return err
}

func (l *LogMovie) beforeAdd(ctx context.Context) error {
	err := utils.V(l.validate)
	if err != nil {
		return err
	}
	return nil
}

func (l *LogMovie) AddLogRequestMovie(ctx context.Context) error {
	err := l.beforeAdd(ctx)
	if err != nil {
		log.Info(err.Error(), ctx)
		return err
	}
	r, err := cinit.Mysql.Exec(`INSERT INTO logs_movie (log_request,action,created_at) VALUES (?,?,?)`,
		l.LogRequest, l.Action, l.CreatedAt)
	l.ID, err = utils.ID(r, err)
	if err != nil {
		log.Error(err.Error(), ctx)
		return err
	}
	return nil
}
