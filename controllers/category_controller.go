package controllers

import (
	"blog-app/configs"
	"blog-app/helpers"
	"blog-app/models"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllCategory(ctx echo.Context) error {
	var data []models.Category

	if err := configs.DB.Find(&data).Error; err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.ResponseDefault(http.StatusInternalServerError, err.Error(), nil),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.ResponseDefault(http.StatusOK, "Success", data),
	)
}

func GetOneCategory(ctx echo.Context) error {
	var id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.ResponseDefault(http.StatusBadRequest, "Invalid input.", nil),
		)
	}

	var data models.Category

	if err := configs.DB.First(&data, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(
				http.StatusNotFound,
				helpers.ResponseDefault(http.StatusNotFound, "Category not found. Invalid ID", nil),
			)
		}

		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.ResponseDefault(http.StatusInternalServerError, err.Error(), nil),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.ResponseDefault(http.StatusOK, "Success", data),
	)
}

func CreateCategory(ctx echo.Context) error {
	var category models.Category
	err := ctx.Bind(&category)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.ResponseDefault(http.StatusBadRequest, "Invalid input.", nil),
		)
	}

	validate := validator.New()

	err = validate.Struct(category)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.ResponseDefault(http.StatusBadRequest, "Invalid input.", nil),
		)
	}

	if err := configs.DB.Create(&category).Error; err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.ResponseDefault(http.StatusInternalServerError, err.Error(), nil),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.ResponseDefault(http.StatusOK, "Success", category),
	)
}

func UpdateCategory(ctx echo.Context) error {
	var id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.ResponseDefault(http.StatusBadRequest, "Invalid input.", nil),
		)
	}

	var category, updatedCategory models.Category

	if err := configs.DB.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(
				http.StatusNotFound,
				helpers.ResponseDefault(http.StatusNotFound, "Category not found. Invalid ID", nil),
			)
		}

		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.ResponseDefault(http.StatusInternalServerError, err.Error(), nil),
		)
	}

	if err := ctx.Bind(&updatedCategory); err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.ResponseDefault(http.StatusInternalServerError, err.Error(), nil),
		)
	}
	log.Println("category", updatedCategory)

	validate := validator.New()

	err = validate.Struct(updatedCategory)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.ResponseDefault(http.StatusBadRequest, "Invalid input.", nil),
		)
	}

	if err := configs.DB.Where("id = ?", id).Updates(&updatedCategory).Error; err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.ResponseDefault(http.StatusInternalServerError, err.Error(), nil),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.ResponseDefault(http.StatusOK, "Success", nil),
	)
}

func DeleteCategory(ctx echo.Context) error {
	var id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.ResponseDefault(http.StatusBadRequest, "Invalid input.", nil),
		)
	}

	var data models.Category

	if err := configs.DB.Where("id = ?", id).Delete(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(
				http.StatusNotFound,
				helpers.ResponseDefault(http.StatusNotFound, "Category not found. Invalid ID", nil),
			)
		}

		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.ResponseDefault(http.StatusInternalServerError, err.Error(), nil),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.ResponseDefault(http.StatusOK, "Success", nil),
	)
}
