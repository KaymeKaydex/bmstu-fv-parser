package app

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/KaymeKaydex/bmstu-fv-parser.git/internal/app/config"
	"github.com/KaymeKaydex/bmstu-fv-parser.git/internal/app/model"
	"github.com/KaymeKaydex/bmstu-fv-parser.git/internal/app/service"
)

type IService interface {
	ParseFVSite(ctx context.Context, date time.Time) ([]model.WorkingOutItem, error)
	// WriteWorkingOutItems Запись в базу данных
	WriteWorkingOutItems(ctx context.Context, items []model.WorkingOutItem) error
	// ReadWorkingOutItems(ctx context.Context, date time.Time) []model.WorkingOutItem
}

type App struct {
	// корневой контекст
	ctx context.Context

	service IService
}

func New(ctx context.Context) (*App, error) {
	app := &App{
		ctx: ctx,
	}

	srv, err := service.New(ctx)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("cant create service")
	}

	app.service = srv

	return app, nil
}

func (a *App) Run(ctx context.Context) error {
	fvCfg := config.FromContext(ctx).FVConfig

	for {
		// дата за которую будет парсинг
		date := time.Now()

		for i := 0; i < 3; i++ {
			workingOutItems, err := a.service.ParseFVSite(ctx, date)
			if err != nil {
				log.WithError(err).Error("cant parse fv")
			}

			err = a.service.WriteWorkingOutItems(ctx, workingOutItems)
			if err != nil {
				log.WithError(err).Error("cant write to db")
			}

			// следующий день
			date = date.Add(time.Hour * 24)
		}

		time.Sleep(fvCfg.CronTimeout)
	}

	return nil
}
