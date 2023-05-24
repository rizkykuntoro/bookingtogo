package controllers

import (
	"crud_mux/models"
	"encoding/json"
	"log"
	"net/http"
)

type Nationality struct {
	Nationality_name string
	Nationality_code string
	Nationality_id   int
}

func NationalityIndex(w http.ResponseWriter, r *http.Request) {
	nationality_code := r.URL.Query().Get("nationality_code")

	resp := models.NationalityGet(nationality_code)
	w.Write(resp)
}

func NationalityAdd(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var n Nationality
	err := decoder.Decode(&n)
	if err != nil {
		log.Println(err.Error())
	}

	resp := models.NationalityInsert(n.Nationality_name, n.Nationality_code)
	w.Write([]byte(resp + "\n"))
}

func NationalityUpdate(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var n Nationality
	err := decoder.Decode(&n)
	if err != nil {
		panic(err)
	}

	resp := models.NationalityUpdate(n.Nationality_name, n.Nationality_code, n.Nationality_id)
	w.Write([]byte(resp + "\n"))
}

func NationalityDelete(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var n Nationality
	err := decoder.Decode(&n)
	if err != nil {
		panic(err)
	}

	resp := models.NationalityDelete(n.Nationality_id)
	w.Write([]byte(resp + "\n"))
}
