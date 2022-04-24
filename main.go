package main

import (
	"github.com/gin-gonic/gin"
	"github.com/howtri/goRate/handlers"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.NoRoute(func(c *gin.Context) {
		c.File("./ui/index.html")
	})

	r.POST("/skill/add", handlers.AddSkillHandler)
	r.POST("/skill/search", handlers.SearchSkillsHandler)
	r.POST("/skill/rank", handlers.RankSkillHandler)
	r.GET("/skill/:id", handlers.GetSkillHandler)

	err := r.Run(":3002")
	if err != nil {
		panic(err)
	}

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
