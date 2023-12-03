package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/seew0/loginAutomation/middlewares"
	"github.com/seew0/loginAutomation/router"
)

type Server struct {
	port   string
	engine *gin.Engine
	router router.Router
}

func NewServer(port string, engine *gin.Engine) *Server {
	return &Server{
		port:   port,
		engine: engine,
	}
}

func (serve *Server) Start() {
	serve.engine.Use(middlewares.CORSmanager)

	serve.engine.GET("/health", func(ctx *gin.Context) {
		serve.router.Health(ctx)
	})

	serve.engine.POST("/getLoginSmartlead", func(c *gin.Context) {
		serve.router.GetLoginSmartlead(c)
	})

	serve.engine.POST("/updateFiveSimOptions", func(ctx *gin.Context) {
		serve.router.UpdateFiveSimOptions(ctx)
	})
	
	fmt.Printf("Starting Server at port %v \n", serve.port)
	
	err := serve.engine.Run(serve.port)
	if err != nil {
		fmt.Printf("Failed to Start Server")
	}
}
