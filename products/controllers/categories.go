package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"net/http"
	"point-of-sale.go/v1/internal/db/repository"
	"point-of-sale.go/v1/internal/types"
	"point-of-sale.go/v1/internal/web"
)

type getCategoryRequest struct {
	Id types.ID `json:"id" validate:"required" msgpack:"id"`
}

func GetCategoryEndpoint(app *web.App, r *http.Request) *web.APIResponse {
	data := &getCategoryRequest{}
	err := app.Validate(data, r, false, false, true, false)
	if err != nil {
		return web.NewErrorAPIResponse(fmt.Errorf("invalid payload %w", err), 400)
	}

	repo := repository.New(app.DB())

	if err != nil {
		return web.NewErrorAPIResponse(err, 500)
	}

	category, err := repo.GetProductCategory(r.Context(), data.Id.Int64())

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return web.NewErrorAPIResponse(err, 404)
		}
		return web.NewErrorAPIResponse(err, 500)
	}

	return web.NewAPIResponse(category, 200)
}

type createCategoryRequest struct {
	Name     string        `json:"name" validate:"required"`
	ParentID sql.NullInt64 `json:"parent_id"`
}

func CreateCategoryEndpoint(app *web.App, r *http.Request) *web.APIResponse {
	data := &createCategoryRequest{}
	err := app.Validate(data, r, true, false, true, false)
	if err != nil {
		return web.NewErrorAPIResponse(fmt.Errorf("invalid payload %w", err), 400)
	}

	repo := repository.New(app.DB())
	ProductCategory, err := repo.CreateProductCategory(r.Context(), repository.CreateProductCategoryParams{
		Name:     data.Name,
		ParentID: data.ParentID,
	})

	if err != nil {
		return web.NewErrorAPIResponse(err, 500)
	}

	return web.NewAPIResponse(ProductCategory, 201)
}

func ListCategoriesEndpoint(app *web.App, r *http.Request) *web.APIResponse {
	repo := repository.New(app.DB())
	categories, err := repo.ListProductCategories(r.Context())

	if err != nil {
		return web.NewErrorAPIResponse(err, 500)
	}

	return web.NewAPIResponse(categories, 200)
}

type deleteCategoryRequest struct {
	Id int64 `json:"id,string" validate:"required"`
}

func DeleteCategoryEndpoint(app *web.App, r *http.Request) *web.APIResponse {
	data := &deleteCategoryRequest{}
	err := app.Validate(data, r, false, false, true, false)
	if err != nil {
		return web.NewErrorAPIResponse(fmt.Errorf("invalid payload %w", err), 400)
	}

	repo := repository.New(app.DB())

	err = repo.DeleteProductCategory(r.Context(), data.Id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return web.NewErrorAPIResponse(err, 404)
		}
		return web.NewErrorAPIResponse(err, 500)
	}

	return web.NewAPIResponse(map[string]interface{}{
		"category_id": data.Id,
	}, 200)

}
