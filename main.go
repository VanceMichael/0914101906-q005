package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func databasePath() string {
	if path := os.Getenv("BASKETBALL_DB_PATH"); path != "" { return path }
	return "basketball-ranking.db"
}

func openDatabase() (*sql.DB, error) { return sql.Open("sqlite", databasePath()) }
func port() string { if value := os.Getenv("PORT"); value != "" { return value }; return "8080" }

func main() {
	db, err := openDatabase(); if err != nil { panic(err) }; defer db.Close()
	router := gin.Default()
	router.GET("/health", func(context *gin.Context) {
		if err := db.Ping(); err != nil { context.JSON(http.StatusServiceUnavailable, gin.H{"status": "error"}); return }
		context.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	if err := router.Run(":" + port()); err != nil { panic(err) }
}
