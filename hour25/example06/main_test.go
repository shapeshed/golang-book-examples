package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListTasksHandler(t *testing.T) {
	router := httprouter.New()
	router.GET("/", ListTasks)
	r, _ := http.NewRequest("GET", "/", nil)

	var err error
	store, err = NewStore()
	if err != nil {
		log.Fatal(err)
	}
	store.Initialize()
	defer store.Close()

	w := httptest.NewRecorder()
	s := Server{router}
	s.ServeHTTP(w, r)

	expectedContentTypeHeader := "application/json; charset=utf-8"

	if w.Header().Get("Content-Type") != expectedContentTypeHeader {
		t.Errorf("handler returned unexpected Content-Type header: got %v want %v",
			w.Header().Get("Content-Type"), expectedContentTypeHeader)
	}

	expectedCode := http.StatusOK
	if w.Code != expectedCode {
		t.Errorf("handler returned unexpected status code: got %v want %v",
			w.Code, expectedCode)
	}

}
