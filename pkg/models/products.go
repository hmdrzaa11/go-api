package models

import (
	"errors"

	"github.com/hmdrzaa11/got/pkg/dto"
)

// ProductRepository its going to represent all the methods this model have
type ProductRepository interface {
	GetAll() []*Product
	Add(prod *dto.ProductCreate)
	FindOneById(id int) (*Product, error)
}

// Product represents data in the database.
type Product struct {
	Id    int
	Name  string
	Price float64
}

// ConvertToDTO is going to turn a SINGLE product into dto
func (p *Product) ConvertToDTO() *dto.ProductResponse {
	return &dto.ProductResponse{
		Id:    p.Id,
		Name:  p.Name,
		Price: p.Price,
	}
}

// ProductRepositoryDB this is the struct that is going to staisfy the "ProductRepository" interface
type ProductRepositoryDB struct {
	//all the dependencies that you need to intract with db
}

// GetAll returns all the products
func (p *ProductRepositoryDB) GetAll() []*Product {
	return productsStubs
}

func (p *ProductRepositoryDB) Add(d *dto.ProductCreate) {
	prod := &Product{
		Id:    generateProdId(),
		Name:  d.Name,
		Price: d.Price,
	}
	productsStubs = append(productsStubs, prod)
}

func (p *ProductRepositoryDB) FindOneById(id int) (*Product, error) {
	for _, prod := range productsStubs {
		if prod.Id == id {
			return prod, nil
		}
	}
	return nil, errors.New("product not found")
}

// NewProductRepository returns a new productRepo that also satisfies the "ProductRepository"
func NewProductRepository() *ProductRepositoryDB {
	return &ProductRepositoryDB{}
}

func generateProdId() int {
	return productsStubs[len(productsStubs)-1].Id + 1
}

var productsStubs = []*Product{
	{Id: 1, Name: "BMW", Price: 32000},
	{Id: 2, Name: "Benz", Price: 42000},
	{Id: 3, Name: "Ferrari", Price: 132000},
}
