package main

import (
	"database/sql"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (this *product) getProduct(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM products where id=$1", this.ID).Scan(&this.Name, &this.Price)
}

func (this *product) updateProduct(db *sql.DB) error {
	_, err := db.Exec("UPDATE products SET name=$1, price=$2, WHERE id=$3", this.Name, this.Price, this.ID)
	return err
}

func (this *product) deleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", this.ID)

	return err
}

func (this *product) createProduct(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO products(name, price) VALUES ($1, $2) RETURNING id", this.Name, this.Price).Scan(&this.ID)
	if err != nil {
		return err
	}
	return nil
}

func getProducts(db *sql.DB, start, count int) ([]product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products LIMIT $1 OFFSET $2", count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)

	}
	return products, nil
}
