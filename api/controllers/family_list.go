package controllers

import (
	"crud_mux/models"
	"encoding/json"
	"log"
	"net/http"
)

type Family_list struct {
	Fl_id       int
	Cst_id      int
	Fl_relation string
	Fl_name     string
	Fl_dob      string
}

func FamilyIndex(w http.ResponseWriter, r *http.Request) {
	fl_id := r.URL.Query().Get("fl_id")

	resp := models.FamilyGet(fl_id)
	w.Write(resp)
}

func FamilyAdd(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var n Family_list
	err := decoder.Decode(&n)
	if err != nil {
		log.Println(err.Error())
	}

	resp := models.FamilyInsert(n.Cst_id, n.Fl_relation, n.Fl_name, n.Fl_dob)
	w.Write([]byte(resp + "\n"))
}

func FamilyUpdate(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var n Family_list
	err := decoder.Decode(&n)
	if err != nil {
		log.Println(err.Error())
	}

	resp := models.FamilyUpdate(n.Cst_id, n.Fl_relation, n.Fl_name, n.Fl_dob, n.Fl_id)
	w.Write([]byte(resp + "\n"))
}

func FamilyDelete(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var n Family_list
	err := decoder.Decode(&n)
	if err != nil {
		log.Println(err.Error())
	}

	resp := models.FamilyDelete(n.Fl_id)
	w.Write([]byte(resp + "\n"))
}
