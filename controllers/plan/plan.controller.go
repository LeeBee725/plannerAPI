package plan

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"leebee.io/planner/rest/models/plan"
)

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

func getAllPlansHandler(ctx *gin.Context) {
	log.Printf("plan: %+v\n", Plans)
	ctx.JSON(http.StatusOK, Plans)
}

func getPlanByIdHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		panic(err)
	}
	log.Printf("plan: %+v\n", Plans[uint32(id)])
	ctx.JSON(http.StatusOK, Plans[uint32(id)])
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

	Plans[uint32(id)].Update(plan)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   Plans[uint32(id)],
	})
}
