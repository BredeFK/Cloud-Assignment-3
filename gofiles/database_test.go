package gofiles

import (
	"testing"
)


func TestMongoDB_Add(t *testing.T) {
	/*
	data2d := GetCurrency()
	db := SetupDB()
	db.Init()
	db.Add(data2d)
	*/
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


