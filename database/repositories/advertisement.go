package repositories

import (
	"context"
	"fmt"
	"test/database/models"

	"github.com/jmoiron/sqlx"
)

type AdvertisementRepository interface {
	Get(id int, ctx context.Context) (*models.Advertisement, error)
	Add(advertisement models.Advertisement, ctx context.Context) error
	Update(advertisement models.AdvertisementUpdate, ctx context.Context) error
	Delete(id int, ctx context.Context) error
}

type advertisementRepository struct {
	db *sqlx.DB
}

func NewAdvertisementRepository(db *sqlx.DB) AdvertisementRepository {
	return &advertisementRepository{
		db: db,
	}
}

func (a *advertisementRepository) Get(id int, ctx context.Context) (*models.Advertisement, error) {
	adv := models.Advertisement{}
	err := a.db.GetContext(ctx, &adv, "SELECT * FROM advertisement WHERE id=$1", id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &adv, nil
}

func (a *advertisementRepository) Add(advertisement models.Advertisement, ctx context.Context) error {
	tx, err := a.db.BeginTx(ctx, nil)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(
		"INSERT INTO advertisement (title, CPM, AdCompany_id, create_date) VALUES ($1, $2, $3, $4)",
		advertisement.Title, advertisement.CPM, advertisement.AdCompany_id, advertisement.Create_date,
	)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (a *advertisementRepository) Update(advertisement models.AdvertisementUpdate, ctx context.Context) error {
	tx, err := a.db.BeginTx(ctx, nil)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(
		"UPDATE advertisement SET title=$1, CPM=$2 WHERE id=$3",
		advertisement.Title, advertisement.CPM, advertisement.Id,
	)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (a *advertisementRepository) Delete(id int, ctx context.Context) error {
	tx, err := a.db.BeginTx(ctx, nil)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(
		"DELETE FROM advertisement WHERE id=$1",
		id,
	)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
