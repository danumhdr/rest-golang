package helper

import (
	"golearning/restapi/model/domain"
	"golearning/restapi/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(list []domain.Category) []web.CategoryResponse {
	var listResp []web.CategoryResponse
	for _, category := range list {
		listResp = append(listResp, ToCategoryResponse(category))
	}
	return listResp
}
