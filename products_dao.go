package main

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "products"
)

// Establish a connection to database
func (m *ProductsDAO) Connect() error {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		return err
	}
	db = session.DB(m.Database)
	return nil
}

// Find list of movies
func (m *ProductsDAO) FindAll() ([]Product, error) {
	var products []Product
	err := db.C(COLLECTION).Find(bson.M{}).All(&products)
	return products, err
}

// Find a movie by its id
func (m *ProductsDAO) FindById(id string) (Product, error) {
	var product Product
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&product)
	return product, err
}

// Insert a movie into database
func (m *ProductsDAO) Insert(product Product) error {
	err := db.C(COLLECTION).Insert(&product)
	return err
}

// Delete an existing movie
func (m *ProductsDAO) Delete(product Product) error {
	err := db.C(COLLECTION).Remove(&product)
	return err
}

// Update an existing movie
func (m *ProductsDAO) Update(product Product) error {
	err := db.C(COLLECTION).UpdateId(product.ID, &product)
	return err
}
