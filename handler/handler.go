package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST opening",
	})
}
func ShowOpeningHandle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
			"msg": "GET opening",
	})
}
func DeleteOpeningHandle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET opening",
	})
}
func UpdateOpeningHandle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET opening",
	})
}
func ListOpeningsHandle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET opening",
	})
}