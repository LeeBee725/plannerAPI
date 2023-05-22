package main

import (
	"encoding/json"
	"log"
	"net/http"

	"leebee.io/planner/rest/plan"
)

var plans = map[uint32]*plan.Plan{}

func main() {
	log.SetPrefix("\033[32mLOG::\033[0m")
	defer func() {
		r := recover()
		if r != nil {
			log.Panic(r)
		}
	}()

	http.HandleFunc("/plans/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(plans)
		case http.MethodPost:
			var p plan.Plan
			json.NewDecoder(r.Body).Decode(&p)
			plans[p.Id()] = &p
			json.NewEncoder(w).Encode(&p)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
