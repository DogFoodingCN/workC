package routers

import (
	"github.com/DogFoodingCN/workC/pkg/setting"
	job2 "github.com/DogFoodingCN/workC/service/job"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.Use(Cors())
	
	gin.SetMode(setting.RunMode)

	job := r.Group("/job")
	{
		job.POST("/save", job2.SaveJobs)
	}

	return r
}
