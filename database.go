package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"fmt"
)

func SetupDB() *MongoDB {
	db := MongoDB{
		os.Getenv("MONGODB_URI"),
		"heroku_pgvgprmm",
		"currencyCollection",
	}

	fmt.Println(db.DatabaseURL)


	session, err := mgo.Dial(db.DatabaseURL)
	defer session.Close()

	if err != nil{
		log.Fatal(err.Error())
	}

	return &db
}

// Init initialising the db
func (db *MongoDB) Init() {

	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	index := mgo.Index{
		Key:        []string{"currencyid"},
		Unique:     true,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	}

	err = session.DB(db.DatabaseName).C(db.ColCurrency).EnsureIndex(index)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Add adds the db
func (db *MongoDB) Add(p Currency) error {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer session.Close()

	err = session.DB(db.DatabaseName).C(db.ColCurrency).Insert(p)

	if err != nil {
		log.Printf("Could not add to db, error in Insert(): %v", err.Error())
		return err
	}

	return nil
}

// Count counts the colCurrency
func (db *MongoDB) Count() int {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	count, err := session.DB(db.DatabaseName).C(db.ColCurrency).Count()
	if err != nil {
		log.Printf("Error in Count(): %v", err.Error())
		return -1
	}

	return count
}