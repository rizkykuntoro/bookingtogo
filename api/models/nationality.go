package models

import (
	"encoding/json"
	"log"

	_ "github.com/lib/pq"
)

type Nationality struct {
	Nationality_id   int    `json:"nationality_id"`
	Nationality_name string `json:"nationality_name"`
	Nationality_code string `json:"nationality_code"`
}

func NationalityGet(nationality_codes string) []byte {
	db := Init()

	defer db.Close()

	var q string

	if nationality_codes != "" {
		q = `SELECT * FROM "nationality" WHERE nationality_code = '` + nationality_codes + `'`
	} else {
		q = `SELECT * FROM "nationality"`
	}

	rows, err := db.Query(q)

	if err != nil {
		log.Println(err.Error())
	}

	defer rows.Close()

	var nationality []Nationality

	for rows.Next() {
		var nationality_name string
		var nationality_code string
		var nationality_id int

		rows.Scan(&nationality_id, &nationality_name, &nationality_code)
		nationality = append(nationality, Nationality{nationality_id, nationality_name, nationality_code})
	}

	nationalityBytes, _ := json.Marshal(&nationality)

	return nationalityBytes
}

func NationalityInsert(nationality_name, nationality_code string) string {
	db := Init()

	defer db.Close()

	q := `INSERT INTO "nationality"("nationality_name", "nationality_code") VALUES($1, $2)`
	_, err := db.Exec(q, nationality_name, nationality_code)

	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}

	return "insert sukses"
}

func NationalityUpdate(nationality_name, nationality_code string, nationality_id int) string {
	db := Init()

	defer db.Close()

	q := `UPDATE "nationality" SET "nationality_name" = $1, "nationality_code" = $2 WHERE "nationality_id" = $3`
	_, err := db.Exec(q, nationality_name, nationality_code, nationality_id)

	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}

	return "update sukses"
}

func NationalityDelete(nationality_id int) string {
	db := Init()

	defer db.Close()

	q := `DELETE FROM "nationality" WHERE "nationality_id" = $1`
	_, err := db.Exec(q, nationality_id)

	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}

	return "delete sukses"
}
