package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"leebee.io/planner/rest/plan"
)

var plans = map[uint32]*plan.Plan{}

func getAllPlansHandler(ctx *gin.Context) {
	log.Printf("plan: %+v\n", plans)
	ctx.JSON(http.StatusOK, plans)
}

func getPlanByIdHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		panic(err)
	}
	log.Printf("plan: %+v\n", plans[uint32(id)])
	ctx.JSON(http.StatusOK, plans[uint32(id)])
}

func createPlanHandler(ctx *gin.Context) {
	plan := plan.NewPlan()
	if err := ctx.ShouldBindJSON(plan); err != nil {
		log.Println("err: ", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	log.Printf("plan: %+v\n", plan)

	plans[plan.Id()] = plan

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   plan,
	})
}

func updatePlanHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		panic(err)
	}

	var plan plan.Plan
	if err := ctx.ShouldBindJSON(&plan); err != nil {
		log.Println("err: ", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	log.Printf("plan: %+v\n", plan)

	plans[uint32(id)].Update(plan)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   plans[uint32(id)],
	})
}

func main() {
	log.SetPrefix("\033[32mLOG::\033[0m")
	defer func() {
		r := recover()
		if r != nil {
			log.Panic(r)
		}
	}()

	r := gin.Default()
	plans := r.Group("/plans")
	plans.GET("", getAllPlansHandler)
	plans.GET("/:id", getPlanByIdHandler)
	plans.POST("", createPlanHandler)
	plans.PUT("/:id", updatePlanHandler)

	r.Run()
}
