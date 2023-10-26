package server

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"strings"

	"github.com/christiandwi/edot/order-service/response"
	"github.com/christiandwi/edot/order-service/usecase/order"
	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	service order.Service
}

func newOrderHandler(service order.Service) *orderHandler {
	return &orderHandler{service: service}
}

type orderRequest struct {
	ProductID   string  `json:"product_id"`
	WarehouseID string  `json:"warehouse_id"`
	Amount      float64 `json:"amount"`
}

func (g *orderHandler) Checkout() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		bodyAsByteArray, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			response.SetResponse(ctx, false, 400, nil, err)
		}

		var orderReq orderRequest
		json.Unmarshal(bodyAsByteArray, &orderReq)

		token := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]
		userId, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			response.SetResponse(ctx, false, 401, nil, err)
		}

		err = g.service.Checkout(orderReq.ProductID, orderReq.WarehouseID, string(userId), orderReq.Amount)
		if err != nil {
			response.SetResponse(ctx, false, 500, nil, err)
		}

		response.SetResponse(ctx, true, 200, map[string]interface{}{"success": true}, nil)
	}
}
