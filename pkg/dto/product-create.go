package dto

import "errors"

type ProductCreate struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *ProductCreate) Validate() error {
	if p.Name == "" {
		return errors.New("name is required")
	}
	if p.Price == 0 {
		return errors.New("price can not be 0")
	}
	return nil
}
