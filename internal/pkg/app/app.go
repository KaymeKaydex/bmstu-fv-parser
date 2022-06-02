package app

import (
	"context"
)

type App struct {
	// корневой контекст
	ctx context.Context
}

func New(ctx context.Context) (*App, error) {
	app := &App{
		ctx: ctx,
	}

	return app, nil
}

func (a *App) Run(ctx context.Context) error {

	return nil
}
