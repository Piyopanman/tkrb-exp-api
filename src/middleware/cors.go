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
			"GET",
			"POST",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Content-Type",
			"Accept",
			"Origin",
			"Access-Control-Allow-Origin",
		},
		AllowOrigins: []string{
			"http://localhost:3030",
		},
		MaxAge: 24 * time.Hour,
	}))
}
