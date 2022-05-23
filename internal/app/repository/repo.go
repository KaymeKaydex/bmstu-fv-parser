package repository

import (
	"context"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/KaymeKaydex/bmstu-fv-parser.git/internal/app/dsn"
	"github.com/KaymeKaydex/bmstu-fv-parser.git/internal/app/model"
)

type Repository struct {
	// корневой контекст
	ctx context.Context

	db *gorm.DB
}

func New(ctx context.Context) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		log.WithError(err).Println("Cant open postgers connection")

		return nil, err
	}

	return &Repository{
		ctx: ctx,
		db:  db,
	}, nil
}

func (r *Repository) RewriteWorkingOutItems(ctx context.Context, items []model.WorkingOutItem) error {
	var err error

	if len(items) == 0 {
		return fmt.Errorf("zero len slice")
	}

	year, month, d := items[0].CreatedAt.Date()

	sqlDate := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(d)

	err = r.db.Where("created_at = ?", sqlDate).Delete(&items[0]).Error
	if err != nil {
		log.WithError(err).Error("cant delete days")
	}
	for _, item := range items {
		err := r.db.Create(&item).Error
		if err != nil {
			log.WithError(err).Error("cant create item")
		}
	}

	return nil
}
