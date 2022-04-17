package main

import (
	"github.com/gin-gonic/gin"
	"github.com/howtri/goRate/handlers"
)

func main() {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.File("./ui/index.html")
	})

	r.POST("/skill/add", handlers.AddSkillHandler)
	//r.POST("/skill/rank/:id", handlers.AddTodoHandler)
	r.POST("/skill/rank", handlers.RankSkillsHandler)
	// r.GET("/skill/:id", handlers.DeleteTodoHandler)
	r.GET("/skill", handlers.GetAllSkillsHandler)

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
