package server

import (
	"encoding/json"
	"io"

	"github.com/christiandwi/edot/user-service/response"
	"github.com/christiandwi/edot/user-service/usecase/guest"
	"github.com/gin-gonic/gin"
)

type guestHandler struct {
	service guest.Service
}

func newGuestHandler(service guest.Service) *guestHandler {
	return &guestHandler{service: service}
}

func (g *guestHandler) Login() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		bodyAsByteArray, _ := io.ReadAll(ctx.Request.Body)
		jsonMap := make(map[string]string)
		json.Unmarshal(bodyAsByteArray, &jsonMap)

		resp, err := g.service.LoginService(jsonMap["username"], jsonMap["password"])
		if err != nil {
			response.SetResponse(ctx, false, 400, nil, err)
		}

		tokenMap := make(map[string]interface{})
		tokenMap["token"] = resp

		response.SetResponse(ctx, true, 200, tokenMap, nil)
	}
}
