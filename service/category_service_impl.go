package service

import (
	"context"
	"database/sql"
	"golearning/restapi/exception"
	"golearning/restapi/helper"
	"golearning/restapi/model/domain"
	"golearning/restapi/model/web"
	"golearning/restapi/repository"

	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, req web.CategoryCreateRequest) web.CategoryResponse {
	errorValidate := service.Validate.Struct(req)
	helper.PanicIfError(errorValidate)

	trx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(trx)

	category := domain.Category{
		Name: req.Name,
	}

	category = service.CategoryRepository.Save(ctx, trx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, req web.CategoryUpdateRequest) web.CategoryResponse {
	errorValidate := service.Validate.Struct(req)
	helper.PanicIfError(errorValidate)
	trx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(trx)

	category, err := service.CategoryRepository.FindbyId(ctx, trx, req.Id)
	if err != nil {
		panic(exception.NewError404(err.Error()))
	}

	category.Name = req.Name

	category = service.CategoryRepository.Update(ctx, trx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	trx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(trx)

	category, err := service.CategoryRepository.FindbyId(ctx, trx, categoryId)
	if err != nil {
		panic(exception.NewError404(err.Error()))
	}

	category = domain.Category{
		Id: categoryId,
	}

	service.CategoryRepository.Delete(ctx, trx, category)

}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	trx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(trx)

	id, err := service.CategoryRepository.FindbyId(ctx, trx, categoryId)
	if err != nil {
		panic(exception.NewError404(err.Error()))
	}

	return helper.ToCategoryResponse(id)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	trx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(trx)

	list := service.CategoryRepository.FindAll(ctx, trx)

	return helper.ToCategoryResponses(list)
}
