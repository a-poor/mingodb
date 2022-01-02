package main

import (
	"log"
	"os"
)

func main() {
	dbPath := "./test.db"

	db, err := Open(dbPath)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// Create a new collection.
	c, err := db.Collection("users")
	if err != nil {
		log.Panic(err)
	}

	// Insert a new document.
	id, err := c.InsertOne(map[string]interface{}{
		"name": "John Doe",
		"age":  30,
	})
	if err != nil {
		log.Panic(err)
	}
	log.Println("Inserted document with id:", id)

	// Retrieve the inserted document.
	doc, err := c.GetByID(id)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Retrieved document:", doc)

	os.Remove(dbPath)
}
