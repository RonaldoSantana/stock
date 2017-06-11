package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rmsj/stock/app/helper"
	"github.com/rmsj/stock/db/models"
	"github.com/rmsj/stock/db/modelx"
)

// ProductController type
type ProductController struct {
	baseController
}

// AddProduct adds a new product in DB
func (pc *ProductController) AddProduct(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	product := models.Product{}
	err := helper.LoadFromJSON(request.Body, &product)
	if err != nil {
		pc.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	exists, err := modelx.ProductExists(pc.DB, product.Sku)
	if err != nil {
		pc.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}
	if exists {
		pc.WriteError(writer, http.StatusBadRequest, "SKU already registered")
		return
	}

	err = product.Insert(pc.DB)
	if err != nil {
		pc.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}
}

// UpdateProduct updates product data into DB
func (pc *ProductController) UpdateProduct(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	product := models.Product{}
	err := helper.LoadFromJSON(request.Body, &product)
	if err != nil {
		pc.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	err = product.Update(pc.DB)
	if err != nil {
		pc.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}
}

// GetProduct retrieves one product
func (pc *ProductController) GetProduct(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {

	sku := params.ByName("sku")
	product, err := modelx.ProductJSONBySKU(pc.DB, sku)
	if err != nil {
		pc.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	pc.WriteJSON(writer, http.StatusOK, product)
}

// GetProduct retrieves one product
func (pc *ProductController) UpdateStock(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {

	sku := params.ByName("sku")
	quantity := int(params.ByName("quantity"))
	product, err := modelx.ProductBySKU(pc.DB, sku)
	product.Quantity -= quantity
	err = product.Update(pc.DB)
	if err != nil {
		pc.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	pc.WriteJSON(writer, http.StatusCreated, product)
}
