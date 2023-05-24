package controllers

import (
	"crud_mux/models"
	"encoding/json"
	"log"
	"net/http"
)

type Customer struct {
	Cst_id         int
	Nationality_id int
	Cst_name       string
	Cst_dob        string
	Cst_phone      string
	Cst_email      string
}

func CustomerIndex(w http.ResponseWriter, r *http.Request) {
	cst_id := r.URL.Query().Get("cst_id")

	resp := models.CustomerGet(cst_id)
	w.Write(resp)
}

func CustomerAdd(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var n Customer
	err := decoder.Decode(&n)
	if err != nil {
		log.Println(err.Error())
	}

	resp := models.CustomerInsert(n.Nationality_id, n.Cst_name, n.Cst_dob, n.Cst_phone, n.Cst_email)
	w.Write([]byte(resp))
}

func CustomerUpdate(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var n Customer
	err := decoder.Decode(&n)
	if err != nil {
		panic(err)
	}

	resp := models.CustomerUpdate(n.Cst_id, n.Nationality_id, n.Cst_name, n.Cst_dob, n.Cst_phone, n.Cst_email)
	w.Write([]byte(resp + "\n"))
}

func CustomerDelete(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var n Customer
	err := decoder.Decode(&n)
	if err != nil {
		panic(err)
	}

	resp := models.CustomerDelete(n.Cst_id)
	w.Write([]byte(resp + "\n"))
}
