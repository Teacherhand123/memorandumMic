package router

import (
	"micro-memorandum/app/gateway/http"
	"micro-memorandum/app/gateway/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	apiV1 := ginRouter.Group("/api/v1")
	{
		apiV1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "success"})
		})

		apiV1.POST("/user/register", http.UserRegisterHandler)
		apiV1.POST("/user/login", http.UserLoginHandler)

		authed := apiV1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("task", http.CreateTaskHandler)
			authed.POST("update_task", http.UpdateTaskHandler)
			authed.POST("delete_task", http.DeleteTaskHandler)
			authed.GET("get_task", http.GetTaskHandler)
			authed.GET("get_tasks", http.ListTaskHandler)
		}

	}

	return ginRouter
}
