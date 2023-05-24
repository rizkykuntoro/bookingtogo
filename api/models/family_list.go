package models

import (
	"encoding/json"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type FamilyList struct {
	Fl_id       int       `json:"fl_id"`
	Cst_id      int       `json:"cst_id"`
	Fl_relation string    `json:"fl_relation"`
	Fl_name     string    `json:"fl_name"`
	Fl_dob      time.Time `json:"fl_dob"`
}

func FamilyGet(fl_id string) []byte {
	db := Init()

	defer db.Close()

	var q string

	if fl_id != "" {
		q = `SELECT * FROM "family_list" WHERE fl_id = '` + fl_id + `'`
	} else {
		q = `SELECT * FROM "family_list"`
	}

	rows, err := db.Query(q)

	if err != nil {
		log.Println(err.Error())
	}

	defer rows.Close()

	var family []FamilyList

	for rows.Next() {
		var fl_id int
		var cst_id int
		var fl_relation string
		var fl_name string
		var fl_dob time.Time

		rows.Scan(&fl_id, &cst_id, &fl_relation, &fl_name, &fl_dob)
		family = append(family, FamilyList{fl_id, cst_id, fl_relation, fl_name, fl_dob})
	}

	familyBytes, _ := json.Marshal(&family)

	return familyBytes
}

func FamilyInsert(cst_id int, fl_relation, fl_name, fl_dob string) string {
	db := Init()

	defer db.Close()

	q := `INSERT INTO "family_list"("cst_id", "fl_relation", "fl_name", "fl_dob") VALUES($1, $2, $3, $4)`
	_, err := db.Exec(q, cst_id, fl_relation, fl_name, fl_dob)

	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}

	return "insert sukses"
}

func FamilyUpdate(cst_id int, fl_relation, fl_name, fl_dob string, fl_id int) string {
	db := Init()

	defer db.Close()

	q := `UPDATE "family_list" SET "cst_id" = $1, "fl_relation" = $2, "fl_name" = $3, "fl_dob" = $4 WHERE "fl_id" = $5`
	_, err := db.Exec(q, cst_id, fl_relation, fl_name, fl_dob, fl_id)

	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}

	return "update sukses"
}

func FamilyDelete(fl_id int) string {
	db := Init()

	defer db.Close()

	q := `DELETE FROM "family_list" WHERE "fl_id" = $1`
	_, err := db.Exec(q, fl_id)

	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}

	return "delete sukses"
}
