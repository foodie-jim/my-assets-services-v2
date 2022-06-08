package main

import (
	"net/http"
	"os"
	
	"github.com/gin-gonic/gin"
	docs "foodie-jim/my-assets-services-v2/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/files"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	r := gin.Default()
	r.Use(gin.Logger())

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}
	}	

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":" + port)
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context)  {
	g.JSON(http.StatusOK,"helloworld")
}