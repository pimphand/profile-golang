package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
)

func CreatePageLinks(ctx *gin.Context, currentPage, lastPage int) []gin.H {
	var links []gin.H
	path := "https://api-go.naisha.id" + ctx.Request.URL.Path // Mengambil path dari request

	// Link ke halaman sebelumnya
	if currentPage > 1 {
		links = append(links, gin.H{
			"url":    fmt.Sprintf("%s?page=%d", path, currentPage-1),
			"label":  "&laquo; Sebelumnya",
			"active": false,
		})
	} else {
		links = append(links, gin.H{
			"url":    nil,
			"label":  "&laquo; Sebelumnya",
			"active": false,
		})
	}

	// Link ke setiap halaman
	for i := 1; i <= lastPage; i++ {
		links = append(links, gin.H{
			"url":    fmt.Sprintf("%s?page=%d", path, i),
			"label":  fmt.Sprintf("%d", i),
			"active": currentPage == i,
		})
	}

	// Link ke halaman berikutnya
	if currentPage < lastPage {
		links = append(links, gin.H{
			"url":    fmt.Sprintf("%s?page=%d", path, currentPage+1),
			"label":  "Berikutnya &raquo;",
			"active": false,
		})
	} else {
		links = append(links, gin.H{
			"url":    nil,
			"label":  "Berikutnya &raquo;",
			"active": false,
		})
	}

	return links
}

func PaginateMeta(ctx *gin.Context, total, perPage, currentPage int) gin.H {

	lastPage := int(math.Ceil(float64(total) / float64(perPage)))

	// Meta pagination
	meta := gin.H{
		"current_page": currentPage,
		"from":         (currentPage-1)*perPage + 1,
		"last_page":    lastPage,
		"links":        CreatePageLinks(ctx, currentPage, lastPage),
		"path":         ctx.Request.URL.Path,
		"per_page":     perPage,
		"to":           currentPage * perPage,
		"total":        total,
	}

	if currentPage*perPage > total {
		meta["to"] = total
	}

	return meta
}
