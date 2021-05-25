package person

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lhendges/simple-go-crud/config"
)

var db *sql.DB

type Person struct {
	Id   int64
	Name string
	Age  int64
}

func init() {
	config.Connect()
	db = config.GetDb()
}

func (p *Person) CreatePerson() (int64, error) {
	insertedId := 0
	query := fmt.Sprintf("INSERT INTO PERSONS (name, age) VALUES ('%s', %d) RETURNING Id", p.Name, p.Age)
	err := db.QueryRow(query).Scan(&insertedId)

	if err != nil {
		fmt.Println("Erro ao inserir o registro: " + err.Error())
		return -1, err
	}
	return int64(insertedId), nil
}

func GetAllPersons() ([]Person, error) {
	rows, err := db.Query("SELECT * FROM PERSONS")

	if err != nil {
		fmt.Println("Erro ao consultar os registros: " + err.Error())
		return nil, err
	}

	persons := []Person{}
	for rows.Next() {
		var r Person

		err := rows.Scan(&r.Id, &r.Name, &r.Age)
		if err != nil {
			return nil, err
		}

		persons = append(persons, r)
	}
	return persons, nil
}

func GetPersonById(id int64) (Person, error) {
	query := fmt.Sprintf("SELECT * FROM PERSONS WHERE ID = %d", id)
	row := db.QueryRow(query)
	person := Person{}

	err := row.Scan(&person.Id, &person.Name, &person.Age)

	if err != nil && err == sql.ErrNoRows {
		return Person{}, errors.New("person not found")
	}

	return person, nil
}

func (p *Person) UpdatePerson(newName string, newAge int64) (Person, error) {
	query := fmt.Sprintf("UPDATE PERSONS SET name = '%s', age = %d WHERE id = %d;", newName, newAge, p.Id)
	_, err := db.Exec(query)

	if err != nil {
		return Person{}, errors.New("record not updated")
	}

	updatedPerson, err := GetPersonById(p.Id)

	return updatedPerson, err
}

func (p *Person) DeletePerson() error {
	query := fmt.Sprintf("DELETE FROM PERSONS WHERE id = %d;", p.Id)
	_, err := db.Exec(query)

	return err
}
