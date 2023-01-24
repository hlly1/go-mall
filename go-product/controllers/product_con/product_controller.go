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
	rows, err := dao.DB.Model(&entities.Product{}).Raw("select * from product where id = ?", ctx.Query("id")).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var product entities.Product
	for rows.Next() {
		dao.DB.ScanRows(rows, &product)
	}
	basecon.ReturnSucceedWithArgus(ctx, basecon.Tuple{Key: "product", Value: product})
}

func (con ProductController) Add(ctx *gin.Context) {
	product := entities.Product{}
	ctx.ShouldBindJSON(&product)
	product.ID = strings.Replace(uuid.NewString(), "-", "", -1)
	dao.DB.Create(product)
	basecon.ReturnSucceed(ctx)
}

func (con ProductController) Update(ctx *gin.Context) {
	product := entities.Product{}
	ctx.ShouldBindJSON(&product)
	dao.DB.Updates(product)
	basecon.ReturnSucceed(ctx)
}

func (con ProductController) Delete(ctx *gin.Context) {
	product := entities.Product{}
	ctx.ShouldBindJSON(&product)
	dao.DB.Delete(product)
	basecon.ReturnSucceed(ctx)
}
