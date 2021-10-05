package controller

import (
	"github.com/gin-gonic/gin"
	"joke/module"
	"net/http"
	"time"
)

// Add one remark
func Add(content *gin.Context) {
	nowTime := time.Now()
	cont := content.PostForm("content")
	remark := &module.Remark{Content: cont, CreateTime: &nowTime, UpdateTime: &nowTime}
	module.AddOne(remark)
}

// GetRandom get one remark randomly
func GetRandom(content *gin.Context) {
	remark := module.GetOne()
	content.HTML(http.StatusOK, "index.html", gin.H{"content": remark.Content})
}

func AddHtml(context *gin.Context) {
	context.HTML(http.StatusOK, "add.html", nil)
}
