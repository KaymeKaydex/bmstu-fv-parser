package enegine

import (
	"context"
)

type Engine struct {
}

func New(ctx context.Context) *Engine {
	return &Engine{}
}

func (e *Engine) Start() error {
	return nil
}
