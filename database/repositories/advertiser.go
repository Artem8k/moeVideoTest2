package repositories

import (
	"context"
	"fmt"
	"test/database/models"

	"github.com/jmoiron/sqlx"
)

type AdvertiserRepository interface {
	GetAll(ctx context.Context) ([]*models.Advertiser, error)
	Get(id int, ctx context.Context) (*models.Advertiser, error)
	Add(advertiser models.Advertiser, ctx context.Context) error
	Update(advertiser models.AdvertiserUpdate, ctx context.Context) error
	Delete(id int, ctx context.Context) error
}

type advertiserRepository struct {
	db *sqlx.DB
}

func NewAdvertiserRepository(db *sqlx.DB) AdvertiserRepository {
	return &advertiserRepository{
		db: db,
	}
}

func (a *advertiserRepository) GetAll(ctx context.Context) ([]*models.Advertiser, error) {
	adv := []*models.Advertiser{}
	err := a.db.SelectContext(ctx, &adv, "SELECT * FROM advertiser")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return adv, err
}

func (a *advertiserRepository) Get(id int, ctx context.Context) (*models.Advertiser, error) {
	adv := models.Advertiser{}
	err := a.db.GetContext(ctx, &adv, "SELECT * FROM advertiser WHERE id=$1", id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &adv, nil
}

func (a *advertiserRepository) Add(advertiser models.Advertiser, ctx context.Context) error {
	tx, err := a.db.BeginTx(ctx, nil)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(
		"INSERT INTO advertiser (title, create_date) VALUES ($1, $2)",
		advertiser.Title, advertiser.Create_date,
	)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (a *advertiserRepository) Update(advertiser models.AdvertiserUpdate, ctx context.Context) error {
	tx, err := a.db.BeginTx(ctx, nil)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(
		"UPDATE advertiser SET title=$1 WHERE id=$2",
		advertiser.Title, advertiser.Id,
	)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (a *advertiserRepository) Delete(id int, ctx context.Context) error {
	tx, err := a.db.BeginTx(ctx, nil)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(
		"DELETE FROM advertiser WHERE id=$1",
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
