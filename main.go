package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"leebee.io/planner/rest/models/plan"
)

var Plans = map[uint32]*plan.Plan{}

func plannerAPI() *gin.Engine {
	r := gin.Default()
	plans := r.Group("/plans")
	plans.GET("", getAllPlansHandler)
	plans.GET("/:id", getPlanByIdHandler)
	plans.POST("", createPlanHandler)
	plans.PUT("/:id", updatePlanHandler)
	return r
}

func main() {
	log.SetPrefix("\033[32mLOG::\033[0m")
	defer func() {
		r := recover()
		if r != nil {
			log.Panic(r)
		}
	}()

	plannerAPI().Run()
}
