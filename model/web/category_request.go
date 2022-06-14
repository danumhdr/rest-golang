package web

type CategoryCreateRequest struct {
	Name string `validate:"required"`
}

type CategoryUpdateRequest struct {
	Id   int
	Name string `validate:"required"`
}
