package services

import (
	"github.com/hmdrzaa11/got/pkg/dto"
	"github.com/hmdrzaa11/got/pkg/models"
)

// ProductsService is going to contains all the methods that a "productsService" is going to have
type ProductsService interface {
	GetAllProducts() []*dto.ProductResponse
	AddNewProduct(*dto.ProductCreate)
	FindOneProductById(id int) (*models.Product, error)
}

// DefaultProductsService is our struct that satisfies the "ProductsService" interface
type DefaultProductsService struct {
	products models.ProductRepository
}

func (s *DefaultProductsService) GetAllProducts() []*dto.ProductResponse {
	prods := s.products.GetAll()
	var prodsResponse []*dto.ProductResponse

	for _, prod := range prods {
		prodsResponse = append(prodsResponse, prod.ConvertToDTO())
	}

	return prodsResponse
}

func (s *DefaultProductsService) AddNewProduct(prod *dto.ProductCreate) {
	s.products.Add(prod)
}

func (s *DefaultProductsService) FindOneProductById(id int) (*models.Product, error) {
	return s.products.FindOneById(id)

}

// NewDefaultProductsService is going to return a defaultProductsService
func NewDefaultProductsService(repo models.ProductRepository) *DefaultProductsService {
	return &DefaultProductsService{repo}
}
