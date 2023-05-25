package plan

import (
	"github.com/gin-gonic/gin"
)

type PlanController interface {
	CreatePlan(ctx *gin.Context)
	GetAllPlans(ctx *gin.Context)
	GetPlanById(ctx *gin.Context)
	UpdatePlan(ctx *gin.Context)
}

// type PlanControllerImpl struct {
// 	psrv srv.PlanService
// }

// func (pc *PlanControllerImpl) CreatePlan(ctx *gin.Context) {
// 	p := new(model.Plan)
// }
// func (pc *PlanControllerImpl) GetAllPlans(ctx *gin.Context) {

// }
// func (pc *PlanControllerImpl) GetPlanById(ctx *gin.Context) {

// }
// func (pc *PlanControllerImpl) UpdatePlan(ctx *gin.Context) {

// }

// func createPlanHandler(ctx *gin.Context) {
// 	plan := plan.NewPlan()
// 	if err := ctx.ShouldBindJSON(plan); err != nil {
// 		log.Println("err: ", err)
// 		ctx.AbortWithStatus(http.StatusBadRequest)
// 	}
// 	log.Printf("plan: %+v\n", plan)

// 	plans[plan.Id()] = plan

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"status": "ok",
// 		"data":   plan,
// 	})
// }

// func getAllPlansHandler(ctx *gin.Context) {
// 	log.Printf("plan: %+v\n", Plans)
// 	ctx.JSON(http.StatusOK, Plans)
// }

// func getPlanByIdHandler(ctx *gin.Context) {
// 	idStr := ctx.Param("id")
// 	id, err := strconv.ParseUint(idStr, 10, 32)
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Printf("plan: %+v\n", Plans[uint32(id)])
// 	ctx.JSON(http.StatusOK, Plans[uint32(id)])
// }

// func updatePlanHandler(ctx *gin.Context) {
// 	idStr := ctx.Param("id")
// 	id, err := strconv.ParseUint(idStr, 10, 32)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var plan plan.Plan
// 	if err := ctx.ShouldBindJSON(&plan); err != nil {
// 		log.Println("err: ", err)
// 		ctx.AbortWithStatus(http.StatusBadRequest)
// 	}
// 	log.Printf("plan: %+v\n", plan)

// 	Plans[uint32(id)].Update(plan)

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"status": "ok",
// 		"data":   Plans[uint32(id)],
// 	})
// }
