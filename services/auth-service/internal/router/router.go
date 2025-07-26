// package router

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/handler"
// )

// func SetupRouter(authHandler *handler.AuthHandler) *gin.Engine {
//     r := gin.Default()

//     v1 := r.Group("/v1")
//     {
//         v1.POST("/register", authHandler.Register)
//         v1.POST("/login", authHandler.Login)
//     }

//     return r
// }

package router

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/handler"
)

func SetupRouter(authHandler *handler.AuthHandler, logger *zap.Logger) *gin.Engine {
	r := gin.New() // use gin.New() to avoid default logger/middleware

	// Add zap logger middleware (logs each HTTP request)
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// Add recovery middleware to catch panics and log them with zap
	r.Use(ginzap.RecoveryWithZap(logger, true))

	v1 := r.Group("/v1")
	{
		v1.POST("/register", authHandler.Register)
		v1.POST("/login", authHandler.Login)
	}

	return r
}
