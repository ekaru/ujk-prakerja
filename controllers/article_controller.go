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

func GetAllArticle(ctx echo.Context) error {
	var data []models.Article
	var articleRes []models.ArticleResponse

	if err := configs.DB.Joins("User").Joins("Category").Find(&data).Find(&articleRes).Error; err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.ResponseDefault(http.StatusInternalServerError, err.Error(), nil),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.ResponseDefault(http.StatusOK, "Success", articleRes),
	)
}

func GetOneArticle(ctx echo.Context) error {
	var id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.ResponseDefault(http.StatusBadRequest, "Invalid input.", nil),
		)
	}

	var article models.Article
	var articleRes models.ArticleResponse

	if err := configs.DB.Joins("User").Joins("Category").First(&article, id).First(&articleRes).Error; err != nil {
		log.Println("err:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(
				http.StatusNotFound,
				helpers.ResponseDefault(http.StatusNotFound, "Article not found. Invalid ID", nil),
			)
		}

		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.ResponseDefault(http.StatusInternalServerError, err.Error(), nil),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.ResponseDefault(http.StatusOK, "Success", articleRes),
	)
}

func CreateArticle(ctx echo.Context) error {
	var article models.Article
	err := ctx.Bind(&article)
	if err != nil {
		log.Println("err1:", err)
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.ResponseDefault(http.StatusBadRequest, "Invalid input.", nil),
		)
	}

	validate := validator.New()

	err = validate.Struct(article)
	if err != nil {
		log.Println("err2:", err)
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.ResponseDefault(http.StatusBadRequest, "Invalid input.", nil),
		)
	}

	var user models.User

	if err := configs.DB.First(&user, article.UserID).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(
				http.StatusNotFound,
				helpers.ResponseDefault(http.StatusNotFound, "User not found. Invalid UserID", nil),
			)
		}

		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.ResponseDefault(http.StatusInternalServerError, err.Error(), nil),
		)
	}

	var category models.Category

	if err := configs.DB.First(&category, article.CategoryID).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(
				http.StatusNotFound,
				helpers.ResponseDefault(http.StatusNotFound, "Category not found. Invalid CategoryID", nil),
			)
		}

		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.ResponseDefault(http.StatusInternalServerError, err.Error(), nil),
		)
	}

	if err := configs.DB.Create(&article).Error; err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.ResponseDefault(http.StatusInternalServerError, err.Error(), nil),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.ResponseDefault(http.StatusCreated, "Success", article),
	)
}

func UpdateArticle(ctx echo.Context) error {
	var id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.ResponseDefault(http.StatusBadRequest, "Invalid input.", nil),
		)
	}

	var article, updatedArticle models.Article

	if err := configs.DB.First(&article, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(
				http.StatusNotFound,
				helpers.ResponseDefault(http.StatusNotFound, "Article not found. Invalid ID", nil),
			)
		}

		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.ResponseDefault(http.StatusInternalServerError, err.Error(), nil),
		)
	}

	if err := ctx.Bind(&updatedArticle); err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.ResponseDefault(http.StatusInternalServerError, err.Error(), nil),
		)
	}

	validate := validator.New()

	err = validate.Struct(updatedArticle)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.ResponseDefault(http.StatusBadRequest, "Invalid input.", nil),
		)
	}

	var user models.User
	if article.UserID != 0 {
		if err := configs.DB.First(&user, article.UserID).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return ctx.JSON(
					http.StatusNotFound,
					helpers.ResponseDefault(http.StatusNotFound, "User not found. Invalid UserID", nil),
				)
			}

			return ctx.JSON(
				http.StatusInternalServerError,
				helpers.ResponseDefault(http.StatusInternalServerError, err.Error(), nil),
			)
		}
	}

	var category models.Category
	if article.CategoryID != 0 {
		if err := configs.DB.First(&category, article.CategoryID).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return ctx.JSON(
					http.StatusNotFound,
					helpers.ResponseDefault(http.StatusNotFound, "Category not found. Invalid CategoryID", nil),
				)
			}

			return ctx.JSON(
				http.StatusInternalServerError,
				helpers.ResponseDefault(http.StatusInternalServerError, err.Error(), nil),
			)
		}
	}

	if err := configs.DB.Where("id = ?", id).Updates(&updatedArticle).Error; err != nil {
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

func DeleteArticle(ctx echo.Context) error {
	var id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.ResponseDefault(http.StatusBadRequest, "Invalid input.", nil),
		)
	}

	var article models.Article

	if err := configs.DB.First(&article, id).Delete(&article).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(
				http.StatusNotFound,
				helpers.ResponseDefault(http.StatusNotFound, "Article not found. Invalid ID", nil),
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
