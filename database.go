package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
)

// SetupDB sets up the database
func SetupDB() *MongoDB {
	db := MongoDB{
		os.Getenv("MONGODB"), // Environment variable from Heroku
		"heroku_pgvgprmm",
		"currencyCollection",
	}

	fmt.Println(db.DatabaseURL)

	session, err := mgo.Dial(db.DatabaseURL)
	defer session.Close()

	if err != nil {
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
func (db *MongoDB) Add(data Data2d) error {

	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer session.Close()

	err = session.DB(db.DatabaseName).C(db.ColCurrency).Insert(data)

	if err != nil {
		log.Printf("Could not add to db, error in Insert(): %v", err.Error())
		return err
	}

	return nil
}

// GetLatest gets the latest currencies with date as index
func (db *MongoDB) GetLatest(date string) (Data2d, bool) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	data2d := Data2d{}
	notToday := true

	err = session.DB(db.DatabaseName).C(db.ColCurrency).Find(bson.M{"date": date}).One(&data2d)
	if err != nil {
		notToday = false
	}

	return data2d, notToday
}

// DailyCurrencyAdder adds currency once a day
func DailyCurrencyAdder() {
	data2d := GetCurrency()
	db := SetupDB()
	db.Init()
	db.Add(data2d)
}
