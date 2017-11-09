package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Server struct {
	r *httprouter.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	s.r.ServeHTTP(w, r)
}

func ListTasks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "ListTasks\n")
}

func CreateTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "CreateTasks\n")
}

func ReadTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "ReadTask\n")
}

func UpdateTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "UpdateTask\n")
}

func DeleteTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "DeleteTask\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", ListTasks)
	router.POST("/", CreateTask)
	router.GET("/:id", ReadTask)
	router.PUT("/:id", UpdateTask)
	router.DELETE("/:id", DeleteTask)

	log.Fatal(http.ListenAndServe(":8080", &Server{router}))

}
