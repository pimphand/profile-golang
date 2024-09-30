package product_controller

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/responses"
	"gin-gonic-gorm/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAllProducts(ctx *gin.Context) {
	var products []responses.Product

	perPage := ctx.DefaultQuery("perpage", "20")
	perPageInt, _ := strconv.Atoi(perPage)

	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)

	search := ctx.DefaultQuery("search", "")

	var total int64
	database.DB.Table("products").Where("deleted_at IS NULL").Count(&total)

	err := database.DB.Table("products").
		Preload("Category").
		Preload("Skus.Images").
		Offset((pageInt - 1) * perPageInt).
		Limit(perPageInt).
		Scopes(responses.Search(search)).
		Order("id desc").
		Find(&products).Error

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "Failed to get products",
		})
		return
	}

	meta := utils.PaginateMeta(ctx, int(total), perPageInt, pageInt)

	// Mengembalikan response
	ctx.JSON(200, gin.H{
		"data": products,
		"links": gin.H{
			"first": meta["path"].(string) + "?page=1",
			"last":  meta["path"].(string) + "?page=" + strconv.Itoa(meta["last_page"].(int)),
			"prev":  meta["path"].(string) + "?page=" + strconv.Itoa(pageInt-1),
			"next":  meta["path"].(string) + "?page=" + strconv.Itoa(pageInt+1),
		},
		"meta": meta,
	})
}
