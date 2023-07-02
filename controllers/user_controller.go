package controllers

import (
	"blog-app/configs"
	"blog-app/helpers"
	"blog-app/models"
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(ctx echo.Context) error {
	var user models.User

	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.ResponseDefault(http.StatusBadRequest, "Invalid input.", nil),
		)
	}

	validate := validator.New()

	err = validate.Struct(user)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.ResponseDefault(http.StatusBadRequest, "Invalid input.", nil),
		)
	}

	checkEmail, _ := getByEmail(user.Email)
	if checkEmail {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.ResponseDefault(http.StatusBadRequest, "Email has already been used.", nil),
		)
	}

	hashedPassword, err := helpers.Hash(user.Password)
	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.ResponseDefault(http.StatusInternalServerError, err.Error(), nil),
		)
	}

	user.Password = string(hashedPassword)

	if err := configs.DB.Create(&user).Error; err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			helpers.ResponseDefault(http.StatusInternalServerError, err.Error(), nil),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.ResponseDefault(http.StatusCreated, "Success", user),
	)
}

func Login(ctx echo.Context) error {
	var user models.UserLogin

	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.ResponseDefault(http.StatusBadRequest, "Invalid input.", nil),
		)
	}

	validate := validator.New()

	err = validate.Struct(user)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.ResponseDefault(http.StatusBadRequest, "Invalid input.", nil),
		)
	}

	data, err := signIn(user.Email, user.Password)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(
				http.StatusBadRequest,
				helpers.ResponseDefault(http.StatusBadRequest, "Email not found.", nil),
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

func signIn(email, password string) (interface{}, error) {
	var (
		user models.User
		err  error
	)

	result := configs.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", gorm.ErrRecordNotFound
		}

		return "", result.Error
	}

	log.Println("result", result.RowsAffected)

	err = helpers.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err2 := helpers.GenerateToken(user.Id, user.Username)

	data := map[string]string{
		"username": user.Username,
		"token":    token,
	}

	return data, err2

}

func getByEmail(email string) (bool, error) {
	var data models.User

	result := configs.DB.First(&data, "email = ?", email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, errors.New("email not found")
		}

		return false, result.Error
	}

	return true, result.Error
}
