package controllers

import (
	"net/http"

	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/rmsj/stock/app/helper"
	"github.com/rmsj/stock/db/models"
	"github.com/rmsj/stock/db/modelx"
)

// ProductController type
type ProductController struct{}

// NewProductController returns a ProductController
func NewProductController() *ProductController {
	return &ProductController{}
}

// Add adds a new product in DB
func (pc *ProductController) Add(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	product := models.Product{}
	err := helper.LoadFromJSON(request.Body, &product)
	if err != nil {
		WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	exists, err := modelx.ProductExists(db, product.Sku)
	if err != nil {
		WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}
	if exists {
		WriteError(writer, http.StatusBadRequest, "SKU already registered")
		return
	}

	err = product.Insert(db)
	if err != nil {
		WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}
}

// Update updates product data into DB
func (pc *ProductController) Update(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	product := models.Product{}
	err := helper.LoadFromJSON(request.Body, &product)
	if err != nil {
		WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	err = product.Update(db)
	if err != nil {
		WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}
}

// Get retrieves one product
func (pc *ProductController) Get(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {

	sku := params.ByName("sku")
	product, err := modelx.ProductJSONBySKU(db, sku)
	if err != nil {
		WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	WriteJSON(writer, http.StatusOK, product)
}

// UpdateStock retrieves one product
func (pc *ProductController) UpdateStock(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {

	sku := params.ByName("sku")
	quantity, err := strconv.Atoi(params.ByName("quantity"))
	if err != nil {
		WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}
	product, err := modelx.ProductBySKU(db, sku)
	if err != nil {
		WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}
	product.Quantity -= quantity
	err = product.Update(db)
	if err != nil {
		WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	WriteJSON(writer, http.StatusCreated, product)
}
