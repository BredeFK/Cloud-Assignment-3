package gofiles

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

// SetupDB sets up the database
func SetupDB() *MongoDB {
	db := MongoDB{
		"mongodb://user:user123@ds149855.mlab.com:49855/heroku_pgvgprmm",
		"heroku_pgvgprmm",
		"currencyCollection",
	}

	fmt.Println(db.DatabaseURL) // TODO : What's this for?

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

// Count counts the number of currency jsons
func (db *MongoDB) Count() int {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	count, err := session.DB(db.DatabaseName).C(db.ColCurrency).Count()
	if err != nil {
		log.Printf("Error in Count(): %v", err.Error())
		return -1
	}

	return count
}

// GetLatest gets the latest currencies with date as index
func (db *MongoDB) GetLatest(date string) (Data2d, bool) {

	count := db.Count()
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	data2d := Data2d{}
	notToday := true

	if date != "noDate" {
		err = session.DB(db.DatabaseName).C(db.ColCurrency).Find(bson.M{"date": date}).One(&data2d)
	} else {
		err = session.DB(db.DatabaseName).C(db.ColCurrency).Find(nil).Skip(count - 1).One(&data2d)
	}
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

// GetValue gets value from db
func GetValue(s1 string, s2 string) float64 {

	if s1 == s2{
		return 1
	}

	// Set up the database
	db := SetupDB()
	db.Init()

	// Get today's currencies for today
	data2d, ok := db.GetLatest("noDate")

	// If there isn't any data in the db for today
	if ok == false {
		// If there's no data, log to heroku
		log.Println("There is no data to get", 404)
		return 0
	}
	return data2d.Data[s1][s2]
}
