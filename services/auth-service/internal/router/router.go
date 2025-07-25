package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/handler"
)

func SetupRouter(authHandler *handler.AuthHandler) *gin.Engine {
    r := gin.Default()

    v1 := r.Group("/v1")
    {
        v1.POST("/register", authHandler.Register)
        v1.POST("/login", authHandler.Login)
    }

    return r
}
