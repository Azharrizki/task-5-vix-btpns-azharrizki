package router

import (
	"task-5-vix-btpns/controllers/photocontroller"
	"task-5-vix-btpns/controllers/usercontroller"
	"task-5-vix-btpns/middlewares"

	"github.com/gin-gonic/gin"
)

func Routers() {
	r := gin.Default()

	r.POST("/users/register", usercontroller.Register)
	r.GET("/users/login", usercontroller.Login)
	r.PUT("/users/:userId", usercontroller.Update)
	r.DELETE("/users/:userId", usercontroller.Delete)

	authGroup := r.Group("/auth")
	authGroup.Use(middlewares.AuthMiddleware())

	authGroup.GET("/photos", photocontroller.Index)
	authGroup.POST("/photos", photocontroller.Create)
	authGroup.PUT("/photos/:photoId", photocontroller.Update)
	authGroup.DELETE("/photos/:photoId", photocontroller.Delete)

	r.Run()
}

func AuthMiddleware() {
	panic("unimplemented")
}
