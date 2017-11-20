package gofiles

import (
	"gopkg.in/mgo.v2"
	"log"
	"testing"
)

func SetupTestDB() *MongoDB {

	db := MongoDB{
		"mongodb://localhost",
		"testDB",
		"currency",
	}

	session, err := mgo.Dial(db.DatabaseURL)
	defer session.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	return &db
}

func (db *MongoDB) DropDB() {

	session, err := mgo.Dial(db.DatabaseURL)
	defer session.Close()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = session.DB(db.DatabaseName).DropDatabase()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func TestMongoDB_Add(t *testing.T) {

	var data Data2d
	data.Data = make(map[string]map[string]float64)
	Add2d(data.Data, "EUR", "NOK", 9.5)

	db := SetupTestDB()
	db.Init()
	ok := db.Add(data)
	if ok != nil {
		t.Fatalf("Could not add to testDB")
	}
}

func TestMongoDB_GetLatest(t *testing.T) {

	var data Data2d
	data.Data = make(map[string]map[string]float64)
	Add2d(data.Data, "EUR", "NOK", 9.5)

	db := SetupTestDB()
	db.Init()

	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		t.Fatal(err.Error())
	}

	defer session.Close()

	testData, ok := db.GetLatest("noDate")
	if ok == false {
		t.Fatal("Could not get latest currency in db")
	}

	if testData.Data["EUR"]["NOK"] != data.Data["EUR"]["NOK"] {
		t.Fatalf("Could not get correct data in testGetLatest()")
	}
}

func TestMongoDB_Count(t *testing.T) {

	db := SetupTestDB()
	db.Init()

	count := db.Count()

	if count == -1 {
		t.Errorf("Could not get count")
	}

	if count > 1 {
		t.Errorf("Too much data in db")
	}
}

func TestGetValue(t *testing.T) {
	out := []string{"NOK", "EUR"}

	// Set up the database
	db := SetupTestDB()
	db.Init()

	// Get today's currencies for today
	data2d, ok := db.GetLatest("noDate")

	// If there isn't any data in the db
	if ok == false {
		t.Fatalf("ERROR could not retrieve data from db")
	}

	realValue := data2d.Data[out[0]][out[1]]

	testValue := db.GetValue(out[0], out[1])

	if realValue != testValue {
		t.Fatalf("ERROR, expected %v, got %v\n", realValue, testValue)
	}
}

func TestDailyCurrencyAdder(t *testing.T) {

	var data Data2d
	data.Data = make(map[string]map[string]float64)
	Add2d(data.Data, "NOK", "DDK", 0.8)

	db := SetupTestDB()
	db.Init()
	db.DailyCurrencyAdder(data)
	db.DropDB()
}
