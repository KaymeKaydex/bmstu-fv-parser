package app

import (
	"context"

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
	c.GetWorkingOut(fv.RequestGetWorkingOut{})

	return nil
}
