package main

import (
	"database/sql"
	"log"
	"os"
	_ "modernc.org/sqlite"
)

func main() {
	path := os.Getenv("BASKETBALL_DB_PATH"); if path == "" { path = "basketball-ranking.db" }
	db, err := sql.Open("sqlite", path); if err != nil { log.Fatal(err) }; defer db.Close()
	if _, err = db.Exec("create table if not exists schema_version(version integer not null)"); err != nil { log.Fatal(err) }
	log.Println("数据库已初始化")
}
