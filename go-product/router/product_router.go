package router

import (
	prod_con "go-product/controllers/product_con"

	cate_con "go-product/controllers/cate_con"

	"github.com/gin-gonic/gin"
)

func ProductRouter(r *gin.Engine) {

	productRouter := r.Group("/product")
	{
		productRouter.POST("/add", prod_con.ProductController{}.Add)

		productRouter.GET("/get", prod_con.ProductController{}.Index)

		productRouter.POST("/update", prod_con.ProductController{}.Update)

		productRouter.DELETE("/delete", prod_con.ProductController{}.Delete)

	}

	categoryRouter := r.Group("/category")
	{
		categoryRouter.GET("/get", cate_con.CategoryController{}.GetCateTree)
	}
}
