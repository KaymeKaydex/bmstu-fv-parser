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
	c := fv.New(ctx)
	c.GetWorkingOut("")

	return nil
}
