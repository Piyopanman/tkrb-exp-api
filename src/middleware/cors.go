package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//Cors CORSの設定
func Cors() gin.HandlerFunc {
	return (cors.New(cors.Config{
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Content-Type",
			"Accept",
			"Origin",
		},
		AllowOrigins: []string{
			"*",
		},
		MaxAge: 24 * time.Hour,
	}))
}
