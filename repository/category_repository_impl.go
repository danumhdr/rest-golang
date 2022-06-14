package repository

import (
	"context"
	"database/sql"
	"errors"
	"golearning/restapi/helper"
	"golearning/restapi/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, trx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into category(Name) values (?)"
	result, err := trx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, trx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set Name = ? where Id = ?"
	_, err := trx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, trx *sql.Tx, category domain.Category) {
	SQL := "delete from category where Id = ?"
	_, err := trx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindbyId(ctx context.Context, trx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select Id,Name from category where Id = ?"
	result, err := trx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer result.Close()

	category := domain.Category{}
	if result.Next() {
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("Category Not Found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, trx *sql.Tx) []domain.Category {
	SQL := "select Id,Name from category"
	result, err := trx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer result.Close()

	var categories []domain.Category
	for result.Next() {
		category := domain.Category{}
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
