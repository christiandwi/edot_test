package server

import (
	"fmt"

	"github.com/christiandwi/go-asynq/response"
	"github.com/christiandwi/go-asynq/usecase/guest"
	"github.com/gin-gonic/gin"
)

type guestHandler struct {
	service guest.Service
}

func newGuestHandler(service guest.Service) *guestHandler {
	return &guestHandler{service: service}
}

func (g *guestHandler) Hello() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if ctx.Query("name") == "error" {
			response.SetResponse(ctx, false, 500, nil, fmt.Errorf("error hit hello"))
			return
		}

		resp := map[string]interface{}{"name": ctx.Query("name")}
		response.SetResponse(ctx, true, 200, resp, nil)
	}
}

func (g *guestHandler) Service() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		resp := g.service.HelloService(ctx.Query("name"))
		response.SetResponse(ctx, true, 200, resp, nil)
	}
}
