package repository

import (
	"context"
	"database/sql"
	"golearning/restapi/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, trx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, trx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, trx *sql.Tx, category domain.Category)
	FindbyId(ctx context.Context, trx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, trx *sql.Tx) []domain.Category
}
