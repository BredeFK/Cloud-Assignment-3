package gofiles

import (
	"testing"
	"time"
)

func TestAdd2d(t *testing.T) {
	var data2d Data2d
	data2d.Data = make(map[string]map[string]float64)

	//sets date
	data2d.Date = time.Now().Format("2006-01-02")

	out := []string{"test1", "test2"}
	outF := 4.20

	Add2d(data2d.Data, out[0], out[1], outF)

	if data2d.Data[out[0]] == nil {
		t.Fatalf("ERROR, expected: %s%v, got %s", out[1], outF, "nil")
	}

	if data2d.Data[out[0]][out[1]] != outF {
		t.Fatalf("ERROR, expected: %v, got %v\n", outF, data2d.Data[out[0]][out[1]])
	}

}

func TestGetCurrency(t *testing.T) {

	out := []string{"EUR", "NOK"}
	db := SetupTestDB()
	db.Init()

	data2d := GetCurrency()
	db.Add(data2d)

	testValue := db.GetValue(out[0], out[1])
	db.DropDB()

	if testValue != data2d.Data[out[0]][out[1]] {
		t.Fatalf("ERROR expected: %v got %v", testValue, data2d.Data[out[0]][out[1]])
	}

}
