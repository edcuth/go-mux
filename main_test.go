package main_test

import (
	"log"
	"os"
	"testing"
)

var a main.App

func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)

}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreatingQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM PRODUCTS")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

const tableCreatingQuery = `CREATE TABLE IF NOT EXISTS products (
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	price NUMERIC(10, 2) NOT NULL DEFAULT 0.00)`
