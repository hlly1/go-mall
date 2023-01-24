package router

import (
	productcon "go-product/controllers/product_con"

	"github.com/gin-gonic/gin"
)

func ProductRouter(r *gin.Engine) {

	productRouter := r.Group("/product")
	{
		productRouter.POST("/add", productcon.ProductController{}.Add)

		productRouter.GET("/get", productcon.ProductController{}.Index)

		productRouter.POST("/update", productcon.ProductController{}.Update)

		productRouter.DELETE("/delete", productcon.ProductController{}.Delete)

	}
}
