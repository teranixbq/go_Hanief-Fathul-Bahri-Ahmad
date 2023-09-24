package route

import (
	"github.com/labstack/echo"
	"praktikum/rest/controller"
)

func BlogsRoute(e *echo.Echo) {
	blogs := e.Group("/blogs")
	blogs.POST("", controller.CreateBlogsController)
	blogs.GET("", controller.GetAllBlogsController)
	blogs.GET("/:id", controller.GetBlogsByIdController)
	blogs.PUT("/:id", controller.UpdateBlogsController)
	blogs.DELETE("/:id", controller.DeleteBlogsByIdController)
}
