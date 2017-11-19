package gofiles

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
)

// SetupDB sets up the database
func SetupDB() *MongoDB {

	// Make new MongoDB struct
	db := MongoDB{
		os.Getenv("MONGODB"),
		"heroku_pgvgprmm",
		"currencyCollection",
	}

	// Dial the dbURL
	session, err := mgo.Dial(db.DatabaseURL)

	// Close session
	defer session.Close()

	// If error is not 0
	if err != nil {

		// Log error to heroku logs
		log.Fatal(err.Error())
	}

	// Return struct db
	return &db
}

// Init initialising the db
func (db *MongoDB) Init() {

	// Dial the dbURL
	session, err := mgo.Dial(db.DatabaseURL)

	// Close session
	defer session.Close()

	// If error is not 0
	if err != nil {

		// Log error to heroku logs
		log.Fatal(err.Error())
	}

	// Index
	index := mgo.Index{
		Key:        []string{"currencyid"},
		Unique:     true,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	}

	// Ensure Index
	err = session.DB(db.DatabaseName).C(db.ColCurrency).EnsureIndex(index)

	// If error is not 0
	if err != nil {

		// Log error to heroku logs
		log.Fatal(err.Error())
	}
}

// Add adds the db
func (db *MongoDB) Add(data Data2d) error {

	// Dial the dbURL
	session, err := mgo.Dial(db.DatabaseURL)

	// Close session
	defer session.Close()

	// If error is not 0
	if err != nil {

		// Log error to heroku logs
		log.Fatal(err.Error())
	}

	// Insert data(Data2d struct) to <dbName><dbCollection>
	err = session.DB(db.DatabaseName).C(db.ColCurrency).Insert(data)

	// If error is not 0
	if err != nil {

		// Log error with error message to heroku logs
		log.Printf("Could not add to db, error in Insert(): %v", err.Error())

		// Return error
		return err
	}

	// Return 0 (return 0 errors)
	return nil
}

// Count counts the number of currencies in Collection
func (db *MongoDB) Count() int {

	// Dial the dbURL
	session, err := mgo.Dial(db.DatabaseURL)

	// Close session
	defer session.Close()

	// If error is not 0
	if err != nil {

		// Log error to heroku logs
		log.Fatal(err)
	}

	// Count number of currencies in <dbName><dbCollection>
	count, err := session.DB(db.DatabaseName).C(db.ColCurrency).Count()

	// If error is not 0
	if err != nil {

		// Log error with error message to heroku logs
		log.Printf("Error in Count(): %v", err.Error())

		// return -1
		return -1
	}

	// Return count of currencies
	return count
}

// GetLatest gets the latest currencies with date as index
func (db *MongoDB) GetLatest(date string) (Data2d, bool) {

	// Get count from Count func
	count := db.Count()

	// Dial the dbURL
	session, err := mgo.Dial(db.DatabaseURL)

	// Close session
	defer session.Close()

	// If error is not 0
	if err != nil {

		// Log error to heroku logs
		log.Fatal(err.Error())
	}

	// data2d is an Data2d struct
	data2d := Data2d{}

	// Declare bool and set it to be true
	notToday := true

	// If dev. hasn't set *date string* to be "NoDate" (This is for easier testing)
	if date != "noDate" {

		// Get currency from said date from <dbName><Collection>
		err = session.DB(db.DatabaseName).C(db.ColCurrency).Find(bson.M{"date": date}).One(&data2d)

		// Else if the date is "noDate"
	} else {

		// Get latest inserted currency from <dbName><Collection>
		err = session.DB(db.DatabaseName).C(db.ColCurrency).Find(nil).Skip(count - 1).One(&data2d)
	}

	// If error is not null
	if err != nil {

		// Date is not found. (This could be better implemented :/ )
		notToday = false

		// Log error to heroku logs
		log.Fatal(err.Error())
	}

	// Return struct and the bool: notToday
	return data2d, notToday
}

// DailyCurrencyAdder adds currency once a day
func (db *MongoDB) DailyCurrencyAdder(data Data2d) {

	// Initialize
	db.Init()

	// Add
	db.Add(data)
}

// GetValue gets value from db
func GetValue(s1 string, s2 string) float64 {

	// If base is equal to target currency
	if s1 == s2 {

		// return 1
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

		// Return 0
		return 0
	}

	return data2d.Data[s1][s2]
}
