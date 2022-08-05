package handlers

import (
	"github.com/gin-gonic/gin"
	"my_project/logger"
	_ "my_project/model/author"
	"net/http"
)

func Home(ctx *gin.Context) {
	Log := logger.GetLogger()
	Log.Info("запуск домашней страницы")
	ctx.JSON(http.StatusOK, gin.H{
		"response": gin.H{
			"method":  http.MethodGet,
			"code":    http.StatusOK,
			"message": "Добро пожаловать в мою библиотеку, перед вами список авторов и книг.",
		},
	})
}
