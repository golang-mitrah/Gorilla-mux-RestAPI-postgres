package models

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// Defining the product struct
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Get all products
func GetProducts(db *pgx.Conn, start, count int) ([]Product, error) {
	// $ signs are dummy sql variables to replace original value
	var rows, err = db.Query(context.Background(), "SELECT id, name,  price FROM products LIMIT $1 OFFSET $2",
		count, start)

	// rows, err := db.Query(
	// 	"SELECT id, name,  price FROM products LIMIT $1 OFFSET $2",
	// 	count, start)

	if err != nil {
		return nil, err
	}

	// Delaying the execution and closing it when the need is over
	defer rows.Close()

	// Closing the context background is actually closing the connection. So commenting it out.
	// defer db.Close(context.Background())

	products := []Product{}

	for rows.Next() {
		var p Product
		// Scan is used to read and put data into the reference we used below
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

// Get a particular product
func (p *Product) GetProduct(db *pgx.Conn) error {
	// Scan is used to read and put data into the reference we used below
	return db.QueryRow(context.Background(), "SELECT name, price FROM products WHERE id=$1",
		p.ID).Scan(&p.Name, &p.Price)
}

// Update a product
func (p *Product) UpdateProduct(db *pgx.Conn) error {
	_, err :=
		db.Exec(context.Background(), "UPDATE products SET name=$1, price=$2 WHERE id=$3",
			p.Name, p.Price, p.ID)

	return err
}

// Delete a product
func (p *Product) DeleteProduct(db *pgx.Conn) error {
	_, err := db.Exec(context.Background(), "DELETE FROM products WHERE id=$1", p.ID)

	return err
}

// Create a product
func (p *Product) CreateProduct(db *pgx.Conn) error {
	err := db.QueryRow(context.Background(),
		"INSERT INTO products(name, price) VALUES($1, $2) RETURNING id",
		p.Name, p.Price).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}
