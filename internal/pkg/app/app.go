package app

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/KaymeKaydex/bmstu-fv-parser.git/internal/pkg/clients/fv"
)

type App struct {
	// корневой контекст
	ctx context.Context
}

func New(ctx context.Context) (*App, error) {
	return &App{}, nil
}

func (a *App) Run(ctx context.Context) error {
	/*
		на след занятие

		db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
		if err != nil {
			log.WithError(err).Println("Cant open postgers connection")

			return err
		}
	*/

	c := fv.New(ctx)
	resp, err := c.GetWorkingOut(fv.RequestGetWorkingOut{
		Id:            14,
		Date:          time.Date(2022, 05, 17, 0, 0, 0, 0, time.Local),
		SecurityLSKey: "bb536826f9118d389119e1f36a2e208a",
	})
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("cant do request to fv")
	}

	fmt.Println(resp)

	return nil
}
