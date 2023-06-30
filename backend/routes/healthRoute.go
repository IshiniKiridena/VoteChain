package routes

import (
	"github.com/IshiniKiridena/VoteChainBackend/middleware"
    "github.com/gin-gonic/gin"
)

func HealthRouter(router *gin.Engine)  {
	router.GET("/health", middleware.Health)
}