package main

import (
	database "contacts/databases"
	"contacts/models"
	"fmt"
	"log"
	"time"
)

var (
	dsn string
)

func main() {

	dsn := "host=localhost user=postgres password=postgres dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := database.GetConnection(dsn)
	if err != nil {
		log.Fatal(err)
	}

	contactdb := database.Contact{DB: db}
	contact := new(models.Contact)
	contact.Name = "Jiten"
	contact.Email = "Jitenp@outlook.Com"
	contact.Status = "Active"
	contact.Lastmodified = int64(time.Now().Unix())

	err = contactdb.Create(contact)
	if err != nil {
		log.Println(err)
	}

	c1, err := contactdb.Get("1")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(c1)

}
