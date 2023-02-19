package routers

import (
	"github.com/EDDYCJY/go-gin-example/routers/frontend"
	"github.com/gin-gonic/gin"
)

func InitFrontendRouter(r *gin.Engine) *gin.Engine {
	r.POST("/login", frontend.Login)
	r.POST("/register", frontend.Register)
	return r
}
