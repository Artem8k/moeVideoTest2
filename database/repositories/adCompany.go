package repositories

import (
	"context"
	"fmt"
	"test/database/models"

	"github.com/jmoiron/sqlx"
)

type AdCompanyRepository interface {
	Get(id int, ctx context.Context) (*models.AdCompany, error)
	Add(advertisement models.AdCompany, ctx context.Context) error
	Update(advertisement models.AdCompanyUpdate, ctx context.Context) error
	Delete(id int, ctx context.Context) error
}

type adCompanyRepository struct {
	db *sqlx.DB
}

func NewAdCompanyRepository(db *sqlx.DB) AdCompanyRepository {
	return &adCompanyRepository{
		db: db,
	}
}

func (a *adCompanyRepository) Get(id int, ctx context.Context) (*models.AdCompany, error) {
	adc := models.AdCompany{}
	err := a.db.GetContext(ctx, &adc, `SELECT * FROM "adCompany" WHERE id=$1`, id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &adc, nil
}

func (a *adCompanyRepository) Add(advertisement models.AdCompany, ctx context.Context) error {
	tx, err := a.db.BeginTx(ctx, nil)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(
		`INSERT INTO "adCompany" (title, create_date, advertiser_id) VALUES ($1, $2, $3)`,
		advertisement.Title, advertisement.Create_date, advertisement.Advertiser_id,
	)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (a *adCompanyRepository) Update(advertisement models.AdCompanyUpdate, ctx context.Context) error {
	tx, err := a.db.BeginTx(ctx, nil)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(
		`UPDATE "adCompany" SET title=$1 WHERE id=$2`,
		advertisement.Title, advertisement.Id,
	)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (a *adCompanyRepository) Delete(id int, ctx context.Context) error {
	tx, err := a.db.BeginTx(ctx, nil)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(
		`DELETE FROM "adCompany" WHERE id=$1`,
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
