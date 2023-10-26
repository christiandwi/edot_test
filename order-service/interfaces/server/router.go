package server

import "github.com/gin-gonic/gin"

func setupRouter(app *gin.Engine, handler *handler) {
	appRouter := app.Group("/api")
	v1 := appRouter.Group("/v1")

	//Order grouping
	order := v1.Group("/orders")
	order.POST("", handler.orderHandler.Checkout())
}
