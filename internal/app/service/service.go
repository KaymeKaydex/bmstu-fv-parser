package service

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/KaymeKaydex/bmstu-fv-parser.git/internal/app/model"
	"github.com/KaymeKaydex/bmstu-fv-parser.git/internal/pkg/clients/fv"
)

type Service struct {
	fvClient *fv.Client
}

func New(ctx context.Context) (*Service, error) {
	fvClient := fv.New(ctx)

	return &Service{
		fvClient: fvClient,
	}, nil
}

func (s *Service) ParseFVSite(ctx context.Context, date time.Time) ([]model.WorkingOutItem, error) {
	_, err := s.fvClient.GetWorkingOut(fv.RequestGetWorkingOut{
		Id:            14,
		Date:          time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local),
		SecurityLSKey: "bb536826f9118d389119e1f36a2e208a",
	})
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("cant do request to fv")
	}

	return nil, nil
}

func (s *Service) WriteWorkingOutItems(ctx context.Context, items []model.WorkingOutItem) error {
	return nil
}

/*
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		log.WithError(err).Println("Cant open postgers connection")

		return err
	}
*/
