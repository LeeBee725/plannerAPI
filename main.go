package main

import (
	"fmt"
	"log"
	"time"

	"leebee.io/planner/rest/plan"
)

func main() {
	log.SetPrefix("\033[32mLOG::\033[0m")
	defer func() {
		r := recover()
		if r != nil {
			log.Panic(r)
		}
	}()
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		panic(err)
	}
	if loc == nil {
		fmt.Println("haha")
	}
	p := plan.NewPlan()
	p.SetStartTime(2023, 5, 17, 14, 0, 0, 0, loc)
	p.SetEndTime(2023, 5, 17, 15, 0, 0, 0, loc)
	p.SetContent("Test 짜보기")
	p1 := plan.NewPlan()
	p1.SetStartTime(2023, 5, 17, 14, 0, 0, 0, loc)
	p1.SetEndTime(2023, 5, 17, 15, 0, 0, 0, loc)
	p1.SetContent("Test 짜보기")
	fmt.Println(p, p1)
	// err := http.ListenAndServe(":8080", nil)
}
