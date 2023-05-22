package main

import (
	"fmt"
	"log"

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

	p := plan.NewPlan("2023-05-17T14:00:00Z", "2023-05-17T15:00:00Z", "Test 짜보기")
	p1 := plan.NewPlan("2023-05-18T15:21:00Z", "2023-05-18T19:32:00Z", "haha")
	fmt.Println(p, p1)
	// err := http.ListenAndServe(":8080", nil)
}
