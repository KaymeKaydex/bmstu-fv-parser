package main

import (
	"context"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/KaymeKaydex/bmstu-fv-parser.git/internal/app/config"
	"github.com/KaymeKaydex/bmstu-fv-parser.git/internal/pkg/app"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	ctx := context.Background()

	cfg, err := config.NewConfig(ctx)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("cant init config")

		os.Exit(2)
	}

	ctx = config.WrapContext(ctx, cfg)

	fmt.Println(cfg)

	application, err := app.New(ctx)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("cant create application")

		os.Exit(2)
	}

	err = application.Run()
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("cant run application")

		os.Exit(2)
	}
}
