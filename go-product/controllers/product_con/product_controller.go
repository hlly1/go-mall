package productcon

import (
	"log"

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

	log.Println("[Get Product]: Request Params - id: ", ctx.Query("id"))
	rows, err := dao.DB.Model(&entities.Product{}).Raw("select * from product where id = ?", ctx.Query("id")).Rows()
	if err != nil {
		panic(err)
	}
	log.Println("[Get Product]: debug: ", rows)
	var product entities.Product
	for rows.Next() {
		dao.DB.ScanRows(rows, &product)
	}
	if !(len(product.ID) > 0) {
		basecon.ReturnCustom(ctx, "0001", "Not Found")
		ctx.Next()
	} else {
		log.Println("[Get Product]: Found Target - ", product)
		basecon.ReturnSucceedWithArgus(ctx, basecon.NewJson("product", product))
	}
}

func (con ProductController) Add(ctx *gin.Context) {
	product := entities.Product{}
	ctx.ShouldBindJSON(&product)
	prod_id := strings.Replace(uuid.NewString(), "-", "", -1)
	product.ID = prod_id
	dao.DB.Create(product)
	basecon.ReturnSucceedWithArgus(ctx, basecon.NewJson("id", prod_id))
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
