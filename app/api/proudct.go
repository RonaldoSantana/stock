package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rmsj/stock/app/helper"
	"github.com/rmsj/stock/db/models"
	"github.com/rmsj/stock/db/modelx"
)

// AddProduct adds a new product in DB
func (api *API) AddProduct(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	product := models.Product{}
	err := helper.LoadFromJSON(request.Body, &product)
	if err != nil {
		api.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	exists, err := modelx.ProductExists(api.DB, product.Sku)
	if err != nil {
		api.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}
	if exists {
		api.WriteError(writer, http.StatusBadRequest, "SKU already registered")
		return
	}

	err = product.Insert(api.DB)
	if err != nil {
		api.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}
}

// UpdateProduct updates product data into DB
func (api *API) UpdateProduct(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	product := models.Product{}
	err := helper.LoadFromJSON(request.Body, &product)
	if err != nil {
		api.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	err = product.Update(api.DB)
	if err != nil {
		api.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}
}

// GetProduct retrieves one product
func (api *API) GetProduct(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {

	sku := params.ByName("sku")
	product, err := modelx.ProductJSONBySKU(api.DB, sku)
	if err != nil {
		api.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	api.WriteJSON(writer, http.StatusOK, product)
}

// GetProduct retrieves one product
func (api *API) UpdateStock(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {

	sku := params.ByName("sku")
	quantity := int(params.ByName("quantity"))
	product, err := modelx.ProductBySKU(api.DB, sku)
	product.Quantity -= quantity
	err = product.Update(api.DB)
	if err != nil {
		api.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	api.WriteJSON(writer, http.StatusCreated, product)
}
