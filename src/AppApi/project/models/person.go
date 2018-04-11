package models

import (
	db "AppApi/project/database"
	"log"
)

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (p Person) DeletePerson() (id int64, err error) {
	res, err := db.SqlDB.Exec("delete from person where id =?", p.Id)
	if err != nil {
		log.Fatalln(err)
		return
	}

	id, err = res.RowsAffected()
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}

func (p Person) AddPerson() (id int64, err error) {
	stmt, err := db.SqlDB.Prepare("insert person (first_name, last_name) values (?, ?)")
	res, err := stmt.Exec(p.FirstName, p.LastName)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	if err != nil {
		return
	}

	return
}

func getAllPerson() (persons []Person, err error) {
	persons = make([]Person, 0)
	rows, err := db.SqlDB.Query("select id, first_name, last_name from person")
	checkErr(err)

	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}

func SearchPerson(index int) (person Person, err error) {
	err = db.SqlDB.QueryRow("select id, first_name, last_name from person where id = ?", index).Scan(
		&person.Id, &person.FirstName, &person.LastName,
	)
	if err != nil {
		return
	}

	return

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
