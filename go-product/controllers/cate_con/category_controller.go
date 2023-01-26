package catecon

import (
	basecon "go-product/controllers/base_con"
	"go-product/dao"
	"go-product/entities"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	basecon.BaseController
}

func (con CategoryController) GetCateTree(ctx *gin.Context) {
	var cateList []entities.Category
	cateRows, err := dao.DB.Model(&entities.Category{}).Raw("select * from category").Rows()
	if err != nil {
		panic(err)
	}
	for cateRows.Next() {
		dao.DB.ScanRows(cateRows, &cateList)
	}
	cateList = tree(cateList, 0)
	basecon.ReturnSucceedWithArgus(ctx, basecon.NewJson("category", cateList))
}

func tree(list []entities.Category, pid int16) []entities.Category {
	var treeList []entities.Category
	for _, r := range list {
		if r.ParentCid == pid {
			r.Children = tree(list, r.CatId)
			treeList = append(treeList, r)
		}
	}
	return treeList
}
