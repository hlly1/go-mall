package router

import (
	"go-product/controllers/product"
	"go-product/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRouter(r *gin.Engine) {

	productRouter := r.Group("/product", middlewares.IsLogin)
	{
		productRouter.GET("/create", product.ProductController{}.Add)

		productRouter.GET("/get", product.ProductController{}.Index)

		productRouter.GET("/update", product.ProductController{}.Update)

		productRouter.GET("/delete", product.ProductController{}.Delete)

	}
}
