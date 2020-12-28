package database

import (
	"testing"
)

func TestInitDB(t *testing.T) {
	InitDB()
	if DBCon == nil {
		t.Error("DB Connection Failed")
	}

	if DBCon != nil {
		t.Log("DB Connected Success")
	}

	DBCon.Debug()
}
