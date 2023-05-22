package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

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
		log.Print(r.Method, " ", r.URL)
		switch r.Method {
		case http.MethodGet:
			s := strings.Split(r.URL.Path, "/")
			if len(s) > 4 {
				http.Error(w, "Not found", 404)
				break
			}
			if s[2] == "" {
				json.NewEncoder(w).Encode(plans)
				break
			}
			id, err := strconv.ParseUint(s[2], 10, 32)
			if err != nil {
				panic(err)
			}
			json.NewEncoder(w).Encode(plans[uint32(id)])
		case http.MethodPost:
			p := plan.NewPlan()
			json.NewDecoder(r.Body).Decode(p)
			plans[p.Id()] = p
			json.NewEncoder(w).Encode(p)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
