package main

import (
	"database/sql"
	"errors"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Pirce float64 `json:"price"`
}

func (this *product) getProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (this *product) updateProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (this *product) deleteProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (this *product) createProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getProdcut(dq *sql.DB, start, count int) ([]product, error) {
	return nil, errors.New("Not implemented")
}
