package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHander(content *gin.Context) {
	content.Next()

	if len(content.Errors) > 0 {
		content.HTML(http.StatusBadRequest, "400", gin.H{
			"errors": content.Errors,
		})
	}
}
