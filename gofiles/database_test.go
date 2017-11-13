package gofiles

import (
	"gopkg.in/mgo.v2"
	"log"
	"testing"
	"time"
)

type Test struct {
	Date     string
	Base     string
	Target   string
	Currency float64
}

func setupTestDB() *MongoDB {
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

func TestMongoDB_Add(t *testing.T) {

	var data Data2d
	data.Data = make(map[string]map[string]float64)
	Add2d(data.Data, "EUR", "NOK", 9.5)

	db := setupTestDB()
	db.Init()
	ok := db.Add(data)
	if ok != nil {
		t.Fatalf("Could not add to testDB")
	}
}

func TestMongoDB_GetLatest(t *testing.T) {

	/*
		db := SetupTestDB()
		db.Init()

		session, err := mgo.Dial(db.DatabaseURL)
		if err !=nil {
			t.Fatal(err.Error())
		}

		defer session.Close()


		data2d, ok := db.GetLatest("noDate")
		if ok == false{
			t.Fatal("Could not get latest currency in db")
		}

	*/
}

func TestMongoDB_Count(t *testing.T) {
	/*
		db := SetupTestDB()
		db.Init()

		count := db.Count()
	*/
}

func TestGetValue(t *testing.T) {
	out := []string{"NOK", "EUR"}
	date := time.Now()
	dateCopy := date.Format("2006-01-02")

	// Set up the database
	db := SetupDB()
	db.Init()

	// Get today's currencies for today
	data2d, ok := db.GetLatest(dateCopy)

	// If there isn't any data in the db
	if ok == false {
		date = date.AddDate(0, 0, -1)
		dateCopy = date.Format("2006-01-02")
		data2d, ok = db.GetLatest(dateCopy)
		if ok == false {
			t.Fatalf("ERROR could not retrieve data from db")
		}
	}

	realValue := data2d.Data[out[0]][out[1]]

	testValue := GetValue(out[0], out[1])

	if realValue != testValue {
		t.Fatalf("ERROR, expected %v, got %v\n", realValue, testValue)
	}
}
