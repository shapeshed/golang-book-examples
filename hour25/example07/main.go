package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

var store *Store

type Server struct {
	r *httprouter.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	s.r.ServeHTTP(w, r)
}

func ListTasks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tasks, err := store.GetTasks()
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(tasks)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func CreateTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := Task{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		log.Fatal(err)
	}
	err = store.CreateTask(&t)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusCreated)
}

func ReadTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Fatal(err)
	}

	task, err := store.GetTask(id)
	if err != nil {
		if err == errKeyNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("{\"message\": \"not found\"}"))
			return
		} else {
			log.Fatal(err)
		}
	}
	b, err := json.Marshal(task)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(b)
}

func UpdateTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Fatal(err)
	}

	t := Task{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&t)
	if err != nil {
		log.Fatal(err)
	}

	t.Id = id

	err = store.UpdateTask(&t)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Fatal(err)
	}

	err = store.DeleteTask(id)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	router := httprouter.New()
	router.GET("/", ListTasks)
	router.POST("/", CreateTask)
	router.GET("/:id", ReadTask)
	router.PUT("/:id", UpdateTask)
	router.DELETE("/:id", DeleteTask)

	var err error
	store, err = NewStore()
	if err != nil {
		log.Fatal(err)
	}
	store.Initialize()
	defer store.Close()

	log.Fatal(http.ListenAndServe(":8080", &Server{router}))
}
