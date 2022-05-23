package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/KaymeKaydex/bmstu-fv-parser.git/internal/app/model"
	"github.com/KaymeKaydex/bmstu-fv-parser.git/internal/app/repository"
	"github.com/KaymeKaydex/bmstu-fv-parser.git/internal/pkg/clients/fv"
)

type Service struct {
	fvClient *fv.Client

	repo *repository.Repository
}

func New(ctx context.Context) (*Service, error) {
	fvClient := fv.New(ctx)

	repo, err := repository.New(ctx)
	if err != nil {
		return nil, err
	}

	return &Service{
		fvClient: fvClient,
		repo:     repo,
	}, nil
}

func (s *Service) ParseFVSite(ctx context.Context, date time.Time) ([]model.WorkingOutItem, error) {
	resp, err := s.fvClient.GetWorkingOut(fv.RequestGetWorkingOut{
		Id:            14,
		Date:          time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local),
		SecurityLSKey: "bb536826f9118d389119e1f36a2e208a",
	})
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("cant do request to fv")

		return nil, err
	}

	workingOutItems := make([]model.WorkingOutItem, len(resp.Items), len(resp.Items))

	for i, elem := range resp.Items {
		item := model.WorkingOutItem{}
		item.CreatedAt = date
		item.Title = elem.Title
		item.Description = elem.Description
		item.AddressTitle = elem.Address.Title
		item.IsAllowJoin = elem.IsAllowJoin
		item.Url = elem.Url
		item.Address = elem.Address.Address

		item.Id, err = strconv.Atoi(elem.Id)

		item.DateBegin, err = s.tryParseTime(elem.DateBegin)
		if err != nil {
			log.WithError(err).Error("cant parse date begin")
		}

		item.DateEnd, err = s.tryParseTime(elem.DateEnd)
		if err != nil {
			log.WithError(err).Error("cant parse date begin")
		}

		item.DateRegistrationEnd, err = s.tryParseTime(elem.DateRegistrationEnd)
		if err != nil {
			log.WithError(err).Error("cant parse date begin")
		}

		workingOutItems[i] = item
	}

	return workingOutItems, nil
}

func (s *Service) tryParseTime(t interface{}) (time.Time, error) {
	switch t.(type) {
	case string:
		stringTime := t.(string)
		res, err := time.Parse("2006-01-02 15:04", stringTime)
		if err != nil {
			return time.Time{}, err
		}

		return res, nil
	default:
		return time.Time{}, fmt.Errorf("smth vent wrong, not string type: %t", t)
	}
}

func (s *Service) WriteWorkingOutItems(ctx context.Context, items []model.WorkingOutItem) error {
	return s.repo.RewriteWorkingOutItems(ctx, items)
}
