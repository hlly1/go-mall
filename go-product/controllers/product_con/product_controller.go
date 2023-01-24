package productcon

import (
	"strings"

	"github.com/gin-gonic/gin"

	"go-product/dao"
	"go-product/entities"

	basecon "go-product/controllers/base_con"

	"github.com/google/uuid"
)

type ProductController struct {
	basecon.BaseController
}

func (con ProductController) Index(ctx *gin.Context) {
	product := basecon.Tuple{Key: "data", Value: map[string]any{"name": "apple", "price": 6, "unit": "/kg", "stock": 10, "stock_unit": "kg"}}
	basecon.ReturnSucceedWithArgus(ctx, product)
}

func (con ProductController) Add(ctx *gin.Context) {
	product := entities.Product{}
	ctx.ShouldBindJSON(&product)
	product.ID = strings.Replace(uuid.NewString(), "-", "", -1)
	dao.DB.Create(product)
	basecon.ReturnSucceed(ctx)
}

func (con ProductController) Update(ctx *gin.Context) {
	basecon.ReturnSucceed(ctx)
}

func (con ProductController) Delete(ctx *gin.Context) {
	basecon.ReturnSucceed(ctx)
}
