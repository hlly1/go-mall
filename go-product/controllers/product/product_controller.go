package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
}

func (con ProductController) Index(ctx *gin.Context) {
	product := map[string]any{"name": "apple", "price": 6, "unit":"/kg", "stock":10, "stock_unit" : "kg"}
	ctx.JSON(http.StatusOK, gin.H{"respCode": "0000", "respDesc": "success", "product": product})
}

func (con ProductController) Add(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"respCode": "0000", "respDesc": "add succeed"})
}

func (con ProductController) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"respCode": "0000", "respDesc": "update succeed"})
}

func (con ProductController) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"respCode": "0000", "respDesc": "delete succeed"})
}
