package routes

import (
	"blog-app/controllers"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) *echo.Echo {
	ctg := e.Group("/categories")
	{
		ctg.GET("", controllers.GetAllCategory)
		ctg.GET("/:id", controllers.GetOneCategory)
		ctg.POST("", controllers.CreateCategory)
		ctg.PUT("/:id", controllers.UpdateCategory)
		ctg.DELETE("/:id", controllers.DeleteCategory)
	}

	usg := e.Group("/users")
	{
		usg.POST("/register", controllers.CreateUser)
		usg.POST("/login", controllers.Login)
	}

	arg := e.Group("/articles")
	{
		arg.GET("", controllers.GetAllArticle)
		arg.GET("/:id", controllers.GetOneArticle)
		arg.POST("", controllers.CreateArticle)
		arg.PUT("/:id", controllers.UpdateArticle)
		arg.DELETE("/:id", controllers.DeleteArticle)
	}

	// cmg := e.Group("/comments")
	// {
	// 	cmg.GET("", controllers.GetAllCategory)
	// 	cmg.GET("/:id", controllers.GetOneCategory)
	// 	cmg.POST("", controllers.CreateCategory)
	// 	cmg.PUT("/:id", controllers.UpdateCategory)
	// 	cmg.DELETE("/:id", controllers.DeleteCategory)
	// }

	return e
}
