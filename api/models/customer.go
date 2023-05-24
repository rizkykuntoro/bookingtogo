package models

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type Customer struct {
	Cst_id         int       `json:"cst_id"`
	Nationality_id int       `json:"nationality_id"`
	Cst_name       string    `json:"cst_name"`
	Cst_dob        time.Time `json:"cst_dob"`
	Cst_phone      string    `json:"cst_phone"`
	Cst_email      string    `json:"cst_email"`
}

func CustomerGet(cst_ids string) []byte {
	db := Init()

	defer db.Close()

	var q string

	if cst_ids != "" {
		q = `SELECT * FROM "customer" WHERE cst_id = '` + cst_ids + `'`
	} else {
		q = `SELECT * FROM "customer"`
	}

	rows, err := db.Query(q)

	if err != nil {
		log.Println(err.Error())
	}

	defer rows.Close()

	var customer []Customer

	for rows.Next() {
		var cst_id int
		var nationality_id int
		var cst_name string
		var cst_dob time.Time
		var cst_phone string
		var cst_email string

		rows.Scan(&cst_id, &nationality_id, &cst_name, &cst_dob, &cst_phone, &cst_email)
		customer = append(customer, Customer{cst_id, nationality_id, cst_name, cst_dob, cst_phone, cst_email})
	}

	customerBytes, _ := json.Marshal(&customer)

	return customerBytes
}

func CustomerInsert(nationality_id int, cst_name, cst_dob, cst_phone, cst_email string) string {
	db := Init()

	defer db.Close()
	lastInsertId := 0
	q := `INSERT INTO "customer"("nationality_id", "cst_name", "cst_dob", "cst_phone", "cst_email") VALUES($1, $2, $3, $4, $5) RETURNING cst_id`
	err := db.QueryRow(q, nationality_id, cst_name, cst_dob, cst_phone, cst_email).Scan(&lastInsertId)

	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}

	return strconv.Itoa(lastInsertId)
}

func CustomerUpdate(cst_id, nationality_id int, cst_name, cst_dob, cst_phone, cst_email string) string {
	db := Init()

	defer db.Close()

	q := `UPDATE "customer" SET "nationality_id" = $1, "cst_name" = $2, "cst_dob" = $3, "cst_phone" = $4, "cst_email" = $5 WHERE "cst_id" = $6`
	_, err := db.Exec(q, nationality_id, cst_name, cst_dob, cst_phone, cst_email, cst_id)

	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}

	return "update sukses"
}

func CustomerDelete(cst_id int) string {
	db := Init()

	defer db.Close()

	q := `DELETE FROM "customer" WHERE "cst_id" = $1`
	_, err := db.Exec(q, cst_id)

	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}

	return "delete sukses"
}
