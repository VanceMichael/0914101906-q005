package main

import "testing"

func TestDatabaseCanBeOpened(t *testing.T) {
	t.Setenv("BASKETBALL_DB_PATH", t.TempDir()+"/test.db")
	db, err := openDatabase(); if err != nil { t.Fatal(err) }; defer db.Close()
	if err := db.Ping(); err != nil { t.Fatal(err) }
}
