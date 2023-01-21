package httpmeth

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"wb/Using"
)

type HandlersOrd struct {
	Group        *gin.RouterGroup
	OrderUseCase Using.UseCase
}

func OrderHandlers(group *gin.RouterGroup, orderUC Using.UseCase) *HandlersOrd {
	return &HandlersOrd{Group: group, OrderUseCase: orderUC}
}

func (HandO *HandlersOrd) GetByUID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		orderUID := ctx.Param("order_uid")
		if orderUID == "" {
			log.Print("order_uid is empty")
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "empty value not accepted"})
		} else {
			m, err := HandO.OrderUseCase.GetByUid(ctx, "orderUID")
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{"msg": "order with UID value not found"})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"status": 200, "data": m})
			}
		}
	}
}

func (HandO *HandlersOrd) MapRoutes() {
	HandO.Group.GET("/:order_uid", HandO.GetByUID())
}
