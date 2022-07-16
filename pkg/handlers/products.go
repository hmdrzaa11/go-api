package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/hmdrzaa11/got/pkg/app"
	"github.com/hmdrzaa11/got/pkg/dto"
	"github.com/hmdrzaa11/got/pkg/services"
)

type Product struct {
	app             *app.Application
	productsService services.ProductsService
}

func (p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.getAllProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	if r.Method == http.MethodPut {
		//we need the id to be extracted
		//the base url for this handler is "/" and id should be after it like "/12" and we need to extrac that "12"
		path := r.URL.Path
		reg := regexp.MustCompile(`/([0-9]+)`)       //we put it in a group "()" because we want to extract it
		group := reg.FindAllStringSubmatch(path, -1) //we want to match all WHY because user may pass a URL like "/12/12/123/344" then if we jst
		//match 1 of them its going to pass WHICH is bad but if we match all of them and we get an array of length more than 1
		//we know that this is invalid url
		//check the length to see if not 1 means or we do not have 1 id or we have many of them in both case its invalid
		if len(group) != 1 {
			p.app.SendErrorResponse(w, http.StatusBadRequest, errors.New("invalid id"))
			return
		}

		//also we need to see if match has 2 items inside of it one is the "/12" and "12"
		if len(group[0]) != 2 {
			p.app.SendErrorResponse(w, http.StatusBadRequest, errors.New("invalid id"))
			return
		}

		//now we need to turn it into id integer
		idStr := group[0][1] //g[0] gives you the group and because index 0 inside is the "/12" and index  1 is the actual "12"
		id, err := strconv.Atoi(idStr)
		if err != nil {
			p.app.SendErrorResponse(w, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		p.updateProduct(w, r, id)
		return
	}

	//catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getAllProducts(w http.ResponseWriter, r *http.Request) {

	prods := p.productsService.GetAllProducts()
	err := json.NewEncoder(w).Encode(prods)
	if err != nil {
		panic(err)
	}
}

func (p *Product) addProduct(w http.ResponseWriter, r *http.Request) {
	rawProd := &dto.ProductCreate{}
	err := json.NewDecoder(r.Body).Decode(rawProd)
	fmt.Println(rawProd)
	if err != nil {
		p.app.Logger.Println(err)
		p.app.SendErrorResponse(w, http.StatusBadRequest, errors.New("invalid data"))
		return
	}
	p.productsService.AddNewProduct(rawProd)
	p.app.SendResponse(w, http.StatusCreated, nil)
}

func (p *Product) updateProduct(w http.ResponseWriter, r *http.Request, id int) {
	//find the product
	prod, err := p.productsService.FindOneProductById(id)
	if err != nil {
		//means not found
		p.app.SendErrorResponse(w, http.StatusNotFound, errors.New("product not found"))
		return
	}
	//turn user input into a dto to run validation on them
	var rawProd dto.ProductUpdate
	err = json.NewDecoder(r.Body).Decode(&rawProd)
	if err != nil {
		p.app.SendErrorResponse(w, http.StatusBadRequest, errors.New("invalid data"))
		return
	}
	prod.Name = rawProd.Name
	prod.Price = rawProd.Price

	p.app.SendResponse(w, http.StatusOK, prod.ConvertToDTO())

}

func NewProductHandler(app *app.Application, service services.ProductsService) *Product {
	return &Product{app, service}
}
