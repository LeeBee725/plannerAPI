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

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
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

	mux := http.NewServeMux()

	planHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.Method, " ", r.URL)
		switch r.Method {
		case http.MethodGet:
			s := strings.Split(r.URL.Path, "/")
			if len(s) != 3 {
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
		case http.MethodPut:
			s := strings.Split(r.URL.Path, "/")
			if len(s) != 3 || s[2] == "" {
				http.Error(w, "Not found", 404)
				break
			}
			id, err := strconv.ParseUint(s[2], 10, 32)
			if err != nil {
				panic(err)
			}
			p := plans[uint32(id)]
			var newP plan.Plan
			err = json.NewDecoder(r.Body).Decode(&newP)
			if err != nil {
				panic(err)
			}
			log.Print(newP)
			p.Update(newP)
			log.Print(p)
			json.NewEncoder(w).Encode(p)
		}
	})

	mux.Handle("/plans/", jsonContentTypeMiddleware(planHandler))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
